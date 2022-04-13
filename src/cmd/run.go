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

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mac641/gotep/src/lib/parser"
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
	if config != "" {
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

// Define httpspec listener based on generated antlr4 source code
type TreeShapeListener struct {
	*parser.BasehttpSpecListener
}

// func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	const separator = ","
// 	fmt.Println(ctx.GetStart(), separator, ctx.GetStop())
// 	fmt.Println(ctx.GetText())
// 	fmt.Println()
// }

func (tsl *TreeShapeListener) ExitRequest(ctx parser.RequestContext) {
	fmt.Println(ctx.GetText())
}

func test(cmd *cobra.Command) {
	tests, err := cmd.Flags().GetString(file)
	cobra.CheckErr(err)
	if tests == "" {
		tests, err := os.Getwd()
		cobra.CheckErr(err)
		tests += "default.http"
	}
	// content, err := ioutil.ReadFile(tests)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Parse test file
	// stringContent := string(content)
	charStream, _ := antlr.NewFileStream(tests)
	lexer := parser.NewhttpSpecLexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewhttpSpecParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(false)) // TODO: handle parsing errors
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(&TreeShapeListener{}, p.File())
}
