// Code generated from ./src/lib/parser/httpspec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpspec

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 233,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 6, 3, 55, 10, 3, 13,
	3, 14, 3, 56, 3, 4, 3, 4, 7, 4, 61, 10, 4, 12, 4, 14, 4, 64, 11, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 5, 5, 5, 76,
	10, 5, 3, 5, 5, 5, 79, 10, 5, 3, 6, 3, 6, 5, 6, 83, 10, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 88, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 93, 10, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 99, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 104, 10, 8, 3, 8, 3,
	8, 5, 8, 108, 10, 8, 3, 9, 3, 9, 5, 9, 112, 10, 9, 3, 10, 3, 10, 3, 10,
	5, 10, 117, 10, 10, 3, 11, 3, 11, 6, 11, 121, 10, 11, 13, 11, 14, 11, 122,
	3, 11, 3, 11, 6, 11, 127, 10, 11, 13, 11, 14, 11, 128, 5, 11, 131, 10,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 137, 10, 12, 3, 13, 3, 13, 3, 13,
	7, 13, 142, 10, 13, 12, 13, 14, 13, 145, 11, 13, 3, 14, 6, 14, 148, 10,
	14, 13, 14, 14, 14, 149, 3, 14, 3, 14, 3, 14, 7, 14, 155, 10, 14, 12, 14,
	14, 14, 158, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 164, 10, 14, 12,
	14, 14, 14, 167, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 180, 10, 17, 3, 18, 3, 18, 3, 18, 5,
	18, 185, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 191, 10, 19, 3, 20,
	3, 20, 5, 20, 195, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 7,
	21, 203, 10, 21, 12, 21, 14, 21, 206, 11, 21, 3, 21, 3, 21, 5, 21, 210,
	10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 226, 10, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 2, 2, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 2, 4, 4, 2, 6, 6, 14,
	14, 3, 2, 3, 5, 2, 238, 2, 48, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 6, 58, 3,
	2, 2, 2, 8, 67, 3, 2, 2, 2, 10, 82, 3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14,
	98, 3, 2, 2, 2, 16, 109, 3, 2, 2, 2, 18, 113, 3, 2, 2, 2, 20, 130, 3, 2,
	2, 2, 22, 132, 3, 2, 2, 2, 24, 143, 3, 2, 2, 2, 26, 147, 3, 2, 2, 2, 28,
	170, 3, 2, 2, 2, 30, 173, 3, 2, 2, 2, 32, 179, 3, 2, 2, 2, 34, 181, 3,
	2, 2, 2, 36, 190, 3, 2, 2, 2, 38, 192, 3, 2, 2, 2, 40, 198, 3, 2, 2, 2,
	42, 211, 3, 2, 2, 2, 44, 225, 3, 2, 2, 2, 46, 227, 3, 2, 2, 2, 48, 49,
	5, 4, 3, 2, 49, 50, 7, 2, 2, 3, 50, 3, 3, 2, 2, 2, 51, 52, 5, 6, 4, 2,
	52, 53, 5, 8, 5, 2, 53, 55, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 55, 56, 3,
	2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 5, 3, 2, 2, 2, 58,
	62, 7, 3, 2, 2, 59, 61, 10, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 2,
	2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 65, 66, 7, 6, 2, 2, 66, 7, 3, 2, 2, 2, 67, 68, 5, 10, 6, 2,
	68, 69, 7, 6, 2, 2, 69, 70, 5, 24, 13, 2, 70, 72, 7, 6, 2, 2, 71, 73, 5,
	32, 17, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2,
	74, 76, 5, 44, 23, 2, 75, 74, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3,
	2, 2, 2, 77, 79, 5, 46, 24, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2,
	79, 9, 3, 2, 2, 2, 80, 81, 7, 20, 2, 2, 81, 83, 7, 12, 2, 2, 82, 80, 3,
	2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 5, 12, 7, 2, 85,
	86, 7, 12, 2, 2, 86, 88, 7, 28, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2,
	2, 2, 88, 11, 3, 2, 2, 2, 89, 93, 7, 22, 2, 2, 90, 93, 5, 14, 8, 2, 91,
	93, 7, 21, 2, 2, 92, 89, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 91, 3, 2,
	2, 2, 93, 13, 3, 2, 2, 2, 94, 95, 7, 25, 2, 2, 95, 96, 7, 16, 2, 2, 96,
	97, 7, 13, 2, 2, 97, 99, 7, 13, 2, 2, 98, 94, 3, 2, 2, 2, 98, 99, 3, 2,
	2, 2, 99, 100, 3, 2, 2, 2, 100, 103, 5, 16, 9, 2, 101, 102, 7, 17, 2, 2,
	102, 104, 7, 29, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	107, 3, 2, 2, 2, 105, 106, 7, 14, 2, 2, 106, 108, 7, 30, 2, 2, 107, 105,
	3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 15, 3, 2, 2, 2, 109, 111, 5, 18,
	10, 2, 110, 112, 7, 24, 2, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2,
	2, 112, 17, 3, 2, 2, 2, 113, 116, 5, 20, 11, 2, 114, 115, 7, 16, 2, 2,
	115, 117, 7, 26, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117,
	19, 3, 2, 2, 2, 118, 120, 7, 18, 2, 2, 119, 121, 7, 31, 2, 2, 120, 119,
	3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2,
	2, 2, 123, 124, 3, 2, 2, 2, 124, 131, 7, 19, 2, 2, 125, 127, 7, 32, 2,
	2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 118, 3, 2, 2, 2, 130, 126,
	3, 2, 2, 2, 131, 21, 3, 2, 2, 2, 132, 133, 7, 8, 2, 2, 133, 136, 7, 6,
	2, 2, 134, 135, 7, 7, 2, 2, 135, 137, 5, 22, 12, 2, 136, 134, 3, 2, 2,
	2, 136, 137, 3, 2, 2, 2, 137, 23, 3, 2, 2, 2, 138, 139, 5, 26, 14, 2, 139,
	140, 7, 6, 2, 2, 140, 142, 3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 142, 145,
	3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 25, 3, 2,
	2, 2, 145, 143, 3, 2, 2, 2, 146, 148, 7, 33, 2, 2, 147, 146, 3, 2, 2, 2,
	148, 149, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150,
	151, 3, 2, 2, 2, 151, 152, 7, 16, 2, 2, 152, 156, 8, 14, 1, 2, 153, 155,
	7, 12, 2, 2, 154, 153, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2,
	2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2,
	159, 160, 8, 14, 1, 2, 160, 161, 5, 22, 12, 2, 161, 165, 8, 14, 1, 2, 162,
	164, 7, 12, 2, 2, 163, 162, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163,
	3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 165, 3, 2,
	2, 2, 168, 169, 8, 14, 1, 2, 169, 27, 3, 2, 2, 2, 170, 171, 7, 8, 2, 2,
	171, 172, 7, 6, 2, 2, 172, 29, 3, 2, 2, 2, 173, 174, 7, 37, 2, 2, 174,
	175, 7, 12, 2, 2, 175, 176, 5, 28, 15, 2, 176, 31, 3, 2, 2, 2, 177, 180,
	5, 34, 18, 2, 178, 180, 5, 38, 20, 2, 179, 177, 3, 2, 2, 2, 179, 178, 3,
	2, 2, 2, 180, 33, 3, 2, 2, 2, 181, 184, 5, 36, 19, 2, 182, 183, 7, 6, 2,
	2, 183, 185, 5, 36, 19, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2,
	185, 35, 3, 2, 2, 2, 186, 187, 10, 3, 2, 2, 187, 188, 7, 8, 2, 2, 188,
	191, 7, 6, 2, 2, 189, 191, 5, 30, 16, 2, 190, 186, 3, 2, 2, 2, 190, 189,
	3, 2, 2, 2, 191, 37, 3, 2, 2, 2, 192, 194, 5, 40, 21, 2, 193, 195, 5, 38,
	20, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2,
	196, 197, 7, 39, 2, 2, 197, 39, 3, 2, 2, 2, 198, 204, 7, 39, 2, 2, 199,
	200, 5, 26, 14, 2, 200, 201, 7, 6, 2, 2, 201, 203, 3, 2, 2, 2, 202, 199,
	3, 2, 2, 2, 203, 206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205, 3, 2,
	2, 2, 205, 207, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 207, 209, 7, 6, 2, 2,
	208, 210, 5, 34, 18, 2, 209, 208, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210,
	41, 3, 2, 2, 2, 211, 212, 7, 8, 2, 2, 212, 213, 7, 6, 2, 2, 213, 43, 3,
	2, 2, 2, 214, 215, 7, 38, 2, 2, 215, 216, 7, 12, 2, 2, 216, 217, 7, 34,
	2, 2, 217, 218, 7, 36, 2, 2, 218, 219, 5, 42, 22, 2, 219, 220, 7, 36, 2,
	2, 220, 221, 7, 35, 2, 2, 221, 226, 3, 2, 2, 2, 222, 223, 7, 38, 2, 2,
	223, 224, 7, 12, 2, 2, 224, 226, 5, 28, 15, 2, 225, 214, 3, 2, 2, 2, 225,
	222, 3, 2, 2, 2, 226, 45, 3, 2, 2, 2, 227, 228, 7, 37, 2, 2, 228, 229,
	7, 38, 2, 2, 229, 230, 7, 12, 2, 2, 230, 231, 5, 28, 15, 2, 231, 47, 3,
	2, 2, 2, 30, 56, 62, 72, 75, 78, 82, 87, 92, 98, 103, 107, 111, 116, 122,
	128, 130, 136, 143, 149, 156, 165, 179, 184, 190, 194, 204, 209, 225,
}
var literalNames = []string{
	"", "'###'", "'<'", "'<>'", "", "", "", "", "", "", "", "'\u002F'", "'\u0023'",
	"", "'\u003A'", "'\u003F'", "'\u005B'", "'\u005D'", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "'\u007B'", "'\u007D'", "'\u0025'",
	"'\u003C'", "'\u003E'",
}
var symbolicNames = []string{
	"", "", "", "", "NEWLINE", "NEWLINEWITHINDENT", "INPUT", "ALPHA", "DIGIT",
	"ID", "WHITESPACE", "SLASH", "HASHTAG", "LINECOMMENT", "COLON", "QUESTIONMARK",
	"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "METHOD", "ASTERISKFORM", "ORIGINFORM",
	"SEGMENT", "ABSOLUTEPATH", "SCHEME", "PORT", "PATHSEPARATOR", "HTTPVERSION",
	"QUERY", "REQUESTFRAGMENT", "IPV6ADDRESS", "IPV4ADDRESSORREGNAME", "FIELDNAME",
	"LEFTCURLYBRACKET", "RIGHTCURLYBRACKET", "PERCENTSIGN", "LESSTHANSIGN",
	"GREATERTHANSIGN", "BOUNDARY",
}

