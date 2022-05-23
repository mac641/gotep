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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mac641/gotep/src/lib/context"
	"github.com/mac641/gotep/src/lib/logger"
	"github.com/mac641/gotep/src/lib/parser"
	"github.com/mac641/gotep/src/lib/validator"
)

// runCmd represents the run command
var (
	ctx    = context.GetContext()
	log    = logger.GetLogger()
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
	ctx.SetVerbose(verbose)

	// Check if test file uses .rest or .http extension
	filePath := args[0]
	filePathExt := filepath.Ext(filePath)
	if filePathExt != ".rest" && filePathExt != ".http" {
		log.Fatal(`error: test file names need to have ".rest" or ".http" extensions`)
	}
	ctx.SetFilePath(filePath)
	log.Infof("using requests file: %s\n", filePath)

	pathPrefix := filepath.Dir(filePath) // Determine path prefix
	ctx.SetPathPrefix(pathPrefix)

	configFilePathBase := filepath.Base(configFilePath)
	if configFilePathBase != "http-client.env.json" && configFilePathBase != "http-client.private.env.json" {
		log.Fatal(`error: config files need to be named "http-client.env.json" or "http-client.private.env.json"`)
	}
	ctx.SetConfigFilePath(configFilePath)
	log.Infof("using config file: %s\n", configFilePath)

	ctx.SetConfigEnvironment(configEnvironment)
}

func test() {
	// TODO: don't load whole files into RAM -> read by byte
	content, err := os.ReadFile(ctx.GetFilePath())
	CheckErr(err)
	stringContent := string(content)

	// Create parser and parse test file
	parser := parser.Parser{}
	err = parser.ParseConfig()
	CheckErr(err)
	preparedRequests, err := parser.Prepare(stringContent)
	CheckErr(err)
	parsedHttpRequests, err := parser.Parse(preparedRequests)
	CheckErr(err)

	// Create validator and send requests
	validator := validator.Validator{}
	responses, err := validator.Send(parsedHttpRequests)
	CheckErr(err)
	succeeded := validator.Analyze(responses)

	// Print result to console
	os.Exit(log.Report(succeeded, len(responses)))
}
