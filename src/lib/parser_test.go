package lib_test

import (
	"regexp"
	"testing"

	"github.com/mac641/gotep/src/lib"
)

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

var (
	parser = lib.Parser{
		ConEnv:     "default",
		Verbose:    false,
		PathPrefix: "../../test_files",
		Config: map[string]string{
			"origin":      "http://localhost:3333",
			"location":    "tests",
			"headervalue": "1234",
			"jsonvalue":   "\"test\": \"some text with\nline break\"",
		},
	}
	regSeparator       = regexp.MustCompile(`(?m)^(###[^\r\n]*(\r?\n|\r))`)
	regComments        = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*(\r?\n|\r))`)
	regEnv             = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
	regResponseHandler = regexp.MustCompile(`(?m)^>[ \t\f]+([^\r\n]*(\r?\n|\r)$|{%(.|(\r?\n|\r))+%})(\r?\n|\r)`)
	regResponseRef     = regexp.MustCompile(`(?m)^<>[ \t\f]+[^\r\n]*$(\r?\n|\r)`)
)

func TestPrepare(t *testing.T) {
	requests := parser.Prepare(testDataForPreparing)

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

func TestParse(t *testing.T) {
	var (
		requestLineOnly           = []string{"192.168.178.2/index/https\n"}
		requestLineOnlyMethod     = []string{"PUT 192.168.178.2/index/https\n"}
		requestLineOnlyMethodHttp = []string{"GET 192.168.178.2/index/https HTTP/1\n"}

		requestLineHeader               = []string{"GET 192.168.178.2/index/https\nAccept: application/json"}
		requestLineHeaderMultipleValues = []string{
			"GET 192.168.178.2/index/https\nAccept: application/json\n  text/html"}
		requestLineMultipleHeaders = []string{
			"GET 192.168.178.2/index/https\nAccept: application/json\nFrom: test@test.com"}

		// 		requestLineHeaderBody = []string{`POST http://www.example.com/test1 HTTP/2
		// Content-Type: application/json

		// {
		//   "html": "Html message"
		// }

		// `}
		// 		requestLineHeaderBodyLineBreak = []string{`POST http://localhost:3333/tests/pdf
		// Content-Type: 1234

		// {
		// "html": "<h1>I'm an example heading!</h1>",
		// "test": "some text with
		// line break"
		// }

		// `}
		// 		requestLineHeaderBodyInputFileRef = []string{`POST http://www.example.com/test1
		// Content-Type: application/json

		// < ./input.json

		// `}
		requestLineHeaderMultipartFormData = []string{`POST http://example.com/api/upload
Content-Type: multipart/form-data; boundary=abcd

--abcd
Content-Disposition: form-data; name="text"

Text
--abcd
Content-Disposition: form-data; name="file_to_send"; filename="input.txt"

< ./input.txt
--abcd--

`}
	)

	// Request line only
	requestLineOnlyParsed := parser.Parse(requestLineOnly)
	requestLineOnlyAssertValue := "http://192.168.178.2/index/https"
	if len(requestLineOnlyParsed) != 1 {
		t.Errorf("Expected requestLineOnly to have one request but got %d!", len(requestLineOnlyParsed))
	}
	if requestLineOnlyParsed[0].URL.String() != requestLineOnlyAssertValue {
		t.Errorf("Expected requestLineOnly to be %s but got %s!", requestLineOnlyAssertValue,
			requestLineOnlyParsed[0].URL.String())
	}

	requestLineOnlyMethodParsed := parser.Parse(requestLineOnlyMethod)
	requestLineOnlyMethodAssertValues := []string{"PUT", "http://192.168.178.2/index/https"}
	if len(requestLineOnlyMethodParsed) != 1 {
		t.Errorf("Expected requestLineOnlyMethod to have one request but got %d!", len(requestLineOnlyMethodParsed))
	}
	if requestLineOnlyMethodParsed[0].Method != requestLineOnlyMethodAssertValues[0] {
		t.Errorf("Expected requestLineOnlyMethod to use method %s but got %s!", requestLineOnlyMethodAssertValues[0],
			requestLineOnlyMethodParsed[0].Method)
	}
	if requestLineOnlyMethodParsed[0].URL.String() != requestLineOnlyMethodAssertValues[1] {
		t.Errorf("Expected requestLineOnlyMethod to use url %s but got %s!", requestLineOnlyMethodAssertValues[1],
			requestLineOnlyMethodParsed[0].URL.String())
	}

	requestLineOnlyMethodHttpParsed := parser.Parse(requestLineOnlyMethodHttp)
	requestLineOnlyMethodHttpAssertValues := []string{"GET", "http://192.168.178.2/index/https", "HTTP/1"}
	if len(requestLineOnlyMethodHttpParsed) != 1 {
		t.Errorf("Expected requestLineOnlyMethodHttp to have one request but got %d!", len(requestLineOnlyMethodHttpParsed))
	}
	if requestLineOnlyMethodHttpParsed[0].Method != requestLineOnlyMethodHttpAssertValues[0] {
		t.Errorf("Expected requestLineOnlyMethodHttp to use method %s but got %s!",
			requestLineOnlyMethodHttpAssertValues[0], requestLineOnlyMethodHttpParsed[0].Method)
	}
	if requestLineOnlyMethodHttpParsed[0].URL.String() != requestLineOnlyMethodHttpAssertValues[1] {
		t.Errorf("Expected requestLineOnlyMethodHttp to use url %s but got %s!",
			requestLineOnlyMethodHttpAssertValues[1], requestLineOnlyMethodHttpParsed[0].URL.String())
	}
	if requestLineOnlyMethodHttpParsed[0].Proto != "HTTP/1.1" && requestLineOnlyMethodHttpParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineOnlyMethodHttp to use proto HTTP/1.1 or HTTP/2 but got %s!",
			requestLineOnlyMethodHttpParsed[0].Proto)
	}

	// Request line and headers combined
	requestLineHeaderParsed := parser.Parse(requestLineHeader)
	requestLineHeaderAssertValues := []string{"GET", "http://192.168.178.2/index/https", "Accept", "application/json"}
	if len(requestLineHeaderParsed) != 1 {
		t.Errorf("Expected requestLineHeader to have one request but got %d!", len(requestLineHeaderParsed))
	}
	if requestLineHeaderParsed[0].Method != requestLineHeaderAssertValues[0] {
		t.Errorf("Expected requestLineHeader to use method %s but got %s!",
			requestLineHeaderAssertValues[0], requestLineHeaderParsed[0].Method)
	}
	if requestLineHeaderParsed[0].URL.String() != requestLineHeaderAssertValues[1] {
		t.Errorf("Expected requestLineHeader to use url %s but got %s!",
			requestLineHeaderAssertValues[1], requestLineHeaderParsed[0].URL.String())
	}
	if len(requestLineHeaderParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeader to have only one header but got %d!",
			len(requestLineHeaderParsed[0].Header))
	}
	for key, value := range requestLineHeaderParsed[0].Header {
		if key != requestLineHeaderAssertValues[2] {
			t.Errorf("Expected requestLineHeader to have header key %s but got %s!",
				requestLineHeaderAssertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeader to have one header value but got %d!", len(value))
		}
		if value[0] != requestLineHeaderAssertValues[3] {
			t.Errorf("Expected requestLineHeader to have header value %s but got %s!",
				requestLineHeaderAssertValues[3], value[0])
		}
	}

	requestLineHeaderMultipleValuesParsed := parser.Parse(requestLineHeaderMultipleValues)
	requestLineHeaderMultipleValuesAssertValues := []string{"GET", "http://192.168.178.2/index/https", "Accept",
		"application/json", "text/html"}
	if len(requestLineHeaderMultipleValuesParsed) != 1 {
		t.Errorf("Expected requestLineHeaderMultipleValues to have one request but got %d!",
			len(requestLineHeaderMultipleValuesParsed))
	}
	if requestLineHeaderMultipleValuesParsed[0].Method != requestLineHeaderMultipleValuesAssertValues[0] {
		t.Errorf("Expected requestLineHeaderMultipleValues to use method %s but got %s!",
			requestLineHeaderMultipleValuesAssertValues[0], requestLineHeaderMultipleValuesParsed[0].Method)
	}
	if requestLineHeaderMultipleValuesParsed[0].URL.String() != requestLineHeaderMultipleValuesAssertValues[1] {
		t.Errorf("Expected requestLineHeaderMultipleValues to use url %s but got %s!",
			requestLineHeaderMultipleValuesAssertValues[1], requestLineHeaderMultipleValuesParsed[0].URL.String())
	}
	if len(requestLineHeaderMultipleValuesParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderMultipleValues to have only one header but got %d!",
			len(requestLineHeaderMultipleValuesParsed[0].Header))
	}
	for key, value := range requestLineHeaderMultipleValuesParsed[0].Header {
		if key != requestLineHeaderMultipleValuesAssertValues[2] {
			t.Errorf("Expected requestLineHeaderMultipleValues to have header key %s but got %s!",
				requestLineHeaderMultipleValuesAssertValues[2], key)
		}
		if len(value) != 2 {
			t.Errorf("Expected requestLineHeaderMultipleValues to have two header values but got %d!", len(value))
		}
		if value[0] != requestLineHeaderMultipleValuesAssertValues[3] {
			t.Errorf("Expected requestLineHeaderMultipleValues' first header value to be %s but got %s!",
				requestLineHeaderAssertValues[3], value[0])
		}
		if value[1] != requestLineHeaderMultipleValuesAssertValues[4] {
			t.Errorf("Expected requestLineHeaderMultipleValues' second header value to be %s but got %s!",
				requestLineHeaderAssertValues[4], value[1])
		}
	}

	requestLineMultipleHeadersParsed := parser.Parse(requestLineMultipleHeaders)
	requestLineMultipleHeadersAssertValues := []string{"GET", "http://192.168.178.2/index/https", "Accept",
		"application/json", "From", "test@test.com"}
	if len(requestLineMultipleHeadersParsed) != 1 {
		t.Errorf("Expected requestLineMultipleHeaders to have one request but got %d!",
			len(requestLineMultipleHeadersParsed))
	}
	if requestLineMultipleHeadersParsed[0].Method != requestLineMultipleHeadersAssertValues[0] {
		t.Errorf("Expected requestLineMultipleHeaders to use method %s but got %s!",
			requestLineMultipleHeadersAssertValues[0], requestLineMultipleHeadersParsed[0].Method)
	}
	if requestLineMultipleHeadersParsed[0].URL.String() != requestLineMultipleHeadersAssertValues[1] {
		t.Errorf("Expected requestLineMultipleHeaders to use url %s but got %s!",
			requestLineMultipleHeadersAssertValues[1], requestLineMultipleHeadersParsed[0].URL.String())
	}
	if len(requestLineMultipleHeadersParsed[0].Header) != 2 {
		t.Errorf("Expected requestLineMultipleHeaders to have two headers but got %d!",
			len(requestLineMultipleHeadersParsed[0].Header))
	}
	for key, value := range requestLineMultipleHeadersParsed[0].Header {
		if key != requestLineMultipleHeadersAssertValues[2] && key != requestLineMultipleHeadersAssertValues[4] {
			t.Errorf("Expected requestLineMultipleHeaders to have header key %s or %s but got %s!",
				requestLineMultipleHeadersAssertValues[2], requestLineMultipleHeadersAssertValues[4], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineMultipleHeaders to have one header value each but got %d!", len(value))
		}
		if value[0] != requestLineMultipleHeadersAssertValues[3] &&
			value[0] != requestLineMultipleHeadersAssertValues[5] {
			t.Errorf("Expected requestLineMultipleHeaders' header values to be %s or %s but got %s!",
				requestLineHeaderAssertValues[3], requestLineHeaderAssertValues[5], value[0])
		}
	}

	// TODO: Insert remaining test cases here

	requestLineHeaderMultipartFormDataParsed := parser.Parse(requestLineHeaderMultipartFormData)
	if len(requestLineHeaderMultipartFormDataParsed) != 0 {
		t.Errorf("Expected requestLineHeaderMultipartFormData to be skipped but got request!")
	}
}