var ruleNames = []string{
	"file", "requests", "requestseparator", "request", "requestline", "requesttarget",
	"absoluteform", "hierpart", "authority", "host", "fieldvalue", "headers",
	"headerfield", "filepath", "inputfileref", "messagebody", "messages", "messageline",
	"multipartformdata", "multipartfield", "handlerscript", "responsehandler",
	"responseref",
}

type httpspecParser struct {
	*antlr.BaseParser
}

// NewhttpspecParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *httpspecParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewhttpspecParser(input antlr.TokenStream) *httpspecParser {
	this := new(httpspecParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "httpspec.g4"

	return this
}

// httpspecParser tokens.
const (
	httpspecParserEOF                  = antlr.TokenEOF
	httpspecParserT__0                 = 1
	httpspecParserT__1                 = 2
	httpspecParserT__2                 = 3
	httpspecParserNEWLINE              = 4
	httpspecParserNEWLINEWITHINDENT    = 5
	httpspecParserINPUT                = 6
	httpspecParserALPHA                = 7
	httpspecParserDIGIT                = 8
	httpspecParserID                   = 9
	httpspecParserWHITESPACE           = 10
	httpspecParserSLASH                = 11
	httpspecParserHASHTAG              = 12
	httpspecParserLINECOMMENT          = 13
	httpspecParserCOLON                = 14
	httpspecParserQUESTIONMARK         = 15
	httpspecParserLEFTSQUAREBRACKET    = 16
	httpspecParserRIGHTSQUAREBRACKET   = 17
	httpspecParserMETHOD               = 18
	httpspecParserASTERISKFORM         = 19
	httpspecParserORIGINFORM           = 20
	httpspecParserSEGMENT              = 21
	httpspecParserABSOLUTEPATH         = 22
	httpspecParserSCHEME               = 23
	httpspecParserPORT                 = 24
	httpspecParserPATHSEPARATOR        = 25
	httpspecParserHTTPVERSION          = 26
	httpspecParserQUERY                = 27
	httpspecParserREQUESTFRAGMENT      = 28
	httpspecParserIPV6ADDRESS          = 29
	httpspecParserIPV4ADDRESSORREGNAME = 30
	httpspecParserFIELDNAME            = 31
	httpspecParserLEFTCURLYBRACKET     = 32
	httpspecParserRIGHTCURLYBRACKET    = 33
	httpspecParserPERCENTSIGN          = 34
	httpspecParserLESSTHANSIGN         = 35
	httpspecParserGREATERTHANSIGN      = 36
	httpspecParserBOUNDARY             = 37
)

