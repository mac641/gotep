package lib_test

import (
	"bufio"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/mac641/gotep/src/lib"
)

var (
	validator = lib.Validator{}
)

func TestSend(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
	}
	absPathToBody, err := filepath.Abs(path.Join(cwd, "../../test_files/input.json"))
	if err != nil {
		t.Errorf(err.Error())
	}
	file, err := os.Open(absPathToBody)
	if err != nil {
		t.Errorf(err.Error())
	}
	request, err := http.NewRequest(http.MethodPost, "http://example.com", bufio.NewReader(file))
	if err != nil {
		t.Errorf(err.Error())
	}
	requests := []http.Request{}
	requests = append(requests, *request)
	responses := validator.Send(requests)

	if len(responses) > 1 {
		t.Errorf("Expected responses to contain 1 response but got %d", len(responses))
	}
	if reflect.TypeOf(responses[0]).String() != "*http.Response" {
		t.Errorf("Expected response to be of type http.Response but got %s", reflect.TypeOf(responses[0]).String())
	}
}
