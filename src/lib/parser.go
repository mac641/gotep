package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

// Takes array of strings representing prepared http requests and parses them into array of http.Requests
func ParseHttpRequests(requests []string, verbose bool) []http.Request {
	resultRequests := []http.Request{}
	for i := range requests {
		stringRequest := requests[i]

		// Parse request line
		requestLineMatches := regexp.MustCompile("[ \t\f]").Split(strings.TrimRight(regRequestline.FindString(stringRequest), "\r\n"), -1)
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
		stringHeaderMessage := strings.Trim(regRequestline.ReplaceAllLiteralString(stringRequest, ""), "\r\n")

		if !IsUrlValid(url) {
			log.Fatal(url + " is not a valid URL. Exiting...")
		}

		// Separate headers / message body and parse them afterwards
		splitEmptyNewline := regEmptyNewline.Split(stringHeaderMessage, -1)
		parsedHeaders := []string{}
		var message io.Reader
		switch len(splitEmptyNewline) {
		case 0:
			if verbose {
				fmt.Println("Neither headers, nor a message body have/has been provided for\\n" + stringRequest)
			}
		case 1:
			match := splitEmptyNewline[0]
			if regHeaders.MatchString(match) {
				if verbose {
					fmt.Println("No message body has been provided for\n" + stringRequest)
				}

				parsedHeaders = parseHeaders(regLineEnding.Split(splitEmptyNewline[0], -1))

			} else {
				if verbose {
					fmt.Println("No headers have been provided for\n" + stringRequest)
				}

				message = parseMessage(match)
			}
		case 2:
			if verbose {
				fmt.Println("Headers and message body detected in\n" + stringRequest)
			}

			parsedHeaders = parseHeaders(regLineEnding.Split(splitEmptyNewline[0], -1))
			message = parseMessage(splitEmptyNewline[1])

		default:
			fmt.Println(Red + "Too many in-between line breaks detected!" + Reset)
			fmt.Println(Yellow + "Check request below for errors:" + Reset)
			fmt.Println(stringRequest)
		}

		// TODO: use message and parsedHeaders to assemble request
	}

	return resultRequests
}

// Takes string containing http requests and splits them by separators, removes comments and empty requests
// and inserts env variable values, if possible.
// Otherwise, exits program with exit code != 0.
func PrepareHttpRequests(file string, conEnv string) []string {
	// Split file by request separators.
	// Exit, if no requests can be found after splitting.
	result := regSeparator.Split(file, -1)
	if len(result) == 0 {
		log.Fatal("No requests could be parsed!")
		os.Exit(1)
	}

	// Remove empty requests and comments
	for i := range result {
		if len(result[i]) == 0 {
			continue
		}

		result[i] = regComments.ReplaceAllLiteralString(result[i], "")
	}

	// Insert env variable values from json config.
	// Exit, if one variable does not exist in json config.
	allConfigKeys := viper.AllKeys()
	for i := range result {
		env := result[i]
		envMatches := regEnv.FindAllString(env, -1)

		if envMatches != nil {
			matchedConfigKeys := []string{}

			for j := range envMatches {
				match := strings.Trim(envMatches[j], "{}")

				for k := range allConfigKeys {
					if strings.Contains(allConfigKeys[k], match) && strings.Contains(allConfigKeys[k], conEnv) {
						matchedConfigKeys = append(matchedConfigKeys, allConfigKeys[k])
					}
				}
			}

			// TODO: use one by one comparison to ensure every match is represented in env variable config
			if len(envMatches) != len(matchedConfigKeys) {
				log.Fatal("There are undefined env variables present in your requests file!")
			}

			inserted := env
			for j := range matchedConfigKeys {
				configValue := viper.GetString(matchedConfigKeys[j])
				inserted = strings.ReplaceAll(inserted, "{{"+strings.TrimPrefix(matchedConfigKeys[j], conEnv+".")+"}}",
					configValue)
			}
			result[i] = inserted
		}
	}

	// Remove responsehandler from requests and prompt user that they are not going to be used
	totalResponseHandlerCount := 0
	for i := range result {
		request := result[i]
		responseHandlerMatches := regResponseHandler.FindAllString(request, -1)

		if len(responseHandlerMatches) > 1 {
			log.Fatal(`Some of your requests contain too many response handlers!\n
			Ensure there's only one handler per request.`)
		}

		if responseHandlerMatches != nil {
			if totalResponseHandlerCount == 0 {
				fmt.Println(Yellow + "NOTE: Currently response handlers can't be validated and, therefore, will be ignored!" + Reset)
			}
			totalResponseHandlerCount++
			result[i] = regResponseHandler.ReplaceAllLiteralString(request, "")

			// TODO: enhance response handler regex to be able to determine, if '###' and '%}' literals have been used inside the handler block
		}
	}

	totalResponseRefCount := 0
	for i := range result {
		request := result[i]
		responseRefMatches := regResponseRef.FindAllString(request, -1)

		if len(responseRefMatches) > 1 {
			log.Fatal(`Some of your requests contain too many response refs!\n
			Ensure there is only one reference per request.`)
		}

		if responseRefMatches != nil {
			if totalResponseRefCount == 0 {
				fmt.Println(Yellow + "NOTE: Currently response references can't be validated and, therefore, will be ignored!" + Reset)
			}
			totalResponseRefCount++
			result[i] = regResponseRef.ReplaceAllLiteralString(request, "")
		}
	}

	return result
}

// Takes array of header strings splitted by new line, detects related header values and concats them
func parseHeaders(headers []string) []string {
	result := []string{}

	for j := range headers {
		headerMatch := strings.Trim(headers[j], "\r\n \t\f")
		if regHeaders.MatchString(headerMatch) {
			result = append(result, headerMatch)
		} else {
			if len(result) > 0 && headerMatch != "" {
				result[len(result)-1] += ", " + headerMatch
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

	return result
}

// Takes message string, detects whether it is a filepath or direct message strings and wraps them as io.Reader
func parseMessage(message string) io.Reader {
	var file io.Reader
	var err error
	if regInputFileRef.MatchString(message) {
		file, err = os.Open(message)
	} else {
		file, err = strings.NewReader(message), nil
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return bufio.NewReader(file)
}