// httpspecParser rules.
const (
	httpspecParserRULE_file              = 0
	httpspecParserRULE_requests          = 1
	httpspecParserRULE_requestseparator  = 2
	httpspecParserRULE_request           = 3
	httpspecParserRULE_requestline       = 4
	httpspecParserRULE_requesttarget     = 5
	httpspecParserRULE_absoluteform      = 6
	httpspecParserRULE_hierpart          = 7
	httpspecParserRULE_authority         = 8
	httpspecParserRULE_host              = 9
	httpspecParserRULE_fieldvalue        = 10
	httpspecParserRULE_headers           = 11
	httpspecParserRULE_headerfield       = 12
	httpspecParserRULE_filepath          = 13
	httpspecParserRULE_inputfileref      = 14
	httpspecParserRULE_messagebody       = 15
	httpspecParserRULE_messages          = 16
	httpspecParserRULE_messageline       = 17
	httpspecParserRULE_multipartformdata = 18
	httpspecParserRULE_multipartfield    = 19
	httpspecParserRULE_handlerscript     = 20
	httpspecParserRULE_responsehandler   = 21
	httpspecParserRULE_responseref       = 22
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Requests() IRequestsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestsContext)
}

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(httpspecParserEOF, 0)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *httpspecParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, httpspecParserRULE_file)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Requests()
	}
	{
		p.SetState(47)
		p.Match(httpspecParserEOF)
	}

	return localctx
}

// IRequestsContext is an interface to support dynamic dispatch.
type IRequestsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestsContext differentiates from other interfaces.
	IsRequestsContext()
}

type RequestsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestsContext() *RequestsContext {
	var p = new(RequestsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_requests
	return p
}

func (*RequestsContext) IsRequestsContext() {}

func NewRequestsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestsContext {
	var p = new(RequestsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_requests

	return p
}

func (s *RequestsContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestsContext) AllRequestseparator() []IRequestseparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRequestseparatorContext)(nil)).Elem())
	var tst = make([]IRequestseparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRequestseparatorContext)
		}
	}

	return tst
}

func (s *RequestsContext) Requestseparator(i int) IRequestseparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestseparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRequestseparatorContext)
}

func (s *RequestsContext) AllRequest() []IRequestContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRequestContext)(nil)).Elem())
	var tst = make([]IRequestContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRequestContext)
		}
	}

	return tst
}

func (s *RequestsContext) Request(i int) IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RequestsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterRequests(s)
	}
}

func (s *RequestsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitRequests(s)
	}
}

func (p *httpspecParser) Requests() (localctx IRequestsContext) {
	localctx = NewRequestsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, httpspecParserRULE_requests)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpspecParserT__0 {
		{
			p.SetState(49)
			p.Requestseparator()
		}
		{
			p.SetState(50)
			p.Request()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRequestseparatorContext is an interface to support dynamic dispatch.
type IRequestseparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestseparatorContext differentiates from other interfaces.
	IsRequestseparatorContext()
}

type RequestseparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestseparatorContext() *RequestseparatorContext {
	var p = new(RequestseparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_requestseparator
	return p
}

func (*RequestseparatorContext) IsRequestseparatorContext() {}

func NewRequestseparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestseparatorContext {
	var p = new(RequestseparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_requestseparator

	return p
}

func (s *RequestseparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestseparatorContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserNEWLINE)
}

func (s *RequestseparatorContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, i)
}

func (s *RequestseparatorContext) AllHASHTAG() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserHASHTAG)
}

func (s *RequestseparatorContext) HASHTAG(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserHASHTAG, i)
}

func (s *RequestseparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestseparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestseparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterRequestseparator(s)
	}
}

func (s *RequestseparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitRequestseparator(s)
	}
}

func (p *httpspecParser) Requestseparator() (localctx IRequestseparatorContext) {
	localctx = NewRequestseparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, httpspecParserRULE_requestseparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(httpspecParserT__0)
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpspecParserT__0)|(1<<httpspecParserT__1)|(1<<httpspecParserT__2)|(1<<httpspecParserNEWLINEWITHINDENT)|(1<<httpspecParserINPUT)|(1<<httpspecParserALPHA)|(1<<httpspecParserDIGIT)|(1<<httpspecParserID)|(1<<httpspecParserWHITESPACE)|(1<<httpspecParserSLASH)|(1<<httpspecParserLINECOMMENT)|(1<<httpspecParserCOLON)|(1<<httpspecParserQUESTIONMARK)|(1<<httpspecParserLEFTSQUAREBRACKET)|(1<<httpspecParserRIGHTSQUAREBRACKET)|(1<<httpspecParserMETHOD)|(1<<httpspecParserASTERISKFORM)|(1<<httpspecParserORIGINFORM)|(1<<httpspecParserSEGMENT)|(1<<httpspecParserABSOLUTEPATH)|(1<<httpspecParserSCHEME)|(1<<httpspecParserPORT)|(1<<httpspecParserPATHSEPARATOR)|(1<<httpspecParserHTTPVERSION)|(1<<httpspecParserQUERY)|(1<<httpspecParserREQUESTFRAGMENT)|(1<<httpspecParserIPV6ADDRESS)|(1<<httpspecParserIPV4ADDRESSORREGNAME)|(1<<httpspecParserFIELDNAME))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(httpspecParserLEFTCURLYBRACKET-32))|(1<<(httpspecParserRIGHTCURLYBRACKET-32))|(1<<(httpspecParserPERCENTSIGN-32))|(1<<(httpspecParserLESSTHANSIGN-32))|(1<<(httpspecParserGREATERTHANSIGN-32))|(1<<(httpspecParserBOUNDARY-32)))) != 0) {
		{
			p.SetState(57)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == httpspecParserNEWLINE || _la == httpspecParserHASHTAG {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(httpspecParserNEWLINE)
	}

	return localctx
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) Requestline() IRequestlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestlineContext)
}

