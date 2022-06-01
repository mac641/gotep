package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/mac641/gotep/src/lib"
	"github.com/mac641/gotep/src/lib/context"
	"github.com/mac641/gotep/src/lib/logger"
)

type Parser struct{}

// TODO: if needed, check if possible to store bool and string in same map -> until then everything is type of string
var (
	config = map[string]string{}
	ctx    = context.GetContext()
	log    = logger.GetLogger()
)

// Takes array of strings representing prepared http requests and parses them into array of http.Requests.
func (p *Parser) Parse(requests []string) ([]http.Request, error) {
	httpRequests := []http.Request{}
	isFirstHttpVersionOccurrence := true
	for _, request := range requests {
		// Parse request line
		requestLine := lib.RegRequestLine.FindString(request)
		method, requestUrl, httpVersion, err := p.parseRequestLine(requestLine)
		if err != nil {
			return httpRequests, err
		}

		stringHeaderMessage := strings.Trim(strings.ReplaceAll(request, requestLine, ""), "\r\n")

		// Separate headers / message body and parse them afterwards
		if lib.RegMultipartFormDataHeader.MatchString(stringHeaderMessage) {
			// TODO: add multipart/form-data support
			log.Warnln("NOTE: currently multipart/form-data can't be parsed and, therefore, this request will be skipped")
			continue // Skip this loop cycle
		}
		splitEmptyNewline := lib.RegEmptyNewLine.Split(stringHeaderMessage, -1)
		parsedHeaders := make(map[string][]string)
		var message io.Reader
		switch len(splitEmptyNewline) {
		case 0:
			log.Infof("neither headers, nor a message body have/has been provided for\n%s\n", request)
		case 1:
			match := splitEmptyNewline[0]
			if lib.RegHeaders.MatchString(match) {
				log.Infof("no message body has been provided for\n%s\n", request)
				parsedHeaders, err = p.parseHeaders(lib.RegLineEnding.Split(splitEmptyNewline[0], -1))
				if err != nil {
					return httpRequests, err
				}
			} else {
				log.Infof("no headers have been provided for\n%s\n", request)
				message, err = p.parseMessage(match)
				if err != nil {
					return httpRequests, err
				}
			}
		case 2:
			log.Infof("headers and message body detected in\n%s\n", request)
			parsedHeaders, err = p.parseHeaders(lib.RegLineEnding.Split(splitEmptyNewline[0], -1))
			if err != nil {
				return httpRequests, err
			}
			message, err = p.parseMessage(splitEmptyNewline[1])
			if err != nil {
				return httpRequests, err
			}
		default:
			return httpRequests,
				fmt.Errorf("too many in-between line breaks detected!\nCheck request below for errors:\n%s",
					request)
		}

		// Assemble request
		var httpRequest *http.Request
		switch method {
		case "", "GET":
			httpRequest, err = http.NewRequest(http.MethodGet, requestUrl, message)
		case "HEAD":
			httpRequest, err = http.NewRequest(http.MethodHead, requestUrl, message)
		case "POST":
			httpRequest, err = http.NewRequest(http.MethodPost, requestUrl, message)
		case "PUT":
			httpRequest, err = http.NewRequest(http.MethodPut, requestUrl, message)
		case "DELETE":
			httpRequest, err = http.NewRequest(http.MethodDelete, requestUrl, message)
		case "CONNECT":
			httpRequest, err = http.NewRequest(http.MethodConnect, requestUrl, message)
		case "PATCH":
			httpRequest, err = http.NewRequest(http.MethodPatch, requestUrl, message)
		case "OPTIONS":
			httpRequest, err = http.NewRequest(http.MethodOptions, requestUrl, message)
		case "TRACE":
			httpRequest, err = http.NewRequest(http.MethodTrace, requestUrl, message)
		default:
			return httpRequests, fmt.Errorf("undefined http method value found in\n%s", request)
		}
		if err != nil {
			return httpRequests, err
		}

		// NOTE: http version can only be set when acting as server
		if httpVersion != "" {
			if isFirstHttpVersionOccurrence {
				log.Warnln("NOTE: http versions can't be set and, therefore, will be ignored")
				isFirstHttpVersionOccurrence = false
			}
			// httpRequest.Proto = httpVersion
		}

		if len(parsedHeaders) > 0 {
			httpRequest.Header = parsedHeaders
		}

		// NOTE: Handle origin-form and asterisk-form of url in specific way
		if httpRequest.URL.Host == "" && len(parsedHeaders["Host"]) > 0 {
			// TODO: give user chance to decide which host shall be used if their are multiple
			host := parsedHeaders["Host"][0]
			if !lib.RegUrlScheme.MatchString(host) && httpRequest.URL.Scheme == "" {
				httpRequest.URL.Scheme = "http"
				httpRequest.URL.Host = host
			} else {
				hostSplit := strings.Split(host, "://")
				if len(hostSplit) > 2 {
					return httpRequests, fmt.Errorf("the \"Host\" header %s is not formatted properly", host)
				}
				if httpRequest.URL.Scheme == "" {
					httpRequest.URL.Scheme = hostSplit[0]
				}
				httpRequest.URL.Host = hostSplit[1]
			}
		}

		httpRequests = append(httpRequests, *httpRequest)
	}

	return httpRequests, nil
}

