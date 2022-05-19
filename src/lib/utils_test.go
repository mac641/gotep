package lib_test

import (
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/mac641/gotep/src/lib"
)

func TestConvertToAbsolutePath(t *testing.T) {
	const (
		prefix = "/home/testuser"
		p      = "a/relative/path"
	)

	result := lib.ConvertToAbsolutePath(prefix, prefix)
	if result != prefix {
		t.Errorf("Expect absolute path to be returned without editing! Expected: %s, Got: %s", prefix, result)
	}

	result = lib.ConvertToAbsolutePath(p, prefix)
	assertValueJoin := path.Join(prefix, p)
	if result != assertValueJoin {
		t.Errorf("Absolute paths using given prefixes have to be joined properly! Expected: %s, Got %s",
			assertValueJoin, result)
	}

	result = lib.ConvertToAbsolutePath(p, "")
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

	for i := range validUrls {
		valid := lib.IsUrlValid(validUrls[i])
		if !valid {
			t.Errorf("Expected %s to be valid url.", validUrls[i])
		}
	}

	for i := range invalidUrls {
		invalid := lib.IsUrlValid(invalidUrls[i])
		if invalid {
			t.Errorf("Expected %s to be invalid url.", invalidUrls[i])
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