func (s *RequestContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserNEWLINE)
}

func (s *RequestContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, i)
}

func (s *RequestContext) Headers() IHeadersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeadersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeadersContext)
}

func (s *RequestContext) Messagebody() IMessagebodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagebodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessagebodyContext)
}

func (s *RequestContext) Responsehandler() IResponsehandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResponsehandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResponsehandlerContext)
}

func (s *RequestContext) Responseref() IResponserefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResponserefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResponserefContext)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (p *httpspecParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, httpspecParserRULE_request)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Requestline()
	}
	{
		p.SetState(66)
		p.Match(httpspecParserNEWLINE)
	}
	{
		p.SetState(67)
		p.Headers()
	}
	{
		p.SetState(68)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(69)
			p.Messagebody()
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserGREATERTHANSIGN {
		{
			p.SetState(72)
			p.Responsehandler()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserLESSTHANSIGN {
		{
			p.SetState(75)
			p.Responseref()
		}

	}

	return localctx
}

// IRequestlineContext is an interface to support dynamic dispatch.
type IRequestlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestlineContext differentiates from other interfaces.
	IsRequestlineContext()
}

type RequestlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestlineContext() *RequestlineContext {
	var p = new(RequestlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_requestline
	return p
}

func (*RequestlineContext) IsRequestlineContext() {}

func NewRequestlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestlineContext {
	var p = new(RequestlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_requestline

	return p
}

func (s *RequestlineContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestlineContext) Requesttarget() IRequesttargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequesttargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequesttargetContext)
}

func (s *RequestlineContext) METHOD() antlr.TerminalNode {
	return s.GetToken(httpspecParserMETHOD, 0)
}

func (s *RequestlineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserWHITESPACE)
}

func (s *RequestlineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserWHITESPACE, i)
}

func (s *RequestlineContext) HTTPVERSION() antlr.TerminalNode {
	return s.GetToken(httpspecParserHTTPVERSION, 0)
}

func (s *RequestlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterRequestline(s)
	}
}

func (s *RequestlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitRequestline(s)
	}
}

func (p *httpspecParser) Requestline() (localctx IRequestlineContext) {
	localctx = NewRequestlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, httpspecParserRULE_requestline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserMETHOD {
		{
			p.SetState(78)
			p.Match(httpspecParserMETHOD)
		}
		{
			p.SetState(79)
			p.Match(httpspecParserWHITESPACE)
		}

	}
	{
		p.SetState(82)
		p.Requesttarget()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserWHITESPACE {
		{
			p.SetState(83)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(84)
			p.Match(httpspecParserHTTPVERSION)
		}

	}

	return localctx
}

// IRequesttargetContext is an interface to support dynamic dispatch.
type IRequesttargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequesttargetContext differentiates from other interfaces.
	IsRequesttargetContext()
}

type RequesttargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequesttargetContext() *RequesttargetContext {
	var p = new(RequesttargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_requesttarget
	return p
}

func (*RequesttargetContext) IsRequesttargetContext() {}

func NewRequesttargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequesttargetContext {
	var p = new(RequesttargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_requesttarget

	return p
}

func (s *RequesttargetContext) GetParser() antlr.Parser { return s.parser }

func (s *RequesttargetContext) ORIGINFORM() antlr.TerminalNode {
	return s.GetToken(httpspecParserORIGINFORM, 0)
}

func (s *RequesttargetContext) Absoluteform() IAbsoluteformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbsoluteformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbsoluteformContext)
}

func (s *RequesttargetContext) ASTERISKFORM() antlr.TerminalNode {
	return s.GetToken(httpspecParserASTERISKFORM, 0)
}

func (s *RequesttargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequesttargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequesttargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterRequesttarget(s)
	}
}

func (s *RequesttargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitRequesttarget(s)
	}
}

func (p *httpspecParser) Requesttarget() (localctx IRequesttargetContext) {
	localctx = NewRequesttargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, httpspecParserRULE_requesttarget)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpspecParserORIGINFORM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(httpspecParserORIGINFORM)
		}

	case httpspecParserLEFTSQUAREBRACKET, httpspecParserSCHEME, httpspecParserIPV4ADDRESSORREGNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Absoluteform()
		}

	case httpspecParserASTERISKFORM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(httpspecParserASTERISKFORM)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAbsoluteformContext is an interface to support dynamic dispatch.
type IAbsoluteformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsoluteformContext differentiates from other interfaces.
	IsAbsoluteformContext()
}

type AbsoluteformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsoluteformContext() *AbsoluteformContext {
	var p = new(AbsoluteformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_absoluteform
	return p
}

func (*AbsoluteformContext) IsAbsoluteformContext() {}

func NewAbsoluteformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsoluteformContext {
	var p = new(AbsoluteformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_absoluteform

	return p
}

func (s *AbsoluteformContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsoluteformContext) Hierpart() IHierpartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHierpartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHierpartContext)
}

func (s *AbsoluteformContext) SCHEME() antlr.TerminalNode {
	return s.GetToken(httpspecParserSCHEME, 0)
}

func (s *AbsoluteformContext) COLON() antlr.TerminalNode {
	return s.GetToken(httpspecParserCOLON, 0)
}

func (s *AbsoluteformContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserSLASH)
}

func (s *AbsoluteformContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserSLASH, i)
}

func (s *AbsoluteformContext) QUESTIONMARK() antlr.TerminalNode {
	return s.GetToken(httpspecParserQUESTIONMARK, 0)
}