func (p *Parser) ParseConfig() error {
	var tempConfig map[string]interface{}
	// TODO: don't load whole files into RAM -> read by byte
	jsonData, err := os.ReadFile(ctx.GetConfigFilePath())
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &tempConfig)
	if err != nil {
		return err
	}

	envMap, configKeyExists := tempConfig[ctx.GetConfigEnvironment()]
	if !configKeyExists {
		return fmt.Errorf("%s does not exist in your config file", ctx.GetConfigEnvironment())
	}

	for key, val := range envMap.(map[string]interface{}) {
		typeOfVal := reflect.TypeOf(val).Kind().String()
		if typeOfVal != "string" && typeOfVal != "bool" {
			return fmt.Errorf("your config contains key %s with value %s which is not of type string or bool",
				key, val)
		}

		config[key] = val.(string)
	}

	log.Infof("using config environment: %s\n", ctx.GetConfigEnvironment())
	return nil
}

// Takes string containing http requests and splits them by separators, removes comments and empty requests
// and inserts env variable values, if possible.
// Otherwise, exits program with exit code != 0.
func (p *Parser) Prepare(file string) ([]string, error) {
	var err error
	// Split file by request separators.
	// Exit, if no requests can be found after splitting.
	requests := lib.RegSeparator.Split(file, -1)
	if len(requests) == 0 {
		return nil, errors.New("no requests could be parsed")
	}

	// Remove empty requests and comments
	noEmptyRequests := []string{}
	for _, request := range requests {
		if len(request) == 0 {
			continue
		}

		noEmptyRequests = append(noEmptyRequests, lib.RegComments.ReplaceAllLiteralString(request, ""))
	}
	requests = noEmptyRequests

	isFirstResponseHandlerOccurrence := true
	isFirstResponseRefOccurrence := true
	for i, request := range requests {
		// Insert env variable values from json config.
		// Exit, if one variable does not exist in json config.
		envMatches := lib.RegEnv.FindAllString(request, -1)

		if envMatches != nil {
			matchedConfig := p.matchConfig(envMatches)

			// TODO: use one by one comparison to ensure every match is represented in env variable config
			if len(envMatches) != len(matchedConfig) {
				return requests, errors.New("there are undefined env variables present in your requests file")
			}

			for configKey, configValue := range matchedConfig {
				request = strings.ReplaceAll(request, fmt.Sprintf("{{%s}}", configKey), configValue)
			}
		}

		// Remove response handler from request and prompt user that they are not going to be used
		responseHandlerMatches := lib.RegResponseHandler.FindAllString(request, -1)
		if len(responseHandlerMatches) > 1 {
			return requests, errors.New("some of your requests contain too many response handlers")
		}

		// NOTE: following TODOs are only relevant when adding response handler validation support
		// TODO: trim '{%%}' from response handler match
		// TODO: check if request separators or '%}' have been used inside the handler block

		if responseHandlerMatches != nil {
			if isFirstResponseHandlerOccurrence {
				log.Warnln("NOTE: currently response handlers can't be validated and, therefore, will be ignored")
				isFirstResponseHandlerOccurrence = false
			}
			request = lib.RegResponseHandler.ReplaceAllLiteralString(request, "")
		}

		// Remove response ref from request and prompt user that they are not going to be used
		responseRefMatches := lib.RegResponseRef.FindAllString(request, -1)
		if len(responseRefMatches) > 1 {
			return requests, errors.New("some of your requests contain too many response refs")
		}

		if responseRefMatches != nil {
			if isFirstResponseRefOccurrence {
				log.Warnln("NOTE: currently response references can't be validated and, therefore, will be ignored")
				isFirstResponseRefOccurrence = false
			}
			request = lib.RegResponseRef.ReplaceAllLiteralString(request, "")
		}

		requests[i] = request
	}

	return requests, err
}

