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
	testDataForParsing = []string{
		`192.168.178.2/index/https
Accept: application/json

`,
		`GET 192.168.178.3/index
Accept: application/json
 text/html
From: test@test.com
`,
		`POST http://www.example.com/test1
Content-Type: application/json

< ./input.json

`,
		`POST http://www.example.com/test1 HTTP/2
Content-Type: application/json

{
  "html": "Html message"
}

`,
		`GET [1111:2222:abcd:4444:5a6c:7777:9c9f:6789]/index
Accept: application/json

`,
		`GET http://localhost:3333/tests/available
Accept: application/json

`,
		`POST http://localhost:3333/tests/pdf
Content-Type: 1234

{
  "html": "<h1>I'm an example heading!</h1>",
  "test": "some text with
line break"
}

`,
		`POST http://example.com/auth
Content-Type: application/json

< input.json

`,
		`OPTIONS * HTTP/1.1
Host: http://example.com:8080
`,
		`OPTIONS * HTTP/1.1
`,
		`OPTIONS * HTTP/1.1

{
  "html": "<h1>I'm an example heading!</h1>",
  "test": "some text with
line break"
}
`,
	}
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
	regSeparator = regexp.MustCompile(`(?m)^(###[^\r\n]*(\r?\n|\r))`)
	regComments  = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*(\r?\n|\r))`)
	regEnv       = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
	// TODO: enhance response handler regex to be able to determine, if '###' and '%}' literals have been used inside the handler block
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
	httpRequests := parser.Parse(testDataForParsing)

	for i := range httpRequests {
		request := httpRequests[i]

		// TODO: use better url validation for testing
		if request.URL.String() == "" {
			t.Errorf("Expected requests to contain valid url!")
		}
	}
}