func (s *AbsoluteformContext) QUERY() antlr.TerminalNode {
	return s.GetToken(httpspecParserQUERY, 0)
}

func (s *AbsoluteformContext) HASHTAG() antlr.TerminalNode {
	return s.GetToken(httpspecParserHASHTAG, 0)
}

func (s *AbsoluteformContext) REQUESTFRAGMENT() antlr.TerminalNode {
	return s.GetToken(httpspecParserREQUESTFRAGMENT, 0)
}

func (s *AbsoluteformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsoluteformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsoluteformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterAbsoluteform(s)
	}
}

func (s *AbsoluteformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitAbsoluteform(s)
	}
}

func (p *httpspecParser) Absoluteform() (localctx IAbsoluteformContext) {
	localctx = NewAbsoluteformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, httpspecParserRULE_absoluteform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserSCHEME {
		{
			p.SetState(92)
			p.Match(httpspecParserSCHEME)
		}
		{
			p.SetState(93)
			p.Match(httpspecParserCOLON)
		}
		{
			p.SetState(94)
			p.Match(httpspecParserSLASH)
		}
		{
			p.SetState(95)
			p.Match(httpspecParserSLASH)
		}

	}
	{
		p.SetState(98)
		p.Hierpart()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserQUESTIONMARK {
		{
			p.SetState(99)
			p.Match(httpspecParserQUESTIONMARK)
		}
		{
			p.SetState(100)
			p.Match(httpspecParserQUERY)
		}

	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserHASHTAG {
		{
			p.SetState(103)
			p.Match(httpspecParserHASHTAG)
		}
		{
			p.SetState(104)
			p.Match(httpspecParserREQUESTFRAGMENT)
		}

	}

	return localctx
}

// IHierpartContext is an interface to support dynamic dispatch.
type IHierpartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHierpartContext differentiates from other interfaces.
	IsHierpartContext()
}

type HierpartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHierpartContext() *HierpartContext {
	var p = new(HierpartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_hierpart
	return p
}

func (*HierpartContext) IsHierpartContext() {}

func NewHierpartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HierpartContext {
	var p = new(HierpartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_hierpart

	return p
}

func (s *HierpartContext) GetParser() antlr.Parser { return s.parser }

func (s *HierpartContext) Authority() IAuthorityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuthorityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *HierpartContext) ABSOLUTEPATH() antlr.TerminalNode {
	return s.GetToken(httpspecParserABSOLUTEPATH, 0)
}

func (s *HierpartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HierpartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HierpartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterHierpart(s)
	}
}

func (s *HierpartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitHierpart(s)
	}
}

func (p *httpspecParser) Hierpart() (localctx IHierpartContext) {
	localctx = NewHierpartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, httpspecParserRULE_hierpart)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Authority()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserABSOLUTEPATH {
		{
			p.SetState(108)
			p.Match(httpspecParserABSOLUTEPATH)
		}

	}

	return localctx
}

// IAuthorityContext is an interface to support dynamic dispatch.
type IAuthorityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuthorityContext differentiates from other interfaces.
	IsAuthorityContext()
}

type AuthorityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorityContext() *AuthorityContext {
	var p = new(AuthorityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_authority
	return p
}

func (*AuthorityContext) IsAuthorityContext() {}

func NewAuthorityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorityContext {
	var p = new(AuthorityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_authority

	return p
}

func (s *AuthorityContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthorityContext) Host() IHostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHostContext)
}

func (s *AuthorityContext) COLON() antlr.TerminalNode {
	return s.GetToken(httpspecParserCOLON, 0)
}

func (s *AuthorityContext) PORT() antlr.TerminalNode {
	return s.GetToken(httpspecParserPORT, 0)
}

func (s *AuthorityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuthorityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterAuthority(s)
	}
}

func (s *AuthorityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitAuthority(s)
	}
}

func (p *httpspecParser) Authority() (localctx IAuthorityContext) {
	localctx = NewAuthorityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, httpspecParserRULE_authority)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Host()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserCOLON {
		{
			p.SetState(112)
			p.Match(httpspecParserCOLON)
		}
		{
			p.SetState(113)
			p.Match(httpspecParserPORT)
		}

	}

	return localctx
}

// IHostContext is an interface to support dynamic dispatch.
type IHostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostContext differentiates from other interfaces.
	IsHostContext()
}

type HostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostContext() *HostContext {
	var p = new(HostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_host
	return p
}

func (*HostContext) IsHostContext() {}

func NewHostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostContext {
	var p = new(HostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_host

	return p
}

func (s *HostContext) GetParser() antlr.Parser { return s.parser }

func (s *HostContext) LEFTSQUAREBRACKET() antlr.TerminalNode {
	return s.GetToken(httpspecParserLEFTSQUAREBRACKET, 0)
}

func (s *HostContext) RIGHTSQUAREBRACKET() antlr.TerminalNode {
	return s.GetToken(httpspecParserRIGHTSQUAREBRACKET, 0)
}

func (s *HostContext) AllIPV6ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserIPV6ADDRESS)
}

func (s *HostContext) IPV6ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserIPV6ADDRESS, i)
}

func (s *HostContext) AllIPV4ADDRESSORREGNAME() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserIPV4ADDRESSORREGNAME)
}

func (s *HostContext) IPV4ADDRESSORREGNAME(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserIPV4ADDRESSORREGNAME, i)
}

func (s *HostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterHost(s)
	}
}

func (s *HostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitHost(s)
	}
}

func (p *httpspecParser) Host() (localctx IHostContext) {
	localctx = NewHostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, httpspecParserRULE_host)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpspecParserLEFTSQUAREBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(httpspecParserLEFTSQUAREBRACKET)
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == httpspecParserIPV6ADDRESS {
			{
				p.SetState(117)
				p.Match(httpspecParserIPV6ADDRESS)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(httpspecParserRIGHTSQUAREBRACKET)
		}

	case httpspecParserIPV4ADDRESSORREGNAME:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == httpspecParserIPV4ADDRESSORREGNAME {
			{
				p.SetState(123)
				p.Match(httpspecParserIPV4ADDRESSORREGNAME)
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldvalueContext is an interface to support dynamic dispatch.
type IFieldvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldvalueContext differentiates from other interfaces.
	IsFieldvalueContext()
}

type FieldvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldvalueContext() *FieldvalueContext {
	var p = new(FieldvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_fieldvalue
	return p
}

func (*FieldvalueContext) IsFieldvalueContext() {}

func NewFieldvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldvalueContext {
	var p = new(FieldvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_fieldvalue

	return p
}

func (s *FieldvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldvalueContext) INPUT() antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, 0)
}

func (s *FieldvalueContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *FieldvalueContext) NEWLINEWITHINDENT() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINEWITHINDENT, 0)
}

