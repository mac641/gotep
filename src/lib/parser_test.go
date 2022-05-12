package lib_test

import (
	"io"
	"regexp"
	"testing"

	"github.com/mac641/gotep/src/lib"
)

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
)

func TestPrepare(t *testing.T) {
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
		regSeparator       = regexp.MustCompile(`(?m)^(###[^\r\n]*(\r?\n|\r))`)
		regComments        = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*(\r?\n|\r))`)
		regEnv             = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
		regResponseHandler = regexp.MustCompile(`(?m)^>[ \t\f]+([^\r\n]*(\r?\n|\r)$|{%(.|(\r?\n|\r))+%})(\r?\n|\r)`)
		regResponseRef     = regexp.MustCompile(`(?m)^<>[ \t\f]+[^\r\n]*$(\r?\n|\r)`)
	)

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

		requestLineHeaderBody = []string{`POST http://www.example.com/test1 HTTP/2
Content-Type: application/json

{
	"html": "Html message"
}

`}
		requestLineHeaderBodyLineBreak = []string{`POST http://localhost:3333/tests/pdf
Content-Type: 1234

{
	"html": "<h1>I'm an example heading!</h1>",
	"test": "some text with
	line break"
}

`}
		requestLineHeaderBodyInputFileRef = []string{`POST http://www.example.com/test1
Content-Type: application/json

< ../test_files/input.json

