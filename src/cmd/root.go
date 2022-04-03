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
	"github.com/mac641/gotep/src/lib/parser"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Define httpspec listener based on generated antlr4 source code
type TreeShapeListener struct {
	*parser.BasehttpSpecListener

	*parser.FileContext
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	const separator = ","
	fmt.Println(ctx.GetStart(), separator, ctx.GetStop())
	fmt.Println(ctx.GetText())
	fmt.Println()
}

// func (tsl *TreeShapeListener) ExitFile(ctx parser.FileContext) {
// 	fmt.Println(ctx.GetText())
// }

// rootCmd represents the base command when called without any subcommands
var (
	testFile   string
	configFile string
	verbose    bool

	rootCmd = &cobra.Command{
		Use:   "gotep",
		Short: "gotep is a terminal-based REST client.",
		Long: `gotep is a terminal-based REST client designed to execute HTTP tests based on the Jetbrains
HTTP-Client.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cobra.OnInitialize(initConfig, initTests)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"print log messages (default is false)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	checkCmd.Flags().StringVarP(&configFile, "config", "c", "",
		"config file path (default is current working directory + \"http-client.env.json\")")
	checkCmd.Flags().StringVarP(&testFile, "file", "f", "",
		"test file path (default is current working directory + \"default.http\")")
	runCmd.Flags().StringVarP(&configFile, "config", "c", "",
		"config file path (default is current working directory + \"http-client.env.json\")")
	runCmd.Flags().StringVarP(&testFile, "file", "f", "",
		"test file path (default is current working directory + \"default.http\")")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
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

func initTests() {
	if testFile == "" {
		testFile, err := os.Getwd()
		cobra.CheckErr(err)
		testFile += "default.http"
	}
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: invoke parsing of test_files here
	// Parse test file
	stringContent := string(content)
	charStream := antlr.NewInputStream(stringContent)
	lexer := parser.NewhttpSpecLexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewhttpSpecParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.File()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	// TODO: validate if check has been called - maybe outsource it to check.go

	// fmt.Println(stringContent)
}