func (s *FieldvalueContext) Fieldvalue() IFieldvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldvalueContext)
}

func (s *FieldvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterFieldvalue(s)
	}
}

func (s *FieldvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitFieldvalue(s)
	}
}

func (p *httpspecParser) Fieldvalue() (localctx IFieldvalueContext) {
	localctx = NewFieldvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, httpspecParserRULE_fieldvalue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(httpspecParserINPUT)
	}
	{
		p.SetState(131)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserNEWLINEWITHINDENT {
		{
			p.SetState(132)
			p.Match(httpspecParserNEWLINEWITHINDENT)
		}
		{
			p.SetState(133)
			p.Fieldvalue()
		}

	}

	return localctx
}

// IHeadersContext is an interface to support dynamic dispatch.
type IHeadersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeadersContext differentiates from other interfaces.
	IsHeadersContext()
}

type HeadersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeadersContext() *HeadersContext {
	var p = new(HeadersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_headers
	return p
}

func (*HeadersContext) IsHeadersContext() {}

func NewHeadersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadersContext {
	var p = new(HeadersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_headers

	return p
}

func (s *HeadersContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadersContext) AllHeaderfield() []IHeaderfieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem())
	var tst = make([]IHeaderfieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderfieldContext)
		}
	}

	return tst
}

func (s *HeadersContext) Headerfield(i int) IHeaderfieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderfieldContext)
}

func (s *HeadersContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserNEWLINE)
}

func (s *HeadersContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, i)
}

func (s *HeadersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterHeaders(s)
	}
}

func (s *HeadersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitHeaders(s)
	}
}

func (p *httpspecParser) Headers() (localctx IHeadersContext) {
	localctx = NewHeadersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, httpspecParserRULE_headers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserFIELDNAME {
		{
			p.SetState(136)
			p.Headerfield()
		}
		{
			p.SetState(137)
			p.Match(httpspecParserNEWLINE)
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderfieldContext is an interface to support dynamic dispatch.
type IHeaderfieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderfieldContext differentiates from other interfaces.
	IsHeaderfieldContext()
}

type HeaderfieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderfieldContext() *HeaderfieldContext {
	var p = new(HeaderfieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_headerfield
	return p
}

func (*HeaderfieldContext) IsHeaderfieldContext() {}

func NewHeaderfieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderfieldContext {
	var p = new(HeaderfieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_headerfield

	return p
}

func (s *HeaderfieldContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderfieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(httpspecParserCOLON, 0)
}

func (s *HeaderfieldContext) Fieldvalue() IFieldvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldvalueContext)
}

func (s *HeaderfieldContext) AllFIELDNAME() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserFIELDNAME)
}

func (s *HeaderfieldContext) FIELDNAME(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserFIELDNAME, i)
}

func (s *HeaderfieldContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserWHITESPACE)
}

func (s *HeaderfieldContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserWHITESPACE, i)
}

func (s *HeaderfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderfieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterHeaderfield(s)
	}
}

func (s *HeaderfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitHeaderfield(s)
	}
}

func (p *httpspecParser) Headerfield() (localctx IHeaderfieldContext) {
	localctx = NewHeaderfieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, httpspecParserRULE_headerfield)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpspecParserFIELDNAME {
		{
			p.SetState(144)
			p.Match(httpspecParserFIELDNAME)
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Match(httpspecParserCOLON)
	}
	ignoreWs = true
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserWHITESPACE {
		{
			p.SetState(151)
			p.Match(httpspecParserWHITESPACE)
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	ignoreWs = false
	{
		p.SetState(158)
		p.Fieldvalue()
	}
	ignoreWs = true
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserWHITESPACE {
		{
			p.SetState(160)
			p.Match(httpspecParserWHITESPACE)
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	ignoreWs = false

	return localctx
}

// IFilepathContext is an interface to support dynamic dispatch.
type IFilepathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilepathContext differentiates from other interfaces.
	IsFilepathContext()
}

type FilepathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilepathContext() *FilepathContext {
	var p = new(FilepathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_filepath
	return p
}

func (*FilepathContext) IsFilepathContext() {}

func NewFilepathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilepathContext {
	var p = new(FilepathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_filepath

	return p
}

func (s *FilepathContext) GetParser() antlr.Parser { return s.parser }

func (s *FilepathContext) INPUT() antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, 0)
}

func (s *FilepathContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *FilepathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilepathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilepathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterFilepath(s)
	}
}

func (s *FilepathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitFilepath(s)
	}
}

func (p *httpspecParser) Filepath() (localctx IFilepathContext) {
	localctx = NewFilepathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, httpspecParserRULE_filepath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(httpspecParserINPUT)
	}
	{
		p.SetState(169)
		p.Match(httpspecParserNEWLINE)
	}

	return localctx
}

// IInputfilerefContext is an interface to support dynamic dispatch.
type IInputfilerefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputfilerefContext differentiates from other interfaces.
	IsInputfilerefContext()
}

type InputfilerefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputfilerefContext() *InputfilerefContext {
	var p = new(InputfilerefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_inputfileref
	return p
}

func (*InputfilerefContext) IsInputfilerefContext() {}

func NewInputfilerefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputfilerefContext {
	var p = new(InputfilerefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_inputfileref

	return p
}

func (s *InputfilerefContext) GetParser() antlr.Parser { return s.parser }

func (s *InputfilerefContext) LESSTHANSIGN() antlr.TerminalNode {
	return s.GetToken(httpspecParserLESSTHANSIGN, 0)
}

func (s *InputfilerefContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(httpspecParserWHITESPACE, 0)
}

func (s *InputfilerefContext) Filepath() IFilepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilepathContext)
}

func (s *InputfilerefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputfilerefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputfilerefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterInputfileref(s)
	}
}

