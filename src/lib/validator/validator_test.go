package validator_test

import (
	"bufio"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/mac641/gotep/src/lib/validator"
)

var (
	validator_test = validator.Validator{}
)

func TestAnalyze(t *testing.T) {
	resSuccessful := http.Response{
		StatusCode: 200,
	}
	resFailed := http.Response{
		StatusCode: 400,
	}

	succeeded := validator_test.Analyze([]*http.Response{&resSuccessful, &resFailed})
	if succeeded != 1 {
		t.Errorf("Expected analyzing to result in %d item but got %d", 1, succeeded)
	}

}

func TestSend(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
	}
	absPathToBody, err := filepath.Abs(path.Join(cwd, "../../../test_files/input.json"))
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
	responses, err := validator_test.Send(requests)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(responses) > 1 {
		t.Errorf("Expected responses to contain 1 response but got %d", len(responses))
	}
	if reflect.TypeOf(responses[0]).String() != "*http.Response" {
		t.Errorf("Expected response to be of type http.Response but got %s", reflect.TypeOf(responses[0]).String())
	}
}