// Takes array of environment variable matches and compares them to keys stored in config.
// It returns an array of all config keys that have been matched.
func (p *Parser) matchConfig(envMatches []string) map[string]string {
	matchedConfig := make(map[string]string)

	for _, match := range envMatches {
		match = strings.Trim(match, "{}")

		for configKey, configValue := range config {
			if strings.Contains(configKey, match) {
				matchedConfig[match] = configValue
			}
		}
	}

	return matchedConfig
}

// Takes string containing headers and message and separates them.
// It returns headers as string array and message body as single string.
// TODO: outsource separating headers and messages from Parse when adding multipart/form-data support
// func separateHeadersFromMessage(headersMessage string) ([]string, string) {

// }

// Takes array of header strings splitted by new line, detects related header values and adds them to map
func (p *Parser) parseHeaders(headers []string) (map[string][]string, error) {
	parsedHeaders := make(map[string][]string)

	var fieldName string
	var err error
	for _, header := range headers {
		header = strings.Trim(header, "\r\n \t\f")
		if lib.RegHeaders.MatchString(header) {
			headerSubMatch := lib.RegHeaders.FindStringSubmatch(header)
			for i, name := range lib.RegHeaders.SubexpNames() {
				// NOTE: start at index == 1 because first value of array is full match which is irrelevant
				if i >= 1 && i < len(headerSubMatch) {
					switch name {
					case "Fieldname":
						parsedHeaders[headerSubMatch[i]] = []string{}
						fieldName = headerSubMatch[i]
					case "Fieldvalue":
						parsedHeaders[fieldName] = append(parsedHeaders[fieldName], headerSubMatch[i])
					default:
						err = errors.New("misconfigured header found")
					}
				}
			}
		} else {
			if len(parsedHeaders) > 0 && header != "" && fieldName != "" {
				parsedHeaders[fieldName] = append(parsedHeaders[fieldName], header)
			} else if header == "" {
				continue
			} else {
				err = errors.New("misconfigured header found")
			}
		}
	}

	return parsedHeaders, err
}

// Takes message string, detects whether it is a filepath or direct message strings and returns them as io.Reader
func (p *Parser) parseMessage(message string) (io.Reader, error) {
	var file io.Reader
	if lib.RegInputFileRef.MatchString(message) {
		message = strings.TrimSpace(strings.TrimPrefix(message, "<"))
		absMessagePath, err := lib.ConvertToAbsolutePath(message)
		if err != nil {
			return file, err
		}
		return os.Open(absMessagePath)
	} else {
		return strings.NewReader(message), nil
	}
}

