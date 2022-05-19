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
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mac641/gotep/src/lib"
)

// runCmd represents the run command
var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute tests",
		Long: `Execute tests based on a test file and its associated environment definitions.
For example run: gotep run -c http-client.env.json -f tests.http -e default`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("run called")
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

func test(cmd *cobra.Command) {
	// Check if test file has been passed and uses .rest or .http extension
	tests, err := cmd.Flags().GetString(file)
	testFileSplit := strings.Split(tests, ".")
	testFileFormat := strings.ToLower(testFileSplit[len(testFileSplit)-1])
	if testFileFormat != "rest" && testFileFormat != "http" {
		log.Fatalf(`Test file names need to have ".rest" or ".http" extensions.`)
	}
	pathPrefix := filepath.Dir(tests)
	cobra.CheckErr(err)
	if tests == "" {
		tests, err := os.Getwd()
		cobra.CheckErr(err)
		pathPrefix = tests

		tests += "default.http"
	}
	lib.LogVerbose(fmt.Sprintf("Using requests file: %s", tests), verbose)

	// TODO: don't load whole files into RAM -> read by byte
	content, err := os.ReadFile(tests)
	cobra.CheckErr(err)
	stringContent := string(content)

	// Check if environment config is not empty and use it, if so
	config, err := cmd.Flags().GetString(config)
	cobra.CheckErr(err)
	var configPath string
	if config != "" {
		configPath = config
	} else {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		configPath = path.Join(cwd, "http-client.env.json")
	}
	conEnv, err := cmd.Flags().GetString(configEnv)
	cobra.CheckErr(err)

	// Create parser
	parser := lib.Parser{
		ConfigEnv:  conEnv,
		ConfigPath: configPath,
		PathPrefix: pathPrefix,
		Verbose:    verbose,
	}

	// Parse test file
	parser.ParseConfig()
	preparedRequests := parser.Prepare(stringContent)
	// fmt.Println(preparedRequests)
	parsedHttpRequests := parser.Parse(preparedRequests)
	// fmt.Println(parsedHttpRequests)

	// TODO: collect requests and print ALL results on console, rather than exiting at every failed request
	responses := []*http.Response{}
	client := &http.Client{}
	for i := range parsedHttpRequests {
		r, err := client.Do(&parsedHttpRequests[i])
		cobra.CheckErr(err)
		responses = append(responses, r)
		fmt.Println(r.Status)
	}
	fmt.Println(responses)
}