`}
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
		assertValues = []string{}
	)

	// Request line only
	requestLineOnlyParsed := parser.Parse(requestLineOnly)
	assertValues = []string{"http://192.168.178.2/index/https"}
	if len(requestLineOnlyParsed) != 1 {
		t.Errorf("Expected requestLineOnly to have one request but got %d!", len(requestLineOnlyParsed))
	}
	if requestLineOnlyParsed[0].URL.String() != assertValues[0] {
		t.Errorf("Expected requestLineOnly to be %s but got %s!", assertValues[0],
			requestLineOnlyParsed[0].URL.String())
	}

	// Request line with method prepended
	requestLineOnlyMethodParsed := parser.Parse(requestLineOnlyMethod)
	assertValues = []string{"PUT", "http://192.168.178.2/index/https"}
	if len(requestLineOnlyMethodParsed) != 1 {
		t.Errorf("Expected requestLineOnlyMethod to have one request but got %d!", len(requestLineOnlyMethodParsed))
	}
	if requestLineOnlyMethodParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineOnlyMethod to use method %s but got %s!", assertValues[0],
			requestLineOnlyMethodParsed[0].Method)
	}
	if requestLineOnlyMethodParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineOnlyMethod to use url %s but got %s!", assertValues[1],
			requestLineOnlyMethodParsed[0].URL.String())
	}

	// Request line with method and HTTP version
	requestLineOnlyMethodHttpParsed := parser.Parse(requestLineOnlyMethodHttp)
	assertValues = []string{"GET", "http://192.168.178.2/index/https"}
	if len(requestLineOnlyMethodHttpParsed) != 1 {
		t.Errorf("Expected requestLineOnlyMethodHttp to have one request but got %d!", len(requestLineOnlyMethodHttpParsed))
	}
	if requestLineOnlyMethodHttpParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineOnlyMethodHttp to use method %s but got %s!",
			assertValues[0], requestLineOnlyMethodHttpParsed[0].Method)
	}
	if requestLineOnlyMethodHttpParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineOnlyMethodHttp to use url %s but got %s!",
			assertValues[1], requestLineOnlyMethodHttpParsed[0].URL.String())
	}
	if requestLineOnlyMethodHttpParsed[0].Proto != "HTTP/1.1" && requestLineOnlyMethodHttpParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineOnlyMethodHttp to use proto HTTP/1.1 or HTTP/2 but got %s!",
			requestLineOnlyMethodHttpParsed[0].Proto)
	}

	// Request line and one header
	requestLineHeaderParsed := parser.Parse(requestLineHeader)
	assertValues = []string{"GET", "http://192.168.178.2/index/https", "Accept", "application/json"}
	if len(requestLineHeaderParsed) != 1 {
		t.Errorf("Expected requestLineHeader to have one request but got %d!", len(requestLineHeaderParsed))
	}
	if requestLineHeaderParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeader to use method %s but got %s!",
			assertValues[0], requestLineHeaderParsed[0].Method)
	}
	if requestLineHeaderParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeader to use url %s but got %s!",
			assertValues[1], requestLineHeaderParsed[0].URL.String())
	}
	if len(requestLineHeaderParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeader to have only one header but got %d!",
			len(requestLineHeaderParsed[0].Header))
	}
	for key, value := range requestLineHeaderParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeader to have header key %s but got %s!",
				assertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeader to have one header value but got %d!", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeader to have header value %s but got %s!",
				assertValues[3], value[0])
		}
	}

	// Request line with one header and multiple header values
	requestLineHeaderMultipleValuesParsed := parser.Parse(requestLineHeaderMultipleValues)
	assertValues = []string{"GET", "http://192.168.178.2/index/https", "Accept", "application/json", "text/html"}
	if len(requestLineHeaderMultipleValuesParsed) != 1 {
		t.Errorf("Expected requestLineHeaderMultipleValues to have one request but got %d!",
			len(requestLineHeaderMultipleValuesParsed))
	}
	if requestLineHeaderMultipleValuesParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderMultipleValues to use method %s but got %s!",
			assertValues[0], requestLineHeaderMultipleValuesParsed[0].Method)
	}
	if requestLineHeaderMultipleValuesParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderMultipleValues to use url %s but got %s!",
			assertValues[1], requestLineHeaderMultipleValuesParsed[0].URL.String())
	}
	if len(requestLineHeaderMultipleValuesParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderMultipleValues to have only one header but got %d!",
			len(requestLineHeaderMultipleValuesParsed[0].Header))
	}
	for key, value := range requestLineHeaderMultipleValuesParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderMultipleValues to have header key %s but got %s!",
				assertValues[2], key)
		}
		if len(value) != 2 {
			t.Errorf("Expected requestLineHeaderMultipleValues to have two header values but got %d!", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderMultipleValues' first header value to be %s but got %s!",
				assertValues[3], value[0])
		}
		if value[1] != assertValues[4] {
			t.Errorf("Expected requestLineHeaderMultipleValues' second header value to be %s but got %s!",
				assertValues[4], value[1])
		}
	}

	// Request line with multiple headers with one value each
	requestLineMultipleHeadersParsed := parser.Parse(requestLineMultipleHeaders)
	assertValues = []string{"GET", "http://192.168.178.2/index/https", "Accept",
		"application/json", "From", "test@test.com"}
	if len(requestLineMultipleHeadersParsed) != 1 {
		t.Errorf("Expected requestLineMultipleHeaders to have one request but got %d!",
			len(requestLineMultipleHeadersParsed))
	}
	if requestLineMultipleHeadersParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineMultipleHeaders to use method %s but got %s!",
			assertValues[0], requestLineMultipleHeadersParsed[0].Method)
	}
	if requestLineMultipleHeadersParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineMultipleHeaders to use url %s but got %s!",
			assertValues[1], requestLineMultipleHeadersParsed[0].URL.String())
	}
	if len(requestLineMultipleHeadersParsed[0].Header) != 2 {
		t.Errorf("Expected requestLineMultipleHeaders to have two headers but got %d!",
			len(requestLineMultipleHeadersParsed[0].Header))
	}
	for key, value := range requestLineMultipleHeadersParsed[0].Header {
		if key != assertValues[2] && key != assertValues[4] {
			t.Errorf("Expected requestLineMultipleHeaders to have header key %s or %s but got %s!",
				assertValues[2], assertValues[4], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineMultipleHeaders to have one header value each but got %d!", len(value))
		}
		if value[0] != assertValues[3] &&
			value[0] != assertValues[5] {
			t.Errorf("Expected requestLineMultipleHeaders' header values to be %s or %s but got %s!",
				assertValues[3], assertValues[5], value[0])
		}
	}

	// Request line with one header and body
	requestLineHeaderBodyParsed := parser.Parse(requestLineHeaderBody)
	assertValues = []string{"POST", "http://www.example.com/test1", "Content-Type",
		"application/json",
		`{
	"html": "Html message"
}`}
	if len(requestLineHeaderBodyParsed) != 1 {
		t.Errorf("Expected requestLineHeaderBody to have one request but got %d!",
			len(requestLineHeaderBodyParsed))
	}
	if requestLineHeaderBodyParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderBody to use method %s but got %s!",
			assertValues[0], requestLineHeaderBodyParsed[0].Method)
	}
	if requestLineHeaderBodyParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderBody to use url %s but got %s!",
			assertValues[1], requestLineHeaderBodyParsed[0].URL.String())
	}
	if requestLineHeaderBodyParsed[0].Proto != "HTTP/1.1" && requestLineHeaderBodyParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineHeaderBody to use proto HTTP/1.1 or HTTP/2 but got %s!",
			requestLineHeaderBodyParsed[0].Proto)
	}
	if len(requestLineHeaderBodyParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderBody to have one header but got %d!",
			len(requestLineHeaderBodyParsed[0].Header))
	}
	for key, value := range requestLineHeaderBodyParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderBody to have header key %s or %s but got %s!",
				assertValues[2], assertValues[4], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeaderBody to have one header value but got %d!", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderBody header value to be %s but got %s!",
				assertValues[3], value[0])
		}
	}
	requestLineHeaderBodyParsedRead, err := io.ReadAll(requestLineHeaderBodyParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineHeaderBody's body due to \n%s", err.Error())
	}
	if string(requestLineHeaderBodyParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineHeaderBody's body value to be \"%s\" but got \"%s\"!", assertValues[4],
			string(requestLineHeaderBodyParsedRead))
	}

	// Request line with one header and body containing line break
	requestLineHeaderBodyLineBreakParsed := parser.Parse(requestLineHeaderBodyLineBreak)
	assertValues = []string{"POST", "http://localhost:3333/tests/pdf", "Content-Type", "1234",
		`{
	"html": "<h1>I'm an example heading!</h1>",
	"test": "some text with
	line break"
}`}
	if len(requestLineHeaderBodyLineBreakParsed) != 1 {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to have one request but got %d!",
			len(requestLineHeaderBodyLineBreakParsed))
	}
	if requestLineHeaderBodyLineBreakParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to use method %s but got %s!",
			assertValues[0], requestLineHeaderBodyLineBreakParsed[0].Method)
	}
	if requestLineHeaderBodyLineBreakParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to use url %s but got %s!",
			assertValues[1], requestLineHeaderBodyLineBreakParsed[0].URL.String())
	}
	if len(requestLineHeaderBodyLineBreakParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to have one header but got %d!",
			len(requestLineHeaderBodyLineBreakParsed[0].Header))
	}
	for key, value := range requestLineHeaderBodyLineBreakParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderBodyLineBreak to have header key %s or %s but got %s!",
				assertValues[2], assertValues[4], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeaderBodyLineBreak to have one header value but got %d!", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderBodyLineBreak header value to be %s but got %s!",
				assertValues[3], value[0])
		}
	}
	requestLineHeaderBodyLineBreakParsedRead, err := io.ReadAll(requestLineHeaderBodyLineBreakParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineHeaderBodyLineBreak's body due to \n%s", err.Error())
	}
	if string(requestLineHeaderBodyLineBreakParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineHeaderBodyLineBreak's body value to be \"%s\" but got \"%s\"!", assertValues[4],
			string(requestLineHeaderBodyLineBreakParsedRead))
	}

	// Request line with one header and body as input file ref
	requestLineHeaderBodyInputFileRefParsed := parser.Parse(requestLineHeaderBodyInputFileRef)
	assertValues = []string{"POST", "http://www.example.com/test1", "Content-Type", "application/json",
		`{
  "html": "<h1>I'm an example heading!</h1>",
  "testnumber": 234145
}
`}
	if len(requestLineHeaderBodyInputFileRefParsed) != 1 {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to have one request but got %d!",
			len(requestLineHeaderBodyInputFileRefParsed))
	}
	if requestLineHeaderBodyInputFileRefParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to use method %s but got %s!",
			assertValues[0], requestLineHeaderBodyInputFileRefParsed[0].Method)
	}
	if requestLineHeaderBodyInputFileRefParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to use url %s but got %s!",
			assertValues[1], requestLineHeaderBodyInputFileRefParsed[0].URL.String())
	}
	if len(requestLineHeaderBodyInputFileRefParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to have one header but got %d!",
			len(requestLineHeaderBodyInputFileRefParsed[0].Header))
	}
	for key, value := range requestLineHeaderBodyInputFileRefParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderBodyInputFileRef to have header key %s or %s but got %s!",
				assertValues[2], assertValues[4], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeaderBodyInputFileRef to have one header value but got %d!", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderBodyInputFileRef header value to be %s but got %s!",
				assertValues[3], value[0])
		}
	}
	requestLineHeaderBodyInputFileRefParsedRead, err := io.ReadAll(requestLineHeaderBodyInputFileRefParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineHeaderBodyInputFileRef's body due to \n%s", err.Error())
	}
	if string(requestLineHeaderBodyInputFileRefParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef's body value to be \"%s\" but got \"%s\"!",
			assertValues[4], string(requestLineHeaderBodyInputFileRefParsedRead))
	}

	// Request line with one header and body as multipart/form-data
	// TODO: expand test when adding multipart/form-data parsing support
	requestLineHeaderMultipartFormDataParsed := parser.Parse(requestLineHeaderMultipartFormData)
	if len(requestLineHeaderMultipartFormDataParsed) != 0 {
		t.Errorf("Expected requestLineHeaderMultipartFormData to be skipped but got request!")
	}
}
