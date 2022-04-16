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
	"os"
	"regexp"

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

	if err := viper.ReadInConfig(); err == nil {
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
	content, err := ioutil.ReadFile(tests)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: invoke parsing of test_files here
	// Parse test file
	stringContent := string(content)
	requests := readHttpRequests(stringContent)
	for i := range requests {
		fmt.Print(requests[i])
		fmt.Println("--------------------------------------------")
	}

	// TODO: replace file refs with their contents and panic if they don't exist
}

// Reads http requests from given string, splits them by separators, removes comments and
// inserts env variables if possible, otherwise returns error
func readHttpRequests(file string) []string {
	// NOTE: enable multi-line mode flag (?m) to match all occurrences in string
	regSeparator := regexp.MustCompile("(?m)^(###[^\u000D\u000A]*\u000D?\u000A)") // \r\n
	splitSeparator := regSeparator.Split(file, -1)

	regComments := regexp.MustCompile("(?m)^(//|#)([^\u000D\u000A]*\u000D?\u000A)") // \r\n
	removedComments := []string{}
	for i := range splitSeparator {
		if len(splitSeparator[i]) == 0 {
			continue
		}

		removed := regComments.ReplaceAllLiteralString(splitSeparator[i], "")
		removedComments = append(removedComments, removed)
	}

	// TODO: add env var support, panic if one does not exist
	// regEnv := regexp.MustCompile("{{[ \u0009\u000C]*[A-Za-z0-9\\-_]+[ \u0009\u000C]*}}") // space\t\f
	// insertEnv := []string{}
	// allConfigKeys := viper.AllKeys()
	// for i := range removedComments {
	// 	env := removedComments[i]
	// 	envMatches := regEnv.FindAllString(env, -1)
	// 	if envMatches != nil {

	// 		matchedConfigKeys := []string{}
	// 		x := 0
	// 		for j := range envMatches {
	// 			match := strings.TrimLeft(envMatches[j], "{")
	// 			match = strings.TrimRight(match, "}")
	// 			x := sort.Search(len(allConfigKeys), func(i int) bool { return allConfigKeys[x] == match })
	// 			if x < len(allConfigKeys) {
	// 				matchedConfigKeys = append(matchedConfigKeys, allConfigKeys[x])
	// 			}
	// 		}

	// 		for j := range matchedConfigKeys {
	// 			insertEnv = append(insertEnv, regEnv.ReplaceAllString(env, matchedConfigKeys[j]))
	// 		}
	// 	}

	// }

	return removedComments
}