func (s *InputfilerefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitInputfileref(s)
	}
}

func (p *httpspecParser) Inputfileref() (localctx IInputfilerefContext) {
	localctx = NewInputfilerefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, httpspecParserRULE_inputfileref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(httpspecParserLESSTHANSIGN)
	}
	{
		p.SetState(172)
		p.Match(httpspecParserWHITESPACE)
	}
	{
		p.SetState(173)
		p.Filepath()
	}

	return localctx
}

// IMessagebodyContext is an interface to support dynamic dispatch.
type IMessagebodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessagebodyContext differentiates from other interfaces.
	IsMessagebodyContext()
}

type MessagebodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagebodyContext() *MessagebodyContext {
	var p = new(MessagebodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_messagebody
	return p
}

func (*MessagebodyContext) IsMessagebodyContext() {}

func NewMessagebodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagebodyContext {
	var p = new(MessagebodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_messagebody

	return p
}

func (s *MessagebodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagebodyContext) Messages() IMessagesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessagesContext)
}

func (s *MessagebodyContext) Multipartformdata() IMultipartformdataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultipartformdataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultipartformdataContext)
}

func (s *MessagebodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagebodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagebodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterMessagebody(s)
	}
}

func (s *MessagebodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitMessagebody(s)
	}
}

func (p *httpspecParser) Messagebody() (localctx IMessagebodyContext) {
	localctx = NewMessagebodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, httpspecParserRULE_messagebody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Messages()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Multipartformdata()
		}

	}

	return localctx
}

// IMessagesContext is an interface to support dynamic dispatch.
type IMessagesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessagesContext differentiates from other interfaces.
	IsMessagesContext()
}

type MessagesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagesContext() *MessagesContext {
	var p = new(MessagesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_messages
	return p
}

func (*MessagesContext) IsMessagesContext() {}

func NewMessagesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagesContext {
	var p = new(MessagesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_messages

	return p
}

func (s *MessagesContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagesContext) AllMessageline() []IMessagelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessagelineContext)(nil)).Elem())
	var tst = make([]IMessagelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessagelineContext)
		}
	}

	return tst
}

func (s *MessagesContext) Messageline(i int) IMessagelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessagelineContext)
}

func (s *MessagesContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *MessagesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterMessages(s)
	}
}

func (s *MessagesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitMessages(s)
	}
}

func (p *httpspecParser) Messages() (localctx IMessagesContext) {
	localctx = NewMessagesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, httpspecParserRULE_messages)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Messageline()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserNEWLINE {
		{
			p.SetState(180)
			p.Match(httpspecParserNEWLINE)
		}
		{
			p.SetState(181)
			p.Messageline()
		}

	}

	return localctx
}

// IMessagelineContext is an interface to support dynamic dispatch.
type IMessagelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessagelineContext differentiates from other interfaces.
	IsMessagelineContext()
}

type MessagelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagelineContext() *MessagelineContext {
	var p = new(MessagelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_messageline
	return p
}

func (*MessagelineContext) IsMessagelineContext() {}

func NewMessagelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagelineContext {
	var p = new(MessagelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_messageline

	return p
}

func (s *MessagelineContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagelineContext) INPUT() antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, 0)
}

func (s *MessagelineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *MessagelineContext) Inputfileref() IInputfilerefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputfilerefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputfilerefContext)
}

func (s *MessagelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterMessageline(s)
	}
}

func (s *MessagelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitMessageline(s)
	}
}

func (p *httpspecParser) Messageline() (localctx IMessagelineContext) {
	localctx = NewMessagelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, httpspecParserRULE_messageline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpspecParserT__0)|(1<<httpspecParserT__1)|(1<<httpspecParserT__2))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(185)
			p.Match(httpspecParserINPUT)
		}
		{
			p.SetState(186)
			p.Match(httpspecParserNEWLINE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Inputfileref()
		}

	}

	return localctx
}

// IMultipartformdataContext is an interface to support dynamic dispatch.
type IMultipartformdataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultipartformdataContext differentiates from other interfaces.
	IsMultipartformdataContext()
}

type MultipartformdataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipartformdataContext() *MultipartformdataContext {
	var p = new(MultipartformdataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_multipartformdata
	return p
}

func (*MultipartformdataContext) IsMultipartformdataContext() {}

func NewMultipartformdataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipartformdataContext {
	var p = new(MultipartformdataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_multipartformdata

	return p
}

func (s *MultipartformdataContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipartformdataContext) Multipartfield() IMultipartfieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultipartfieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultipartfieldContext)
}

func (s *MultipartformdataContext) BOUNDARY() antlr.TerminalNode {
	return s.GetToken(httpspecParserBOUNDARY, 0)
}

func (s *MultipartformdataContext) Multipartformdata() IMultipartformdataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultipartformdataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultipartformdataContext)
}

func (s *MultipartformdataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipartformdataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipartformdataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterMultipartformdata(s)
	}
}

func (s *MultipartformdataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitMultipartformdata(s)
	}
}

func (p *httpspecParser) Multipartformdata() (localctx IMultipartformdataContext) {
	localctx = NewMultipartformdataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, httpspecParserRULE_multipartformdata)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Multipartfield()
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(191)
			p.Multipartformdata()
		}

	}
	{
		p.SetState(194)
		p.Match(httpspecParserBOUNDARY)
	}

	return localctx
}

// IMultipartfieldContext is an interface to support dynamic dispatch.
type IMultipartfieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultipartfieldContext differentiates from other interfaces.
	IsMultipartfieldContext()
}

type MultipartfieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipartfieldContext() *MultipartfieldContext {
	var p = new(MultipartfieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_multipartfield
	return p
}

func (*MultipartfieldContext) IsMultipartfieldContext() {}

func NewMultipartfieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipartfieldContext {
	var p = new(MultipartfieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_multipartfield

	return p
}

func (s *MultipartfieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipartfieldContext) BOUNDARY() antlr.TerminalNode {
	return s.GetToken(httpspecParserBOUNDARY, 0)
}

func (s *MultipartfieldContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserNEWLINE)
}

func (s *MultipartfieldContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, i)
}

