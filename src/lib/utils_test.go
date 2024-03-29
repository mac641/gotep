package lib_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/mac641/gotep/src/lib"
	"github.com/mac641/gotep/src/lib/context"
)

var (
	ctx = context.GetContext()
)

func init() {
	if runtime.GOOS == "windows" {
		// NOTE: use D disk for compatibility with GitHub pipeline
		ctx.SetPathPrefix("D:\\home\\testuser")
	} else {
		ctx.SetPathPrefix("/home/testuser")
	}
}

func TestConvertToAbsolutePath(t *testing.T) {
	var (
		p = fmt.Sprintf("a%crelative%cpath", os.PathSeparator, os.PathSeparator)
	)

	result, err := lib.ConvertToAbsolutePath(ctx.GetPathPrefix())
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != ctx.GetPathPrefix() {
		t.Errorf("Expected absolute path to be %s without editing but got %s",
			ctx.GetPathPrefix(), result)
	}

	result, err = lib.ConvertToAbsolutePath(p)
	if err != nil {
		t.Errorf(err.Error())
	}
	assertValueJoin := filepath.Join(ctx.GetPathPrefix(), p)
	if result != assertValueJoin {
		t.Errorf("Expected absolute paths to be joined properly like %s but got %s",
			assertValueJoin, result)
	}

	ctx.SetPathPrefix("")
	result, err = lib.ConvertToAbsolutePath(p)
	if err != nil {
		t.Errorf(err.Error())
	}
	assertValueAbs, err := filepath.Abs(p)
	if err != nil {
		t.Errorf("Error while generating absolute path. Error: %s", err.Error())
	}
	if result != assertValueAbs {
		t.Errorf("Absolute paths in the same directory have to use matching path prefixes! Expected: %s, Got: %s",
			assertValueAbs, result)
	}
}

func TestIsUrlValid(t *testing.T) {
	var (
		validUrls = []string{
			"https://192.168.178.2/index/https",
			"192.168.178.2/index/https",
			"http://www.example.com/test1",
			"[1111:2222:abcd:4444:5a6c:7777:9c9f:6789]/index",
			"/api/get",
			"*",
			"example.com/test?id=1",
			"example.com:80",
		}
		invalidUrls = []string{
			"httphallt",
			"",
			"a/relative/path",
			"htt://example.com",
			"http//example.com",
		}
	)

	for _, url := range validUrls {
		valid := lib.IsUrlValid(url)
		if !valid {
			t.Errorf("Expected %s to be valid url.", url)
		}
	}

	for _, url := range invalidUrls {
		invalid := lib.IsUrlValid(url)
		if invalid {
			t.Errorf("Expected %s to be invalid url.", url)
		}
	}
}

func TestTrimEmptyStringsFromSlice(t *testing.T) {
	testSlice := []string{"", "", "", "hello", "i'm in between", ""}
	assertSlice := []string{"hello", "i'm in between"}

	testResult := lib.TrimEmptyStringsFromSlice(testSlice)
	if !reflect.DeepEqual(testResult, assertSlice) {
		t.Errorf("Expected result to be %s but got %s", assertSlice, testResult)
	}
}

func TestTrimLeftEmptyStringsFromSlice(t *testing.T) {
	testSlice := []string{"", "", "", "hello", "i'm in between"}
	assertSlice := []string{"hello", "i'm in between"}

	testResult := lib.TrimLeftEmptyStringsFromSlice(testSlice)
	if !reflect.DeepEqual(testResult, assertSlice) {
		t.Errorf("Expected result to be %s but got %s", assertSlice, testResult)
	}
}

func TestTrimRightEmptyStringsFromSlice(t *testing.T) {
	testSlice := []string{"hello", "i'm in between", "", "", ""}
	assertSlice := []string{"hello", "i'm in between"}

	testResult := lib.TrimRightEmptyStringsFromSlice(testSlice)
	if !reflect.DeepEqual(testResult, assertSlice) {
		t.Errorf("Expected result to be %s but got %s", assertSlice, testResult)
	}
}
