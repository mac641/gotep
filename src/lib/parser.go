package lib

import (
	"fmt"
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
		// requestLineMatches := regexp.MustCompile("[ \t\f]").Split(strings.TrimRight(regRequestline.FindString(stringRequest), "\r\n"), -1)
		// method := ""
		// url := ""
		// httpVersion := ""
		// switch len(requestLineMatches) {
		// case 1:
		// 	url = requestLineMatches[0]
		// case 2:
		// 	if strings.Contains(requestLineMatches[1], "HTTP/") {
		// 		url = requestLineMatches[0]
		// 		httpVersion = requestLineMatches[1]
		// 	} else {
		// 		method = requestLineMatches[0]
		// 		url = requestLineMatches[1]
		// 	}
		// case 3:
		// 	method = requestLineMatches[0]
		// 	url = requestLineMatches[1]
		// 	httpVersion = requestLineMatches[2]
		// default:
		// 	log.Fatal("One of your request lines could not be parsed!")
		// 	fmt.Println(Yellow + "Check request below for errors:" + Reset)
		// 	fmt.Println(stringRequest)
		// 	os.Exit(1)
		// }
		stringHeaderMessage := strings.Trim(regRequestline.ReplaceAllLiteralString(stringRequest, ""), "\r\n")

		// if !IsUrlValid(url) {
		// 	log.Fatal(url + " is not a valid URL. Exiting...")
		// 	os.Exit(1)
		// }

		// Separate headers / message body and parse them afterwards
		// TODO: validate and reduce code duplication of case 1 and case 2
		splitEmptyNewline := regEmptyNewline.Split(stringHeaderMessage, -1)
		var parsedHeaders []string
		switch len(splitEmptyNewline) {
		case 0:
			if verbose {
				fmt.Println("No headers / message body have been provided for\\n" + stringRequest)
			}
			// TODO: create request directly in here
		case 1:
			match := splitEmptyNewline[0]
			if regHeaders.MatchString(match) {
				if verbose {
					fmt.Println("No message has been provided for\n" + stringRequest)
				}

				headers := regLineEnding.Split(match, -1)
				// fmt.Println(headerSplitNewline)
				parsedHeaders = parseHeaders(headers, *regHeaders)
				if len(parsedHeaders) <= 0 {
					log.Fatal("Misconfigured headers found!")
					fmt.Println(Yellow + "Check request below for errors:" + Reset)
					fmt.Println(stringRequest)
					os.Exit(1)
				}
				// TODO: create request directly in here
			} else {
				if verbose {
					fmt.Println("No headers have been provided for\n" + stringRequest)
				}

				if regInputFileRef.MatchString(match) {
					fmt.Println(match)
					// TODO: use io.Reader to create request directly in here
				} else {
					fmt.Println(match)
					// TODO: use message to create request directly in here; maybe write to message to temp file and use io.Reader, too
				}
			}
		case 2:
			message := splitEmptyNewline[1]

			headers := regLineEnding.Split(splitEmptyNewline[0], -1)
			// fmt.Println(headerSplitNewline)
			parsedHeaders = parseHeaders(headers, *regHeaders)
			if len(parsedHeaders) <= 0 {
				log.Fatal("Misconfigured headers found!")
				fmt.Println(Yellow + "Check request below for errors:" + Reset)
				fmt.Println(stringRequest)
				os.Exit(1)
			}

			if regInputFileRef.MatchString(message) {
				fmt.Println(message)
				// TODO: validate file path
				// TODO: use io.Reader to create request directly in here
			} else {
				fmt.Println(message)
				// TODO: use message to create request directly in here; maybe write to message to temp file and use io.Reader, too
			}
		default:
			log.Fatal("Too many in-between line breaks detected!")
			fmt.Println(Yellow + "Check request below for errors:" + Reset)
			fmt.Println(stringRequest)
			os.Exit(1)
		}
		// TODO: extract message body and add them to requests
	}

	return resultRequests
}

// Takes string containing http requests and splits them by separators, removes comments and empty requests
// and inserts env variable values, if possible.
// Otherwise, exits program with exit code != 0
func PrepareHttpRequests(file string, conEnv string) []string {
	// Split file by request separators.
	// Exit, if no requests can be found after splitting.
	result := regSeparator.Split(file, -1)
	if len(result) == 0 {
		log.Fatal("No requests could be parsed!")
		os.Exit(1)
	}

	// Remove empty requests and comments
	removedComments := []string{}
	for i := range result {
		if len(result[i]) == 0 {
			continue
		}

		removed := regComments.ReplaceAllLiteralString(result[i], "")
		removedComments = append(removedComments, removed)
	}
	result = removedComments

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
				os.Exit(1)
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
			os.Exit(1)
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

	// TODO: remove response ref and prompt user that they're going to be ignored - instead every response code 200 will be treated as successful
	totalResponseRefCount := 0
	for i := range result {
		request := result[i]
		responseRefMatches := regResponseRef.FindAllString(request, -1)

		if len(responseRefMatches) > 1 {
			log.Fatal(`Some of your requests contain too many response refs!\n
			Ensure there is only one reference per request.`)
			os.Exit(1)
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
func parseHeaders(s []string, regHeaders regexp.Regexp) []string {
	result := []string{}

	for j := range s {
		headerMatch := strings.Trim(s[j], "\r\n \t\f")
		if regHeaders.MatchString(headerMatch) {
			result = append(result, headerMatch)
		} else {
			if len(result) > 0 && headerMatch != "" {
				result[len(result)-1] += ", " + headerMatch
			} else if headerMatch == "" {
				continue
			} else {

			}
		}
	}

	return result
}
