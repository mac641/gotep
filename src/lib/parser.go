package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type Parser struct {
	ConfigEnv  string
	ConfigPath string
	PathPrefix string
	Verbose    bool
}

// TODO: if needed, check if possible to store bool and string in same map -> until then everything is type of string
var config = map[string]string{}

// Takes array of strings representing prepared http requests and parses them into array of http.Requests.
func (p *Parser) Parse(requests []string) []http.Request {
	httpRequests := []http.Request{}
	isFirstHttpVersionOccurrence := true
	for i := range requests {
		stringRequest := requests[i]

		// Parse request line
		requestLine := strings.TrimRight(regRequestline.FindString(stringRequest), "\r\n")
		requestLineMatches := regexp.MustCompile("[ \t\f]").Split(requestLine, -1)
		method := ""
		url := ""
		httpVersion := ""
		switch len(requestLineMatches) {
		case 1:
			url = requestLineMatches[0]
		case 2:
			if strings.Contains(requestLineMatches[1], "HTTP/") {
				url = requestLineMatches[0]
				httpVersion = requestLineMatches[1]
			} else {
				method = requestLineMatches[0]
				url = requestLineMatches[1]
			}
		case 3:
			method = requestLineMatches[0]
			url = requestLineMatches[1]
			httpVersion = requestLineMatches[2]
		default:
			fmt.Println(Red + "One of your request lines could not be parsed!" + Reset)
			fmt.Println(Yellow + "Check request below for errors:" + Reset)
			fmt.Println(stringRequest)
			os.Exit(1)
		}
		stringHeaderMessage := strings.Trim(strings.ReplaceAll(stringRequest, requestLine, ""), "\r\n")

		if !IsUrlValid(url) {
			log.Fatalf("%s is not a valid URL. Exiting...", url)
		}
		// Ensure ip addresses will be stored as hosts, otherwise creating requests fails
		if !regUrlScheme.MatchString(url) && regIp.MatchString(url) {
			url = "http://" + url
		}

		// Separate headers / message body and parse them afterwards
		if regMultipartFormDataHeader.MatchString(stringHeaderMessage) {
			// TODO: Add multipart/form-data support
			fmt.Println(Yellow +
				"NOTE: Currently multipart/form-data can't be parsed. Therefore, this request will be skipped!" + Reset)
			continue
		}
		splitEmptyNewline := regEmptyNewline.Split(stringHeaderMessage, -1)
		parsedHeaders := make(map[string][]string)
		var message io.Reader
		switch len(splitEmptyNewline) {
		case 0:
			LogVerbose("Neither headers, nor a message body have/has been provided for\\n"+stringRequest, p.Verbose)
		case 1:
			match := splitEmptyNewline[0]
			if regHeaders.MatchString(match) {
				LogVerbose("No message body has been provided for\n"+stringRequest, p.Verbose)
				parsedHeaders = p.parseHeaders(regLineEnding.Split(splitEmptyNewline[0], -1))

			} else {
				LogVerbose("No headers have been provided for\n"+stringRequest, p.Verbose)
				message = p.parseMessage(match)
			}
		case 2:
			LogVerbose("Headers and message body detected in\n"+stringRequest, p.Verbose)
			parsedHeaders = p.parseHeaders(regLineEnding.Split(splitEmptyNewline[0], -1))
			message = p.parseMessage(splitEmptyNewline[1])
		default:
			fmt.Println(Red + "Too many in-between line breaks detected!" + Reset)
			fmt.Println(Yellow + "Check request below for errors:" + Reset)
			fmt.Println(stringRequest)
			os.Exit(1)
		}

		// Assemble request
		var httpRequest *http.Request
		var err error
		switch method {
		case "", "GET":
			httpRequest, err = http.NewRequest(http.MethodGet, url, message)
		case "HEAD":
			httpRequest, err = http.NewRequest(http.MethodHead, url, message)
		case "POST":
			httpRequest, err = http.NewRequest(http.MethodPost, url, message)
		case "PUT":
			httpRequest, err = http.NewRequest(http.MethodPut, url, message)
		case "DELETE":
			httpRequest, err = http.NewRequest(http.MethodDelete, url, message)
		case "CONNECT":
			httpRequest, err = http.NewRequest(http.MethodConnect, url, message)
		case "PATCH":
			httpRequest, err = http.NewRequest(http.MethodPatch, url, message)
		case "OPTIONS":
			httpRequest, err = http.NewRequest(http.MethodOptions, url, message)
		case "TRACE":
			httpRequest, err = http.NewRequest(http.MethodTrace, url, message)
		default:
			log.Fatal("Undefined http method value found in\n" + stringRequest)
		}
		cobra.CheckErr(err)

		// NOTE: http version can only be set when acting as server
		if httpVersion != "" {
			if isFirstHttpVersionOccurrence {
				fmt.Println(Yellow + "NOTE: Http versions can't be set and, therefore, will be ignored!" + Reset)
				isFirstHttpVersionOccurrence = false
			}
			// httpRequest.Proto = httpVersion
		}

		if len(parsedHeaders) > 0 {
			httpRequest.Header = parsedHeaders
		}

		httpRequests = append(httpRequests, *httpRequest)
	}

	return httpRequests
}

