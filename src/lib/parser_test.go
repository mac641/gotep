package lib

import (
	"log"
	"strings"
	"testing"
)

type PrepareHttpRequestsHelperMock struct{}

// TODO: check if it is better practice to only mock viper calls, rather than the whole getConfig function
func (hm PrepareHttpRequestsHelperMock) getConfig(envMatches []string, conEnv string) map[string]string {
	mockedMatchedConfig := make(map[string]string)

	for i := range envMatches {
		match := strings.Trim(envMatches[i], "{}")

		switch match {
		case "origin":
			mockedMatchedConfig["origin"] = "http://localhost:3333"
		case "location":
			mockedMatchedConfig["location"] = "tests"
		case "headervalue":
			mockedMatchedConfig["headervalue"] = "1234"
		case "jsonvalue":
			mockedMatchedConfig["jsonvalue"] = "\"test\": \"some text with\nline break\""
		default:
			log.Fatal("Something's wrong with your function mock! Compare mock and test files.")
		}
	}

	return mockedMatchedConfig
}

const testDataForPreparing = `
###
### Test ipv4
192.168.178.2/index/https
Accept: application/json

### Test multiple headers
GET 192.168.178.3/index
Accept: application/json
  text/html
// Comment 1
From: test@test.com
### Test with message body and absolute path
POST http://www.example.com/test1
Content-Type: application/json

< ./input.json

### Test with message body and absolute path
POST http://www.example.com/test1 HTTP/2
Content-Type: application/json

{
  "html": "Html message"
}

###
### Test ipv6
GET [1111:2222:abcd:4444:5a6c:7777:9c9f:6789]/index
# Comment 924875q3948579
Accept: application/json

### Test variables in request line
GET {{origin}}/{{location}}/available
Accept: application/json

### Test variables in header / message
POST {{origin}}/{{location}}/pdf
Content-Type: {{headervalue}}

{
  "html": "<h1>I'm an example heading!</h1>",
  {{jsonvalue}}
}

### Test response handler / ref
POST http://example.com/auth
Content-Type: application/json

< input.json
> {% client.global.set("auth", response.body.token);
console.log('Hello world!'); %}
<> previous-response.200.json

### Test asterisk form
OPTIONS * HTTP/1.1
Host: http://example.com:8080
### Test no header and message
OPTIONS * HTTP/1.1
### Test message only
OPTIONS * HTTP/1.1

{
  "html": "<h1>I'm an example heading!</h1>",
  {{jsonvalue}}
}
`

func TestPrepareHttpRequests(t *testing.T) {
	requests := PrepareHttpRequests(testDataForPreparing, "default", PrepareHttpRequestsHelperMock{})

	for i := range requests {
		result := requests[i]
		if regSeparator.MatchString(result) {
			t.Errorf("Expected %s not to contain request separators!", result)
		}

		if result == "" {
			t.Errorf("Expected %s not to be empty!", result)
		}

		if regComments.MatchString(result) {
			t.Errorf("Expected %s not to contain comments!", result)
		}

		// TODO: test if correct env var value has been inserted
		if regEnv.MatchString(result) {
			t.Errorf("Expected %s not to contain env variables!", result)
		}

		if regResponseHandler.MatchString(result) {
			t.Errorf("Expected %s not to contain response handlers!", result)
		}

		if regResponseRef.MatchString(result) {
			t.Errorf("Expected %s not to contain response refs!", result)
		}
	}
}
