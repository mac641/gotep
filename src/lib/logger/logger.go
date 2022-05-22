package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/mac641/gotep/src/lib/context"
)

type logger struct{}

var (
	l *logger

	ctx  = context.GetContext()
	once = sync.Once{}

	// Print functions in blue, red and yellow
	blue     = color.New(color.FgBlue).PrintFunc()
	bluef    = color.New(color.FgBlue).PrintfFunc()
	blueln   = color.New(color.FgBlue).PrintlnFunc()
	red      = color.New(color.FgRed).PrintFunc()
	redf     = color.New(color.FgRed).PrintfFunc()
	redln    = color.New(color.FgRed).PrintlnFunc()
	yellow   = color.New(color.FgYellow).PrintFunc()
	yellowf  = color.New(color.FgYellow).PrintfFunc()
	yellowln = color.New(color.FgYellow).PrintlnFunc()
)

func GetLogger() *logger {
	once.Do(func() {
		l = new(logger)
	})
	return l
}

// Logs message colored red to console but does not exit app
func (l *logger) Error(msg string) {
	red(msg)
}

// Logs format colored red to console using given array of arguments but does not exit app
func (l *logger) Errorf(format string, a ...any) {
	redf(format, a...)
}

// Logs message colored red to console and appends new line but does not exit app
func (l *logger) Errorln(msg string) {
	redln(msg)
}

// Uses Error to log message and exits the app with exit code 1
func (l *logger) Fatal(msg string) {
	l.Error(msg)
	os.Exit(1) // TODO: consider setting exit code via parameter
}

// Uses Errorf to log format using given arguments and exits the app with exit code 1
func (l *logger) Fatalf(format string, a ...any) {
	l.Errorf(format, a...)
	os.Exit(1) // TODO: consider setting exit code via parameter
}

// Logs message colored blue, if verbose flag has been set
func (l *logger) Info(msg string) {
	if ctx.GetVerbose() {
		blue(msg)
	}
}

// Logs formate colored blue to console using given arguments, if verbose flag has been set
func (l *logger) Infof(format string, a ...any) {
	if ctx.GetVerbose() {
		bluef(format, a...)
	}
}

// Logs message colored blue to console and appends new line, if verbose flag has been set
func (l *logger) Infoln(msg string) {
	if ctx.GetVerbose() {
		blueln(msg)
	}
}

// Generates report based on analyzed http responses, prints it to console and returns app exit code
func (l *logger) Report(succeeded int, total int) (exitCode int) {
	fmt.Println()
	fmt.Println("All tests have been run.")
	fmt.Println("----------------------------------------")
	fmt.Println()
	fmt.Println("Results:")
	if succeeded < total {
		color.Set(color.FgRed, color.Bold)
		exitCode = 1
	} else {
		color.Set(color.FgGreen, color.Bold)
		exitCode = 0
	}
	fmt.Printf("Successful %d/%d\n", succeeded, total)
	color.Unset()
	fmt.Println()

	return exitCode
}

// Logs message colored yellow to console
func (l *logger) Warn(msg string) {
	yellow(msg)
}

// Logs format colored yellow to console using given arguments
func (l *logger) Warnf(format string, a ...any) {
	yellowf(format, a...)
}

// Logs message colored yellow to console and appends new line
func (l *logger) Warnln(msg string) {
	yellowln(msg)
}
