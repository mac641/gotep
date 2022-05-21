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
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mac641/gotep/src/lib"
)

// runCmd represents the run command
var (
	ctx    = lib.GetContext()
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute tests",
		Long: `Execute tests based on a test file and its associated environment definitions.
For example run: gotep run -c http-client.env.json -e default tests.http`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// err := cmd.ParseFlags(args)
			// if err != nil {
			// 	log.Fatalf(err.Error())
			// }
			assignContext(cmd, args)
			test()
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

func assignContext(cmd *cobra.Command, args []string) {
	// Check if test file uses .rest or .http extension
	fileName := args[0]
	testFileSplit := strings.Split(fileName, ".")
	testFileFormat := strings.ToLower(testFileSplit[len(testFileSplit)-1])
	if testFileFormat != "rest" && testFileFormat != "http" {
		log.Fatalf(`Test file names need to have ".rest" or ".http" extensions.`)
	}
	ctx.SetFilePath(fileName)
	lib.LogVerbose(fmt.Sprintf("Using requests file: %s", fileName), verbose)

	pathPrefix := filepath.Dir(fileName) // Determine path prefix
	ctx.SetPathPrefix(pathPrefix)

	// Check if environment config is not empty and use it, if so
	ctx.SetConfigFilePath(configFilePath)
	lib.LogVerbose(fmt.Sprintf("Using config file: %s", configFilePath), verbose)

	ctx.SetConfigEnvironment(configEnvironment)

	ctx.SetVerbose(verbose)
}

func test() {
	// TODO: don't load whole files into RAM -> read by byte
	content, err := os.ReadFile(ctx.GetFilePath())
	cobra.CheckErr(err)
	stringContent := string(content)

	// Create parser and parse test file
	parser := lib.Parser{}
	parser.ParseConfig()
	preparedRequests := parser.Prepare(stringContent)
	parsedHttpRequests := parser.Parse(preparedRequests)

	// Create validator and send requests
	validator := lib.Validator{}
	validator.Send(parsedHttpRequests)
}
