package context_test

import (
	"reflect"
	"testing"

	"github.com/mac641/gotep/src/lib/context"
)

var ctx = context.GetContext()

func TestGetContext(t *testing.T) {
	context1 := context.GetContext()
	context2 := context.GetContext()

	if !reflect.DeepEqual(context1, context2) {
		t.Errorf("Expected context to exist only once but got different instances returned")
	}
}

func TestGetConfigEnvironment(t *testing.T) {
	configEnvironment := ctx.GetConfigEnvironment()
	configEnvironmentAssert := "default"

	if configEnvironment != configEnvironmentAssert {
		t.Errorf("Expected configEnvironment to be %s but got %s", configEnvironmentAssert, configEnvironment)
	}
}

func TestGetConfigFilePath(t *testing.T) {
	configFilePath := ctx.GetConfigFilePath()
	configFilePathAssert := "http-client.env.json"

	if configFilePath != configFilePathAssert {
		t.Errorf("Expected configFilePath to be %s but got %s", configFilePathAssert, configFilePath)
	}
}

func TestGetFilePath(t *testing.T) {
	filePath := ctx.GetFilePath()
	filePathAssert := "default.http"

	if filePath != filePathAssert {
		t.Errorf("Expected filePath to be %s but got %s", filePathAssert, filePath)
	}
}

func TestGetPathPrefix(t *testing.T) {
	pathPrefix := ctx.GetPathPrefix()
	pathPrefixAssert := "."

	if pathPrefix != pathPrefixAssert {
		t.Errorf("Expected pathPrefix to be %s but got %s", pathPrefixAssert, pathPrefix)
	}
}

func TestGetVerbose(t *testing.T) {
	verbose := ctx.GetVerbose()
	verboseAssert := false

	if verbose != verboseAssert {
		t.Errorf("Expected verbose to be %t but got %t", verboseAssert, verbose)
	}
}

func TestSetConfigEnvironment(t *testing.T) {
	configEnvironment := "test"
	ctx.SetConfigEnvironment(configEnvironment)

	if configEnvironment != ctx.GetConfigEnvironment() {
		t.Errorf("Expected configEnvironment to be %s but got %s", configEnvironment, ctx.GetConfigEnvironment())
	}
}

func TestSetConfigFilePath(t *testing.T) {
	configFilePath := "http-client.private.env.json"
	ctx.SetConfigFilePath(configFilePath)

	if configFilePath != ctx.GetConfigFilePath() {
		t.Errorf("Expected configFilePath to be %s but got %s", configFilePath, ctx.GetConfigFilePath())
	}
}

func TestSetFilePath(t *testing.T) {
	filePath := "test.rest"
	ctx.SetFilePath(filePath)

	if filePath != ctx.GetFilePath() {
		t.Errorf("Expected filePath to be %s but got %s", filePath, ctx.GetFilePath())
	}
}

func TestSetPathPrefix(t *testing.T) {
	pathPrefix := ".."
	ctx.SetPathPrefix(pathPrefix)

	if pathPrefix != ctx.GetPathPrefix() {
		t.Errorf("Expected pathPrefix to be %s but got %s", pathPrefix, ctx.GetPathPrefix())
	}
}

func TestSetVerbose(t *testing.T) {
	verbose := true
	ctx.SetVerbose(verbose)

	if verbose != ctx.GetVerbose() {
		t.Errorf("Expected verbose to be %t but got %t", verbose, ctx.GetVerbose())
	}
}