// Takes string representing a request line and parses it.
func (p *Parser) parseRequestLine(reqLine string) (method string, requestUrl string, httpVersion string, err error) {
	requestLine := reqLine
	// Check if request line has been split
	requestLineSplit := lib.TrimRightEmptyStringsFromSlice(regexp.MustCompile("\r?\n|\r").Split(requestLine, -1))
	if len(requestLineSplit) > 1 {
		requestLine = ""
		for _, line := range requestLineSplit {
			line = strings.Trim(line, " \t\f")
			line = strings.TrimRight(line, "\r\n")
			requestLine += line
		}
	} else {
		requestLine = strings.TrimRight(requestLine, "\r\n")
	}

	// Assign method, requestUrl and httpVersion
	requestLineMatches := regexp.MustCompile("[ \t\f]").Split(requestLine, -1)
	method = ""
	requestUrl = ""
	httpVersion = ""
	switch len(requestLineMatches) {
	case 1:
		requestUrl = requestLineMatches[0]
	case 2:
		if strings.Contains(requestLineMatches[1], "HTTP/") {
			requestUrl = requestLineMatches[0]
			httpVersion = requestLineMatches[1]
		} else {
			method = requestLineMatches[0]
			requestUrl = requestLineMatches[1]
		}
	case 3:
		method = requestLineMatches[0]
		requestUrl = requestLineMatches[1]
		httpVersion = requestLineMatches[2]
	default:
		return method, requestUrl, httpVersion, errors.New("one of your request lines could not be parsed")
	}

	// Check if requestUrl contains fragment and remove it, because browsers do not send fragments
	fragment := lib.RegRequestLineFragment.FindString(requestUrl)
	if fragment != "" {
		requestUrl = strings.ReplaceAll(requestUrl, fragment, "")
	}

	// Check if requestUrl contains path or query elements and needs to be encoded
	// or requestUrl is already encoded because of whitespace and needs to be decoded
	query := lib.RegRequestLineQuery.FindString(requestUrl)
	if query != "" {
		if !strings.Contains(query, "%") && lib.RegNonAscii.MatchString(query) {
			queryEncoded := url.QueryEscape(query)
			requestUrl = strings.ReplaceAll(requestUrl, query, queryEncoded)
		}
		if strings.Contains(query, "%") {
			queryDecoded, err := url.QueryUnescape(query)
			if err != nil {
				return method, requestUrl, httpVersion, err
			}
			if !lib.RegNonAscii.MatchString(queryDecoded) {
				requestUrl = strings.ReplaceAll(requestUrl, query, queryDecoded)
			}
		}
	}

	pathSegment := lib.RegRequestLinePathSegment.FindString(requestUrl)
	if pathSegment != "" {
		// NOTE: According to go docs, PathUnescape does not convert + to ' ', which is correct, unlike stated by
		// JetBrains grammar
		if !strings.Contains(pathSegment, "%") && lib.RegNonAscii.MatchString(pathSegment) {
			pathSegmentEncoded := url.PathEscape(pathSegment)
			requestUrl = strings.ReplaceAll(requestUrl, pathSegment, pathSegmentEncoded)
		}
		if strings.Contains(pathSegment, "%") {
			pathSegmentDecoded, err := url.PathUnescape(pathSegment)
			if err != nil {
				return method, requestUrl, httpVersion, err
			}
			if !lib.RegNonAscii.MatchString(pathSegmentDecoded) {
				requestUrl = strings.ReplaceAll(requestUrl, pathSegment, pathSegmentDecoded)
			}
		}
	}

	// Validate url and add scheme, if missing
	if !lib.IsUrlValid(requestUrl) {
		return method, requestUrl, httpVersion, fmt.Errorf("%s is not a valid URL", requestUrl)
	}
	// NOTE: ensure ip addresses will be stored as hosts, otherwise creating requests fails
	if !lib.RegUrlScheme.MatchString(requestUrl) && lib.RegIp.MatchString(requestUrl) {
		requestUrl = "http://" + requestUrl
	}

	return method, requestUrl, httpVersion, nil
}
