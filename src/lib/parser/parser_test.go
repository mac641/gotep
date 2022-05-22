package parser_test

import (
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/mac641/gotep/src/lib/context"
	"github.com/mac641/gotep/src/lib/parser"
)

var (
	ctx        = context.GetContext()
	parserTest = parser.Parser{}
)

func init() {
	ctx.SetConfigFilePath("../../../test_files/http-client.env.json")
	ctx.SetPathPrefix("../../../test_files")
}

func TestParse(t *testing.T) {
	var (
		requestLineOnly           = []string{"192.168.178.2/index/https\n"}
		requestLineOnlyMethod     = []string{"PUT 192.168.178.2/index/https\n"}
		requestLineOnlyMethodHttp = []string{"GET 192.168.178.2/index/https HTTP/1\n"}
		requestLineOnlySplit      = []string{`http://example.com/
  api
  /get
`}
		requestLineOnlySplitEncoded = []string{`GET http://example.com/
  %20api%20
  +/get+
`}

		requestLineHeader      = []string{"GET 192.168.178.2/index/https\nAccept: application/json"}
		requestLineSplitHeader = []string{`http://example.com/
  api
  /get
From: eugen@blabla.com`}
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
		requestLineSplitHeaderBody = []string{`POST http://example.com/
  api
  /get
From: eugen@blabla.com

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
		requestLineAsteriskFormHeader = []string{`OPTIONS * HTTP/1.1
Host: http://example.com:8080`}
		requestLineOriginFormHeader = []string{`GET /api/get
Host: example.com`}
		assertValues = []string{}
	)

	// Request line only
	requestLineOnlyParsed, err := parserTest.Parse(requestLineOnly)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"http://192.168.178.2/index/https"}
	if len(requestLineOnlyParsed) != 1 {
		t.Errorf("Expected requestLineOnly to have one request but got %d", len(requestLineOnlyParsed))
	}
	if requestLineOnlyParsed[0].URL.String() != assertValues[0] {
		t.Errorf("Expected requestLineOnly to be %s but got %s", assertValues[0],
			requestLineOnlyParsed[0].URL.String())
	}

	// Request line with method prepended
	requestLineOnlyMethodParsed, err := parserTest.Parse(requestLineOnlyMethod)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"PUT", "http://192.168.178.2/index/https"}
	if len(requestLineOnlyMethodParsed) != 1 {
		t.Errorf("Expected requestLineOnlyMethod to have one request but got %d", len(requestLineOnlyMethodParsed))
	}
	if requestLineOnlyMethodParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineOnlyMethod to use method %s but got %s", assertValues[0],
			requestLineOnlyMethodParsed[0].Method)
	}
	if requestLineOnlyMethodParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineOnlyMethod to use url %s but got %s", assertValues[1],
			requestLineOnlyMethodParsed[0].URL.String())
	}

	// Request line with method and HTTP version
	requestLineOnlyMethodHttpParsed, err := parserTest.Parse(requestLineOnlyMethodHttp)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"GET", "http://192.168.178.2/index/https"}
	if len(requestLineOnlyMethodHttpParsed) != 1 {
		t.Errorf("Expected requestLineOnlyMethodHttp to have one request but got %d", len(requestLineOnlyMethodHttpParsed))
	}
	if requestLineOnlyMethodHttpParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineOnlyMethodHttp to use method %s but got %s",
			assertValues[0], requestLineOnlyMethodHttpParsed[0].Method)
	}
	if requestLineOnlyMethodHttpParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineOnlyMethodHttp to use url %s but got %s",
			assertValues[1], requestLineOnlyMethodHttpParsed[0].URL.String())
	}
	if requestLineOnlyMethodHttpParsed[0].Proto != "HTTP/1.1" && requestLineOnlyMethodHttpParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineOnlyMethodHttp to use proto HTTP/1.1 or HTTP/2 but got %s",
			requestLineOnlyMethodHttpParsed[0].Proto)
	}

	// Request line only in split form
	requestLineOnlySplitParsed, err := parserTest.Parse(requestLineOnlySplit)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"http://example.com/api/get"}
	if len(requestLineOnlySplitParsed) != 1 {
		t.Errorf("Expected requestLineOnlySplit to have one request but got %d", len(requestLineOnlySplitParsed))
	}
	if requestLineOnlySplitParsed[0].URL.String() != assertValues[0] {
		t.Errorf("Expected requestLineOnlySplit to be %s but got %s", assertValues[0],
			requestLineOnlySplitParsed[0].URL.String())
	}

	// Request line only in split form with encoded url
	requestLineOnlySplitEncodedParsed, err := parserTest.Parse(requestLineOnlySplitEncoded)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"GET", "/ api +/get+"}
	if len(requestLineOnlySplitEncodedParsed) != 1 {
		t.Errorf("Expected requestLineOnlySplitEncoded to have one request but got %d",
			len(requestLineOnlySplitEncodedParsed))
	}
	if requestLineOnlySplitEncodedParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineOnlySplitEncoded to use method %s but got %s",
			assertValues[0], requestLineOnlySplitEncodedParsed[0].Method)
	}
	if requestLineOnlySplitEncodedParsed[0].URL.Path != assertValues[1] {
		t.Errorf("Expected requestLineOnlySplitEncoded to be %s but got %s", assertValues[1],
			requestLineOnlySplitEncodedParsed[0].URL.String())
	}

	// Request line and one header
	requestLineHeaderParsed, err := parserTest.Parse(requestLineHeader)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"GET", "http://192.168.178.2/index/https", "Accept", "application/json"}
	if len(requestLineHeaderParsed) != 1 {
		t.Errorf("Expected requestLineHeader to have one request but got %d", len(requestLineHeaderParsed))
	}
	if requestLineHeaderParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeader to use method %s but got %s",
			assertValues[0], requestLineHeaderParsed[0].Method)
	}
	if requestLineHeaderParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeader to use url %s but got %s",
			assertValues[1], requestLineHeaderParsed[0].URL.String())
	}
	if len(requestLineHeaderParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeader to have only one header but got %d",
			len(requestLineHeaderParsed[0].Header))
	}
	for key, value := range requestLineHeaderParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeader to have header key %s but got %s",
				assertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeader to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeader to have header value %s but got %s",
				assertValues[3], value[0])
		}
	}

	// Split request line with one header
	requestLineSplitHeaderParsed, err := parserTest.Parse(requestLineSplitHeader)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"http://example.com/api/get", "From", "eugen@blabla.com"}
	if len(requestLineSplitHeaderParsed) != 1 {
		t.Errorf("Expected requestLineSplitHeader to have one request but got %d", len(requestLineSplitHeaderParsed))
	}
	if requestLineSplitHeaderParsed[0].URL.String() != assertValues[0] {
		t.Errorf("Expected requestLineSplitHeader to use url %s but got %s",
			assertValues[0], requestLineSplitHeaderParsed[0].URL.String())
	}
	if len(requestLineSplitHeaderParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineSplitHeader to have only one header but got %d",
			len(requestLineSplitHeaderParsed[0].Header))
	}
	for key, value := range requestLineSplitHeaderParsed[0].Header {
		if key != assertValues[1] {
			t.Errorf("Expected requestLineSplitHeader to have header key %s but got %s", assertValues[1], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineSplitHeader to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[2] {
			t.Errorf("Expected requestLineSplitHeader to have header value %s but got %s",
				assertValues[2], value[0])
		}
	}

	// Request line with one header and multiple header values
	requestLineHeaderMultipleValuesParsed, err := parserTest.Parse(requestLineHeaderMultipleValues)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"GET", "http://192.168.178.2/index/https", "Accept", "application/json", "text/html"}
	if len(requestLineHeaderMultipleValuesParsed) != 1 {
		t.Errorf("Expected requestLineHeaderMultipleValues to have one request but got %d",
			len(requestLineHeaderMultipleValuesParsed))
	}
	if requestLineHeaderMultipleValuesParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderMultipleValues to use method %s but got %s",
			assertValues[0], requestLineHeaderMultipleValuesParsed[0].Method)
	}
	if requestLineHeaderMultipleValuesParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderMultipleValues to use url %s but got %s",
			assertValues[1], requestLineHeaderMultipleValuesParsed[0].URL.String())
	}
	if len(requestLineHeaderMultipleValuesParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderMultipleValues to have only one header but got %d",
			len(requestLineHeaderMultipleValuesParsed[0].Header))
	}
	for key, value := range requestLineHeaderMultipleValuesParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderMultipleValues to have header key %s but got %s",
				assertValues[2], key)
		}
		if len(value) != 2 {
			t.Errorf("Expected requestLineHeaderMultipleValues to have two header values but got %d", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderMultipleValues' first header value to be %s but got %s",
				assertValues[3], value[0])
		}
		if value[1] != assertValues[4] {
			t.Errorf("Expected requestLineHeaderMultipleValues' second header value to be %s but got %s",
				assertValues[4], value[1])
		}
	}

	// Request line with multiple headers with one value each
	requestLineMultipleHeadersParsed, err := parserTest.Parse(requestLineMultipleHeaders)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"GET", "http://192.168.178.2/index/https", "Accept",
		"application/json", "From", "test@test.com"}
	if len(requestLineMultipleHeadersParsed) != 1 {
		t.Errorf("Expected requestLineMultipleHeaders to have one request but got %d",
			len(requestLineMultipleHeadersParsed))
	}
	if requestLineMultipleHeadersParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineMultipleHeaders to use method %s but got %s",
			assertValues[0], requestLineMultipleHeadersParsed[0].Method)
	}
	if requestLineMultipleHeadersParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineMultipleHeaders to use url %s but got %s",
			assertValues[1], requestLineMultipleHeadersParsed[0].URL.String())
	}
	if len(requestLineMultipleHeadersParsed[0].Header) != 2 {
		t.Errorf("Expected requestLineMultipleHeaders to have two headers but got %d",
			len(requestLineMultipleHeadersParsed[0].Header))
	}
	for key, value := range requestLineMultipleHeadersParsed[0].Header {
		if key != assertValues[2] && key != assertValues[4] {
			t.Errorf("Expected requestLineMultipleHeaders to have header key %s or %s but got %s",
				assertValues[2], assertValues[4], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineMultipleHeaders to have one header value each but got %d", len(value))
		}
		if value[0] != assertValues[3] &&
			value[0] != assertValues[5] {
			t.Errorf("Expected requestLineMultipleHeaders' header values to be %s or %s but got %s",
				assertValues[3], assertValues[5], value[0])
		}
	}

	// Request line with one header and body
	requestLineHeaderBodyParsed, err := parserTest.Parse(requestLineHeaderBody)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"POST", "http://www.example.com/test1", "Content-Type",
		"application/json",
		`{
	"html": "Html message"
}`}
	if len(requestLineHeaderBodyParsed) != 1 {
		t.Errorf("Expected requestLineHeaderBody to have one request but got %d",
			len(requestLineHeaderBodyParsed))
	}
	if requestLineHeaderBodyParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderBody to use method %s but got %s",
			assertValues[0], requestLineHeaderBodyParsed[0].Method)
	}
	if requestLineHeaderBodyParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderBody to use url %s but got %s",
			assertValues[1], requestLineHeaderBodyParsed[0].URL.String())
	}
	if requestLineHeaderBodyParsed[0].Proto != "HTTP/1.1" && requestLineHeaderBodyParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineHeaderBody to use proto HTTP/1.1 or HTTP/2 but got %s",
			requestLineHeaderBodyParsed[0].Proto)
	}
	if len(requestLineHeaderBodyParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderBody to have one header but got %d",
			len(requestLineHeaderBodyParsed[0].Header))
	}
	for key, value := range requestLineHeaderBodyParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderBody to have header key %s but got %s", assertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeaderBody to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderBody header value to be %s but got %s",
				assertValues[3], value[0])
		}
	}
	requestLineHeaderBodyParsedRead, err := io.ReadAll(requestLineHeaderBodyParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineHeaderBody's body due to \n%s", err.Error())
	}
	if string(requestLineHeaderBodyParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineHeaderBody's body value to be \n%s\n but got \n%s", assertValues[4],
			string(requestLineHeaderBodyParsedRead))
	}

	// Split request line with one header and message body
	requestLineSplitHeaderBodyParsed, err := parserTest.Parse(requestLineSplitHeaderBody)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"POST",
		"http://example.com/api/get",
		"From",
		"eugen@blabla.com",
		`{
	"html": "Html message"
}`}
	if len(requestLineSplitHeaderBodyParsed) != 1 {
		t.Errorf("Expected requestLineSplitHeaderBody to have one request but got %d",
			len(requestLineSplitHeaderBodyParsed))
	}
	if requestLineSplitHeaderBodyParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineSplitHeaderBody to use method %s but got %s",
			assertValues[0], requestLineSplitHeaderBodyParsed[0].Method)
	}
	if requestLineSplitHeaderBodyParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineSplitHeaderBody to use url %s but got %s",
			assertValues[1], requestLineSplitHeaderBodyParsed[0].URL.String())
	}
	if len(requestLineSplitHeaderBodyParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineSplitHeaderBody to have one header but got %d",
			len(requestLineSplitHeaderBodyParsed[0].Header))
	}
	for key, value := range requestLineSplitHeaderBodyParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineSplitHeaderBody to have header key %s but got %s", assertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineSplitHeaderBody to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineSplitHeaderBody header value to be %s but got %s",
				assertValues[3], value[0])
		}
	}
	requestLineSplitHeaderBodyParsedRead, err := io.ReadAll(requestLineSplitHeaderBodyParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineSplitHeaderBody's body due to \n%s", err.Error())
	}
	if string(requestLineSplitHeaderBodyParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineSplitHeaderBody's body value to be \n%s\n but got \n%s", assertValues[4],
			string(requestLineSplitHeaderBodyParsedRead))
	}

	// Request line with one header and body containing line break
	requestLineHeaderBodyLineBreakParsed, err := parserTest.Parse(requestLineHeaderBodyLineBreak)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"POST", "http://localhost:3333/tests/pdf", "Content-Type", "1234",
		`{
	"html": "<h1>I'm an example heading!</h1>",
	"test": "some text with
	line break"
}`}
	if len(requestLineHeaderBodyLineBreakParsed) != 1 {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to have one request but got %d",
			len(requestLineHeaderBodyLineBreakParsed))
	}
	if requestLineHeaderBodyLineBreakParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to use method %s but got %s",
			assertValues[0], requestLineHeaderBodyLineBreakParsed[0].Method)
	}
	if requestLineHeaderBodyLineBreakParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to use url %s but got %s",
			assertValues[1], requestLineHeaderBodyLineBreakParsed[0].URL.String())
	}
	if len(requestLineHeaderBodyLineBreakParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderBodyLineBreak to have one header but got %d",
			len(requestLineHeaderBodyLineBreakParsed[0].Header))
	}
	for key, value := range requestLineHeaderBodyLineBreakParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderBodyLineBreak to have header key %s but got %s",
				assertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeaderBodyLineBreak to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderBodyLineBreak header value to be %s but got %s",
				assertValues[3], value[0])
		}
	}
	requestLineHeaderBodyLineBreakParsedRead, err := io.ReadAll(requestLineHeaderBodyLineBreakParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineHeaderBodyLineBreak's body due to \n%s", err.Error())
	}
	if string(requestLineHeaderBodyLineBreakParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineHeaderBodyLineBreak's body value to be \n%s\n but got \n%s",
			assertValues[4], string(requestLineHeaderBodyLineBreakParsedRead))
	}

	// Request line with one header and body as input file ref
	requestLineHeaderBodyInputFileRefParsed, err := parserTest.Parse(requestLineHeaderBodyInputFileRef)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"POST", "http://www.example.com/test1", "Content-Type", "application/json",
		`{
  "html": "<h1>I'm an example heading!</h1>",
  "testnumber": 234145
}
`}
	if len(requestLineHeaderBodyInputFileRefParsed) != 1 {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to have one request but got %d",
			len(requestLineHeaderBodyInputFileRefParsed))
	}
	if requestLineHeaderBodyInputFileRefParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to use method %s but got %s",
			assertValues[0], requestLineHeaderBodyInputFileRefParsed[0].Method)
	}
	if requestLineHeaderBodyInputFileRefParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to use url %s but got %s",
			assertValues[1], requestLineHeaderBodyInputFileRefParsed[0].URL.String())
	}
	if len(requestLineHeaderBodyInputFileRefParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef to have one header but got %d",
			len(requestLineHeaderBodyInputFileRefParsed[0].Header))
	}
	for key, value := range requestLineHeaderBodyInputFileRefParsed[0].Header {
		if key != assertValues[2] {
			t.Errorf("Expected requestLineHeaderBodyInputFileRef to have header key %s but got %s",
				assertValues[2], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineHeaderBodyInputFileRef to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[3] {
			t.Errorf("Expected requestLineHeaderBodyInputFileRef header value to be %s but got %s",
				assertValues[3], value[0])
		}
	}
	requestLineHeaderBodyInputFileRefParsedRead, err := io.ReadAll(requestLineHeaderBodyInputFileRefParsed[0].Body)
	if err != nil {
		t.Errorf("Could not read requestLineHeaderBodyInputFileRef's body due to \n%s", err.Error())
	}
	if string(requestLineHeaderBodyInputFileRefParsedRead) != assertValues[4] {
		t.Errorf("Expected requestLineHeaderBodyInputFileRef's body value to be \n%s\n but got \n%s",
			assertValues[4], string(requestLineHeaderBodyInputFileRefParsedRead))
	}

	// Request line with one header and body as multipart/form-data
	// TODO: expand test when adding multipart/form-data parsing support
	requestLineHeaderMultipartFormDataParsed, err := parserTest.Parse(requestLineHeaderMultipartFormData)
	if err != nil {
		t.Error(err)
	}
	if len(requestLineHeaderMultipartFormDataParsed) != 0 {
		t.Errorf("Expected requestLineHeaderMultipartFormData to be skipped but got request")
	}

	// Request line in asterisk-form with one mandatory header
	requestLineAsteriskFormHeaderParsed, err := parserTest.Parse(requestLineAsteriskFormHeader)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"OPTIONS", "http://example.com:8080/*", "HTTP/1.1", "Host", "http://example.com:8080"}
	if len(requestLineAsteriskFormHeaderParsed) != 1 {
		t.Errorf("Expected requestLineAsteriskFormHeader to have one request but got %d",
			len(requestLineAsteriskFormHeaderParsed))
	}
	if requestLineAsteriskFormHeaderParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineAsteriskFormHeader to use method %s but got %s",
			assertValues[0], requestLineAsteriskFormHeaderParsed[0].Method)
	}
	if requestLineAsteriskFormHeaderParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineAsteriskFormHeader to use url %s but got %s",
			assertValues[1], requestLineAsteriskFormHeaderParsed[0].URL.String())
	}
	if requestLineAsteriskFormHeaderParsed[0].Proto != "HTTP/1.1" &&
		requestLineAsteriskFormHeaderParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineAsteriskFormHeader to use proto HTTP/1.1 or HTTP/2 but got %s",
			requestLineAsteriskFormHeaderParsed[0].Proto)
	}
	if len(requestLineAsteriskFormHeaderParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineAsteriskFormHeader to have one header but got %d",
			len(requestLineAsteriskFormHeaderParsed[0].Header))
	}
	for key, value := range requestLineAsteriskFormHeaderParsed[0].Header {
		if key != assertValues[3] {
			t.Errorf("Expected requestLineAsteriskFormHeader to have header key %s but got %s",
				assertValues[3], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineAsteriskFormHeader to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[4] {
			t.Errorf("Expected requestLineAsteriskFormHeader header value to be %s but got %s",
				assertValues[4], value[0])
		}
	}

	// Request line in origin-form (contains only path segment of url) with one mandatory header
	requestLineOriginFormHeaderParsed, err := parserTest.Parse(requestLineOriginFormHeader)
	if err != nil {
		t.Error(err)
	}
	assertValues = []string{"GET", "http://example.com/api/get", "/api/get", "Host", "example.com"}
	if len(requestLineOriginFormHeaderParsed) != 1 {
		t.Errorf("Expected requestLineOriginFormHeader to have one request but got %d",
			len(requestLineOriginFormHeaderParsed))
	}
	if requestLineOriginFormHeaderParsed[0].Method != assertValues[0] {
		t.Errorf("Expected requestLineOriginFormHeader to use method %s but got %s",
			assertValues[0], requestLineOriginFormHeaderParsed[0].Method)
	}
	if requestLineOriginFormHeaderParsed[0].URL.String() != assertValues[1] {
		t.Errorf("Expected requestLineOriginFormHeader to use url %s but got %s",
			assertValues[1], requestLineOriginFormHeaderParsed[0].URL.String())
	}
	if requestLineOriginFormHeaderParsed[0].URL.Path != assertValues[2] {
		t.Errorf("Expected requestLineOriginFormHeader to use url path %s but got %s",
			assertValues[1], requestLineOriginFormHeaderParsed[0].URL.String())
	}
	if requestLineOriginFormHeaderParsed[0].Proto != "HTTP/1.1" &&
		requestLineOriginFormHeaderParsed[0].Proto != "HTTP/2" {
		t.Errorf("Expected requestLineOriginFormHeader to use proto HTTP/1.1 or HTTP/2 but got %s",
			requestLineOriginFormHeaderParsed[0].Proto)
	}
	if len(requestLineOriginFormHeaderParsed[0].Header) != 1 {
		t.Errorf("Expected requestLineOriginFormHeader to have one header but got %d",
			len(requestLineOriginFormHeaderParsed[0].Header))
	}
	for key, value := range requestLineOriginFormHeaderParsed[0].Header {
		if key != assertValues[3] {
			t.Errorf("Expected requestLineOriginFormHeader to have header key %s but got %s",
				assertValues[3], key)
		}
		if len(value) != 1 {
			t.Errorf("Expected requestLineOriginFormHeader to have one header value but got %d", len(value))
		}
		if value[0] != assertValues[4] {
			t.Errorf("Expected requestLineOriginFormHeader header value to be %s but got %s",
				assertValues[4], value[0])
		}
	}
}

func TestPrepare(t *testing.T) {
	file, err := os.ReadFile("../../../test_files/default.http")
	if err != nil {
		t.Errorf("Could not open test file!\n\n%s", err.Error())
	}
	testDataForPreparing := string(file)

	var (
		regSeparator       = regexp.MustCompile(`(?m)^(###[^\r\n]*(\r?\n|\r))`)
		regComments        = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*(\r?\n|\r))`)
		regEnv             = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
		regResponseHandler = regexp.MustCompile(`(?m)^>[ \t\f]+([^\r\n]*(\r?\n|\r)$|{%(.|(\r?\n|\r))+%})(\r?\n|\r)`)
		regResponseRef     = regexp.MustCompile(`(?m)^<>[ \t\f]+[^\r\n]*$(\r?\n|\r)`)
	)

	err = parserTest.ParseConfig()
	if err != nil {
		t.Error(err)
	}
	requests, err := parserTest.Prepare(testDataForPreparing)
	if err != nil {
		t.Error(err)
	}

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
