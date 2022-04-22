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

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	testFile          string
	configFile        string
	configEnvironment string
	verbose           bool

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

const (
	config    = "config"
	configEnv = "configEnv"
	file      = "file"
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

	// cobra.OnInitialize(initConfig, initTests)

	// TODO: implement custom print queue for better formatted output
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"print log messages (default is false)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	checkCmd.Flags().StringVarP(&configFile, config, "c", "",
		"config file path (default is current working directory + \"http-client.env.json\")")
	checkCmd.Flags().StringVarP(&testFile, file, "f", "",
		"test file path (default is current working directory + \"default.http\")")
	runCmd.Flags().StringVarP(&configFile, config, "c", "",
		"config file path (default is current working directory + \"http-client.env.json\")")
	runCmd.Flags().StringVarP(&configEnvironment, configEnv, "e", "default",
		"environment name specified in config file (default is 'default')")
	runCmd.Flags().StringVarP(&testFile, file, "f", "",
		"test file path (default is current working directory + \"default.http\")")
}