func (s *MultipartfieldContext) AllHeaderfield() []IHeaderfieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem())
	var tst = make([]IHeaderfieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderfieldContext)
		}
	}

	return tst
}

func (s *MultipartfieldContext) Headerfield(i int) IHeaderfieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderfieldContext)
}

func (s *MultipartfieldContext) Messages() IMessagesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessagesContext)
}

func (s *MultipartfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipartfieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipartfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterMultipartfield(s)
	}
}

func (s *MultipartfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitMultipartfield(s)
	}
}

func (p *httpspecParser) Multipartfield() (localctx IMultipartfieldContext) {
	localctx = NewMultipartfieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, httpspecParserRULE_multipartfield)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(httpspecParserBOUNDARY)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserFIELDNAME {
		{
			p.SetState(197)
			p.Headerfield()
		}
		{
			p.SetState(198)
			p.Match(httpspecParserNEWLINE)
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(206)
			p.Messages()
		}

	}

	return localctx
}

// IHandlerscriptContext is an interface to support dynamic dispatch.
type IHandlerscriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHandlerscriptContext differentiates from other interfaces.
	IsHandlerscriptContext()
}

type HandlerscriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHandlerscriptContext() *HandlerscriptContext {
	var p = new(HandlerscriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_handlerscript
	return p
}

func (*HandlerscriptContext) IsHandlerscriptContext() {}

func NewHandlerscriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HandlerscriptContext {
	var p = new(HandlerscriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_handlerscript

	return p
}

func (s *HandlerscriptContext) GetParser() antlr.Parser { return s.parser }

func (s *HandlerscriptContext) INPUT() antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, 0)
}

func (s *HandlerscriptContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *HandlerscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandlerscriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HandlerscriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterHandlerscript(s)
	}
}

func (s *HandlerscriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitHandlerscript(s)
	}
}

func (p *httpspecParser) Handlerscript() (localctx IHandlerscriptContext) {
	localctx = NewHandlerscriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, httpspecParserRULE_handlerscript)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(httpspecParserINPUT)
	}
	{
		p.SetState(210)
		p.Match(httpspecParserNEWLINE)
	}

	return localctx
}

// IResponsehandlerContext is an interface to support dynamic dispatch.
type IResponsehandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResponsehandlerContext differentiates from other interfaces.
	IsResponsehandlerContext()
}

type ResponsehandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponsehandlerContext() *ResponsehandlerContext {
	var p = new(ResponsehandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_responsehandler
	return p
}

func (*ResponsehandlerContext) IsResponsehandlerContext() {}

func NewResponsehandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponsehandlerContext {
	var p = new(ResponsehandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_responsehandler

	return p
}

func (s *ResponsehandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponsehandlerContext) GREATERTHANSIGN() antlr.TerminalNode {
	return s.GetToken(httpspecParserGREATERTHANSIGN, 0)
}

func (s *ResponsehandlerContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(httpspecParserWHITESPACE, 0)
}

func (s *ResponsehandlerContext) LEFTCURLYBRACKET() antlr.TerminalNode {
	return s.GetToken(httpspecParserLEFTCURLYBRACKET, 0)
}

func (s *ResponsehandlerContext) AllPERCENTSIGN() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserPERCENTSIGN)
}

func (s *ResponsehandlerContext) PERCENTSIGN(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserPERCENTSIGN, i)
}

func (s *ResponsehandlerContext) Handlerscript() IHandlerscriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHandlerscriptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHandlerscriptContext)
}

func (s *ResponsehandlerContext) RIGHTCURLYBRACKET() antlr.TerminalNode {
	return s.GetToken(httpspecParserRIGHTCURLYBRACKET, 0)
}

func (s *ResponsehandlerContext) Filepath() IFilepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilepathContext)
}

func (s *ResponsehandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponsehandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponsehandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterResponsehandler(s)
	}
}

func (s *ResponsehandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitResponsehandler(s)
	}
}

func (p *httpspecParser) Responsehandler() (localctx IResponsehandlerContext) {
	localctx = NewResponsehandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, httpspecParserRULE_responsehandler)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(httpspecParserGREATERTHANSIGN)
		}
		{
			p.SetState(213)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(214)
			p.Match(httpspecParserLEFTCURLYBRACKET)
		}
		{
			p.SetState(215)
			p.Match(httpspecParserPERCENTSIGN)
		}
		{
			p.SetState(216)
			p.Handlerscript()
		}
		{
			p.SetState(217)
			p.Match(httpspecParserPERCENTSIGN)
		}
		{
			p.SetState(218)
			p.Match(httpspecParserRIGHTCURLYBRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Match(httpspecParserGREATERTHANSIGN)
		}
		{
			p.SetState(221)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(222)
			p.Filepath()
		}

	}

	return localctx
}

// IResponserefContext is an interface to support dynamic dispatch.
type IResponserefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResponserefContext differentiates from other interfaces.
	IsResponserefContext()
}

type ResponserefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponserefContext() *ResponserefContext {
	var p = new(ResponserefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_responseref
	return p
}

func (*ResponserefContext) IsResponserefContext() {}

func NewResponserefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponserefContext {
	var p = new(ResponserefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_responseref

	return p
}

func (s *ResponserefContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponserefContext) LESSTHANSIGN() antlr.TerminalNode {
	return s.GetToken(httpspecParserLESSTHANSIGN, 0)
}

func (s *ResponserefContext) GREATERTHANSIGN() antlr.TerminalNode {
	return s.GetToken(httpspecParserGREATERTHANSIGN, 0)
}

func (s *ResponserefContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(httpspecParserWHITESPACE, 0)
}

func (s *ResponserefContext) Filepath() IFilepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilepathContext)
}

func (s *ResponserefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponserefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponserefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterResponseref(s)
	}
}

func (s *ResponserefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitResponseref(s)
	}
}

func (p *httpspecParser) Responseref() (localctx IResponserefContext) {
	localctx = NewResponserefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, httpspecParserRULE_responseref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(httpspecParserLESSTHANSIGN)
	}
	{
		p.SetState(226)
		p.Match(httpspecParserGREATERTHANSIGN)
	}
	{
		p.SetState(227)
		p.Match(httpspecParserWHITESPACE)
	}
	{
		p.SetState(228)
		p.Filepath()
	}

	return localctx
}