func (p *Parser) ParseConfig() {
	var tempConfig map[string]interface{}
	// TODO: don't load whole files into RAM -> read by byte
	jsonData, err := os.ReadFile(p.ConfigPath)
	cobra.CheckErr(err)
	err = json.Unmarshal(jsonData, &tempConfig)
	cobra.CheckErr(err)

	envMap, configKeyExists := tempConfig[p.ConfigEnv]
	if !configKeyExists {
		log.Fatalf("%s does not exist in your config file!", p.ConfigEnv)
	}

	for key, val := range envMap.(map[string]interface{}) {
		typeOfVal := reflect.TypeOf(val).Kind().String()
		if typeOfVal != "string" && typeOfVal != "bool" {
			log.Fatalf("Your config contains key \"%s\" with value \"%s\" which is not of type string or bool!",
				key, val)
		}

		config[key] = val.(string)
	}

	LogVerbose(fmt.Sprintf("Using config environment: %s", p.ConfigEnv), p.Verbose)
}

// Takes string containing http requests and splits them by separators, removes comments and empty requests
// and inserts env variable values, if possible.
// Otherwise, exits program with exit code != 0.
func (p *Parser) Prepare(file string) []string {
	// Split file by request separators.
	// Exit, if no requests can be found after splitting.
	requests := regSeparator.Split(file, -1)
	if len(requests) == 0 {
		log.Fatal("No requests could be parsed!")
	}

	// Remove empty requests and comments
	noEmptyRequests := []string{}
	for i := range requests {
		if len(requests[i]) == 0 {
			continue
		}

		noEmptyRequests = append(noEmptyRequests, regComments.ReplaceAllLiteralString(requests[i], ""))
	}
	requests = noEmptyRequests

	isFirstResponseHandlerOccurrence := true
	isFirstResponseRefOccurrence := true
	for i := range requests {
		request := requests[i]

		// Insert env variable values from json config.
		// Exit, if one variable does not exist in json config.
		envMatches := regEnv.FindAllString(request, -1)

		if envMatches != nil {
			matchedConfig := p.matchConfig(envMatches)

			// TODO: use one by one comparison to ensure every match is represented in env variable config
			if len(envMatches) != len(matchedConfig) {
				log.Fatal("There are undefined env variables present in your requests file!\n",
					"Ensure your config variable keys use lowercase letters only.")
			}

			for configKey, configValue := range matchedConfig {
				request = strings.ReplaceAll(request, fmt.Sprintf("{{%s}}", configKey), configValue)
			}
		}

		// Remove response handler from request and prompt user that they are not going to be used
		responseHandlerMatches := regResponseHandler.FindAllString(request, -1)
		if len(responseHandlerMatches) > 1 {
			log.Fatal(`Some of your requests contain too many response handlers!\n
			Ensure there's only one handler per request.`)
		}

		// NOTE: Following TODOs are only relevant when adding response handler validation support
		// TODO: Trim '{%%}' from response handler match
		// TODO: Check if request separators or '%}' have been used inside the handler block

		if responseHandlerMatches != nil {
			if isFirstResponseHandlerOccurrence {
				fmt.Println(Yellow +
					"NOTE: Currently response handlers can't be validated and, therefore, will be ignored!" + Reset)
				isFirstResponseHandlerOccurrence = false
			}
			request = regResponseHandler.ReplaceAllLiteralString(request, "")
		}

		// Remove response ref from request and prompt user that they are not going to be used
		responseRefMatches := regResponseRef.FindAllString(request, -1)
		if len(responseRefMatches) > 1 {
			log.Fatal(`Some of your requests contain too many response refs!\n
			Ensure there is only one reference per request.`)
		}

		if responseRefMatches != nil {
			if isFirstResponseRefOccurrence {
				fmt.Println(Yellow +
					"NOTE: Currently response references can't be validated and, therefore, will be ignored!" + Reset)
				isFirstResponseRefOccurrence = false
			}
			request = regResponseRef.ReplaceAllLiteralString(request, "")
		}

		requests[i] = request
	}

	return requests
}

