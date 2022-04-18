/*
Copyright © 2022 Marcel Arndt <kontakt@marcelarndt.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute tests",
		Long: `Execute tests based on a test file and its associated environment definitions.
For example run: gotep run -c http-client.env.json -f tests.http`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("run called")
			initConfig(cmd)
			test(cmd)
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig(cmd *cobra.Command) {
	config, err := cmd.Flags().GetString(config)
	cobra.CheckErr(err)
	if config != "" && err == nil {
		viper.SetConfigFile(config)
	} else {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(cwd)
		viper.SetConfigType("json")
		viper.SetConfigName("http-client.env")
	}

	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func test(cmd *cobra.Command) {
	tests, err := cmd.Flags().GetString(file)
	cobra.CheckErr(err)
	if tests == "" {
		tests, err := os.Getwd()
		cobra.CheckErr(err)
		tests += "default.http"
	}
	if verbose {
		fmt.Println("Using requests file:", tests)
	}

	content, err := ioutil.ReadFile(tests)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: invoke parsing of test_files here
	// Parse test file
	stringContent := string(content)
	conEnv, err := cmd.Flags().GetString(configEnv)
	cobra.CheckErr(err)
	if verbose {
		fmt.Println("Using environment config:", conEnv)
	}

	preparedRequests := prepareHttpRequests(stringContent, conEnv)
	parseHttpRequests(preparedRequests)
	// for i := range preparedRequests {
	// 	fmt.Print(preparedRequests[i])
	// 	fmt.Println("--------------------------------------------")
	// }

	// TODO: replace file refs with their contents and panic if they don't exist

	// TODO: use http.ReadRequest to read and send requests after preparing them

	// TODO: use http.ReadResponse to compare possible response refs
}

// Takes array of strings representing prepared http requests and parses them into array of http.Requests
func parseHttpRequests(requests []string) []http.Request {
	regRequestline := regexp.MustCompile("(?m)^((GET|HEAD|POST|PUT|DELETE|CONNECT|PATCH|OPTIONS|TRACE)[ \t\f]+)?(([0-9/\\[]|http|https)[^ \t\f]+)([ \t\f]+(HTTP/\\d+(\\.\\d+)?))?\r?\n|\r$")
	// regHeaders := regexp.MustCompile("(?m)^([^:]+):[ \t\f]*(([^\r\n]+((\r?\n|\r)($^[ \t\f]+)?))+)[ \t\f]*$")
	// regBody :=

	resultRequests := []http.Request{}
	for i := range requests {
		stringRequest := requests[i]
		requestLineMatches := strings.Split(regRequestline.FindString(stringRequest), " ")

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
			// FIXME: prevent httpVersion from getting assigned to single new line
			httpVersion = requestLineMatches[2]
		default:
			log.Fatal("One of your request lines could not be parsed!")
			fmt.Println(Yellow + "Check request below for errors:" + Reset)
			fmt.Println(stringRequest)
			os.Exit(1)
		}
		stringRequest = regRequestline.ReplaceAllLiteralString(stringRequest, "")
		fmt.Println(stringRequest+",", method+",", url+",", httpVersion)

		if !isUrlValid(url) {
			log.Fatal(url, "is not a valid URL. Exiting...")
			os.Exit(1)
		}
		// switch method {
		// case "", "GET":
		// 	req, err := http.NewRequest(http.MethodGet, url)
		// 	cobra.CheckErr(err)
		// }

	}

	return resultRequests
}

// Takes string containing http requests and splits them by separators, removes comments and empty requests
// and inserts env variable values, if possible.
// Otherwise, exits program with exit code != 0
func prepareHttpRequests(file string, conEnv string) []string {
	// Split file by request separators.
	// Exit, if no requests can be found after splitting.
	// NOTE: enable multi-line mode flag (?m) to match all occurrences in file string
	regSeparator := regexp.MustCompile("(?m)^(###[^\u000D\u000A]*(\u000D?\u000A|\u000D))") // \r\n
	result := regSeparator.Split(file, -1)
	if len(result) == 0 {
		log.Fatal("No requests could be parsed!")
		os.Exit(1)
	}

	// Remove empty requests and comments
	// NOTE: enable multi-line mode flag (?m) to match all occurrences in file string
	regComments := regexp.MustCompile("(?m)^(//|#)([^\u000D\u000A]*(\u000D?\u000A|\u000D))") // \r\n
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
	regEnv := regexp.MustCompile("{{[ \u0009\u000C]*[A-Za-z0-9\\-_]+[ \u0009\u000C]*}}") // space\t\f
	allConfigKeys := viper.AllKeys()
	for i := range result {
		env := result[i]
		envMatches := regEnv.FindAllString(env, -1)

		if envMatches != nil {
			matchedConfigKeys := []string{}

			for j := range envMatches {
				match := strings.TrimLeft(envMatches[j], "{")
				match = strings.TrimRight(match, "}")

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
	// NOTE: enable multi-line mode flag (?m) to match all occurrences in file string
	regResponseHandler := regexp.MustCompile("(?m)^>[ \t\f]+([^\r\n]*(\u000D?\u000A|\u000D)$|{%(.|(\u000D?\u000A|\u000D))+%})(\u000D?\u000A|\u000D)")
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

	return result
}

func isUrlValid(s string) bool {
	// TODO: use uri.ParseRequestURI(s) to parse urls and afterwards use net.ParseIP(s) to parse IP if needed

	return true
}