// Takes array of environment variable matches and compares them to keys stored in config.
// It returns an array of all config keys that have been matched.
func (p *Parser) matchConfig(envMatches []string) map[string]string {
	matchedConfig := make(map[string]string)

	for i := range envMatches {
		match := strings.Trim(envMatches[i], "{}")

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
// TODO: Outsource separating headers and messages from ParseHttpRequests when adding multipart/form-data support
// func separateHeadersFromMessage(headersMessage string) ([]string, string) {

// }

// Takes array of header strings splitted by new line, detects related header values and adds them to map
func (p *Parser) parseHeaders(headers []string) map[string][]string {
	parsedHeaders := make(map[string][]string)

	var fieldName string
	for i := range headers {
		headerMatch := strings.Trim(headers[i], "\r\n \t\f")
		if regHeaders.MatchString(headerMatch) {
			headerSubMatch := regHeaders.FindStringSubmatch(headerMatch)
			for i, name := range regHeaders.SubexpNames() {
				// NOTE: Start at index == 1 because first value of array is full match which is irrelevant
				if i >= 1 && i < len(headerSubMatch) {
					switch name {
					case "Fieldname":
						parsedHeaders[headerSubMatch[i]] = []string{}
						fieldName = headerSubMatch[i]
					case "Fieldvalue":
						parsedHeaders[fieldName] = append(parsedHeaders[fieldName], headerSubMatch[i])
					default:
						fmt.Println(Red + "Misconfigured header found!" + Reset)
						fmt.Println(Yellow + "Check header below for errors:" + Reset)
						fmt.Println(headerMatch)
						os.Exit(1)
					}
				}
			}
		} else {
			if len(parsedHeaders) > 0 && headerMatch != "" && fieldName != "" {
				parsedHeaders[fieldName] = append(parsedHeaders[fieldName], headerMatch)
			} else if headerMatch == "" {
				continue
			} else {
				fmt.Println(Red + "Misconfigured header found!" + Reset)
				fmt.Println(Yellow + "Check header below for errors:" + Reset)
				fmt.Println(headerMatch)
				os.Exit(1)
			}
		}
	}

	return parsedHeaders
}

// Takes message string, detects whether it is a filepath or direct message strings and returns them as io.Reader
func (p *Parser) parseMessage(message string) io.Reader {
	var file io.Reader
	var err error
	if regInputFileRef.MatchString(message) {
		message = strings.TrimSpace(strings.TrimPrefix(message, "<"))
		absMessagePath := ConvertToAbsolutePath(message, p.PathPrefix)
		file, err = os.Open(absMessagePath)
	} else {
		file, err = strings.NewReader(message), nil
	}

	cobra.CheckErr(err)

	return bufio.NewReader(file)
}
