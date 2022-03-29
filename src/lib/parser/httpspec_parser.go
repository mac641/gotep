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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 188,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 6, 3, 48, 10, 3, 13, 3, 14, 3, 49, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 57, 10, 4, 3, 4, 5, 4, 60, 10, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 3, 5,
	5, 5, 67, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 72, 10, 5, 3, 6, 3, 6, 3, 6, 5,
	6, 77, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 83, 10, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 88, 10, 7, 3, 7, 3, 7, 5, 7, 92, 10, 7, 3, 8, 3, 8, 5, 8, 96, 10,
	8, 3, 9, 3, 9, 3, 9, 5, 9, 101, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	5, 10, 108, 10, 10, 3, 11, 6, 11, 111, 10, 11, 13, 11, 14, 11, 112, 3,
	12, 6, 12, 116, 10, 12, 13, 12, 14, 12, 117, 3, 13, 3, 13, 3, 13, 7, 13,
	123, 10, 13, 12, 13, 14, 13, 126, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 140, 10, 15,
	3, 16, 3, 16, 3, 16, 5, 16, 145, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 152, 10, 17, 3, 18, 3, 18, 5, 18, 156, 10, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 164, 10, 19, 12, 19, 14, 19, 167, 11,
	19, 3, 19, 3, 19, 5, 19, 171, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 186, 10,
	21, 3, 21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 2, 4, 4, 2, 11, 11, 17, 17, 4, 2, 11, 12, 14, 15,
	2, 191, 2, 42, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 66,
	3, 2, 2, 2, 10, 76, 3, 2, 2, 2, 12, 82, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2,
	16, 97, 3, 2, 2, 2, 18, 107, 3, 2, 2, 2, 20, 110, 3, 2, 2, 2, 22, 115,
	3, 2, 2, 2, 24, 124, 3, 2, 2, 2, 26, 127, 3, 2, 2, 2, 28, 139, 3, 2, 2,
	2, 30, 141, 3, 2, 2, 2, 32, 151, 3, 2, 2, 2, 34, 153, 3, 2, 2, 2, 36, 159,
	3, 2, 2, 2, 38, 172, 3, 2, 2, 2, 40, 185, 3, 2, 2, 2, 42, 43, 5, 4, 3,
	2, 43, 44, 7, 2, 2, 3, 44, 3, 3, 2, 2, 2, 45, 46, 7, 18, 2, 2, 46, 48,
	5, 6, 4, 2, 47, 45, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2,
	49, 50, 3, 2, 2, 2, 50, 5, 3, 2, 2, 2, 51, 52, 5, 8, 5, 2, 52, 53, 7, 3,
	2, 2, 53, 54, 5, 24, 13, 2, 54, 56, 7, 3, 2, 2, 55, 57, 5, 28, 15, 2, 56,
	55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 60, 5, 40,
	21, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61,
	63, 7, 40, 2, 2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 7, 3, 2, 2,
	2, 64, 65, 7, 19, 2, 2, 65, 67, 7, 10, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 71, 5, 10, 6, 2, 69, 70, 7, 10, 2,
	2, 70, 72, 7, 27, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 9,
	3, 2, 2, 2, 73, 77, 7, 21, 2, 2, 74, 77, 5, 12, 7, 2, 75, 77, 7, 20, 2,
	2, 76, 73, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 11,
	3, 2, 2, 2, 78, 79, 7, 24, 2, 2, 79, 80, 7, 14, 2, 2, 80, 81, 7, 11, 2,
	2, 81, 83, 7, 11, 2, 2, 82, 78, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84,
	3, 2, 2, 2, 84, 87, 5, 14, 8, 2, 85, 86, 7, 15, 2, 2, 86, 88, 7, 28, 2,
	2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 90,
	7, 12, 2, 2, 90, 92, 7, 29, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2,
	2, 92, 13, 3, 2, 2, 2, 93, 95, 5, 16, 9, 2, 94, 96, 7, 23, 2, 2, 95, 94,
	3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 15, 3, 2, 2, 2, 97, 100, 5, 18, 10,
	2, 98, 99, 7, 14, 2, 2, 99, 101, 7, 25, 2, 2, 100, 98, 3, 2, 2, 2, 100,
	101, 3, 2, 2, 2, 101, 17, 3, 2, 2, 2, 102, 103, 7, 16, 2, 2, 103, 104,
	5, 20, 11, 2, 104, 105, 7, 17, 2, 2, 105, 108, 3, 2, 2, 2, 106, 108, 5,
	22, 12, 2, 107, 102, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 19, 3, 2, 2,
	2, 109, 111, 10, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 21, 3, 2, 2, 2, 114, 116, 10,
	3, 2, 2, 115, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 115, 3, 2, 2,
	2, 117, 118, 3, 2, 2, 2, 118, 23, 3, 2, 2, 2, 119, 120, 5, 26, 14, 2, 120,
	121, 7, 3, 2, 2, 121, 123, 3, 2, 2, 2, 122, 119, 3, 2, 2, 2, 123, 126,
	3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 25, 3, 2,
	2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 7, 30, 2, 2, 128, 129, 7, 14, 2,
	2, 129, 130, 8, 14, 1, 2, 130, 131, 7, 10, 2, 2, 131, 132, 8, 14, 1, 2,
	132, 133, 7, 31, 2, 2, 133, 134, 8, 14, 1, 2, 134, 135, 7, 10, 2, 2, 135,
	136, 8, 14, 1, 2, 136, 27, 3, 2, 2, 2, 137, 140, 5, 30, 16, 2, 138, 140,
	5, 34, 18, 2, 139, 137, 3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 29, 3, 2,
	2, 2, 141, 144, 5, 32, 17, 2, 142, 143, 7, 3, 2, 2, 143, 145, 5, 32, 17,
	2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 31, 3, 2, 2, 2, 146,
	147, 7, 6, 2, 2, 147, 148, 7, 6, 2, 2, 148, 149, 7, 6, 2, 2, 149, 152,
	7, 5, 2, 2, 150, 152, 7, 39, 2, 2, 151, 146, 3, 2, 2, 2, 151, 150, 3, 2,
	2, 2, 152, 33, 3, 2, 2, 2, 153, 155, 5, 36, 19, 2, 154, 156, 5, 34, 18,
	2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157,
	158, 7, 38, 2, 2, 158, 35, 3, 2, 2, 2, 159, 165, 7, 38, 2, 2, 160, 161,
	5, 26, 14, 2, 161, 162, 7, 3, 2, 2, 162, 164, 3, 2, 2, 2, 163, 160, 3,
	2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2,
	2, 166, 168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 170, 7, 3, 2, 2, 169,
	171, 5, 30, 16, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 37,
	3, 2, 2, 2, 172, 173, 7, 5, 2, 2, 173, 39, 3, 2, 2, 2, 174, 175, 7, 36,
	2, 2, 175, 176, 7, 10, 2, 2, 176, 177, 7, 32, 2, 2, 177, 178, 7, 34, 2,
	2, 178, 179, 5, 38, 20, 2, 179, 180, 7, 34, 2, 2, 180, 181, 7, 33, 2, 2,
	181, 186, 3, 2, 2, 2, 182, 183, 7, 36, 2, 2, 183, 184, 7, 10, 2, 2, 184,
	186, 7, 37, 2, 2, 185, 174, 3, 2, 2, 2, 185, 182, 3, 2, 2, 2, 186, 41,
	3, 2, 2, 2, 25, 49, 56, 59, 62, 66, 71, 76, 82, 87, 91, 95, 100, 107, 112,
	117, 124, 139, 144, 151, 155, 165, 170, 185,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "'\u002F'", "'\u0023'", "", "'\u003A'",
	"'\u003F'", "'\u005B'", "'\u005D'", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "'\u007B'", "'\u007D'", "'\u0025'", "'\u003C'", "'\u003E'",
}
var symbolicNames = []string{
	"", "NEWLINE", "NEWLINEWITHINDENT", "LINETAIL", "INPUT", "ALPHA", "DIGIT",
	"ID", "WHITESPACE", "SLASH", "HASHTAG", "LINECOMMENT", "COLON", "QUESTIONMARK",
	"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "REQUESTSEPARATOR", "METHOD",
	"ASTERISKFORM", "ORIGINFORM", "SEGMENT", "ABSOLUTEPATH", "SCHEME", "PORT",
	"PATHSEPARATOR", "HTTPVERSION", "QUERY", "REQUESTFRAGMENT", "FIELDNAME",
	"FIELDVALUE", "LEFTCURLYBRACKET", "RIGHTCURLYBRACKET", "PERCENTSIGN", "LESSTHANSIGN",
	"GREATERTHANSIGN", "FILEPATH", "BOUNDARY", "INPUTFILEREF", "RESPONSEREF",
	"ENVVARIABLE",
}

var ruleNames = []string{
	"file", "requests", "request", "requestline", "requesttarget", "absoluteform",
	"hierpart", "authority", "host", "ipv6address", "ipv4addressorregname",
	"headers", "headerfield", "messagebody", "messages", "messageline", "multipartformdata",
	"multipartfield", "handlerscript", "responsehandler",
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
	httpspecParserEOF                = antlr.TokenEOF
	httpspecParserNEWLINE            = 1
	httpspecParserNEWLINEWITHINDENT  = 2
	httpspecParserLINETAIL           = 3
	httpspecParserINPUT              = 4
	httpspecParserALPHA              = 5
	httpspecParserDIGIT              = 6
	httpspecParserID                 = 7
	httpspecParserWHITESPACE         = 8
	httpspecParserSLASH              = 9
	httpspecParserHASHTAG            = 10
	httpspecParserLINECOMMENT        = 11
	httpspecParserCOLON              = 12
	httpspecParserQUESTIONMARK       = 13
	httpspecParserLEFTSQUAREBRACKET  = 14
	httpspecParserRIGHTSQUAREBRACKET = 15
	httpspecParserREQUESTSEPARATOR   = 16
	httpspecParserMETHOD             = 17
	httpspecParserASTERISKFORM       = 18
	httpspecParserORIGINFORM         = 19
	httpspecParserSEGMENT            = 20
	httpspecParserABSOLUTEPATH       = 21
	httpspecParserSCHEME             = 22
	httpspecParserPORT               = 23
	httpspecParserPATHSEPARATOR      = 24
	httpspecParserHTTPVERSION        = 25
	httpspecParserQUERY              = 26
	httpspecParserREQUESTFRAGMENT    = 27
	httpspecParserFIELDNAME          = 28
	httpspecParserFIELDVALUE         = 29
	httpspecParserLEFTCURLYBRACKET   = 30
	httpspecParserRIGHTCURLYBRACKET  = 31
	httpspecParserPERCENTSIGN        = 32
	httpspecParserLESSTHANSIGN       = 33
	httpspecParserGREATERTHANSIGN    = 34
	httpspecParserFILEPATH           = 35
	httpspecParserBOUNDARY           = 36
	httpspecParserINPUTFILEREF       = 37
	httpspecParserRESPONSEREF        = 38
	httpspecParserENVVARIABLE        = 39
)

// httpspecParser rules.
const (
	httpspecParserRULE_file                 = 0
	httpspecParserRULE_requests             = 1
	httpspecParserRULE_request              = 2
	httpspecParserRULE_requestline          = 3
	httpspecParserRULE_requesttarget        = 4
	httpspecParserRULE_absoluteform         = 5
	httpspecParserRULE_hierpart             = 6
	httpspecParserRULE_authority            = 7
	httpspecParserRULE_host                 = 8
	httpspecParserRULE_ipv6address          = 9
	httpspecParserRULE_ipv4addressorregname = 10
	httpspecParserRULE_headers              = 11
	httpspecParserRULE_headerfield          = 12
	httpspecParserRULE_messagebody          = 13
	httpspecParserRULE_messages             = 14
	httpspecParserRULE_messageline          = 15
	httpspecParserRULE_multipartformdata    = 16
	httpspecParserRULE_multipartfield       = 17
	httpspecParserRULE_handlerscript        = 18
	httpspecParserRULE_responsehandler      = 19
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
		p.SetState(40)
		p.Requests()
	}
	{
		p.SetState(41)
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

func (s *RequestsContext) AllREQUESTSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserREQUESTSEPARATOR)
}

func (s *RequestsContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserREQUESTSEPARATOR, i)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpspecParserREQUESTSEPARATOR {
		{
			p.SetState(43)
			p.Match(httpspecParserREQUESTSEPARATOR)
		}
		{
			p.SetState(44)
			p.Request()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *RequestContext) RESPONSEREF() antlr.TerminalNode {
	return s.GetToken(httpspecParserRESPONSEREF, 0)
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
	p.EnterRule(localctx, 4, httpspecParserRULE_request)
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
		p.SetState(49)
		p.Requestline()
	}
	{
		p.SetState(50)
		p.Match(httpspecParserNEWLINE)
	}
	{
		p.SetState(51)
		p.Headers()
	}
	{
		p.SetState(52)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserINPUT || _la == httpspecParserBOUNDARY || _la == httpspecParserINPUTFILEREF {
		{
			p.SetState(53)
			p.Messagebody()
		}

	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserGREATERTHANSIGN {
		{
			p.SetState(56)
			p.Responsehandler()
		}

	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserRESPONSEREF {
		{
			p.SetState(59)
			p.Match(httpspecParserRESPONSEREF)
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
	p.EnterRule(localctx, 6, httpspecParserRULE_requestline)
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(62)
			p.Match(httpspecParserMETHOD)
		}
		{
			p.SetState(63)
			p.Match(httpspecParserWHITESPACE)
		}

	}
	{
		p.SetState(66)
		p.Requesttarget()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserWHITESPACE {
		{
			p.SetState(67)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(68)
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
	p.EnterRule(localctx, 8, httpspecParserRULE_requesttarget)

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(httpspecParserORIGINFORM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Absoluteform()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(httpspecParserASTERISKFORM)
		}

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
	p.EnterRule(localctx, 10, httpspecParserRULE_absoluteform)
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

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(76)
			p.Match(httpspecParserSCHEME)
		}
		{
			p.SetState(77)
			p.Match(httpspecParserCOLON)
		}
		{
			p.SetState(78)
			p.Match(httpspecParserSLASH)
		}
		{
			p.SetState(79)
			p.Match(httpspecParserSLASH)
		}

	}
	{
		p.SetState(82)
		p.Hierpart()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserQUESTIONMARK {
		{
			p.SetState(83)
			p.Match(httpspecParserQUESTIONMARK)
		}
		{
			p.SetState(84)
			p.Match(httpspecParserQUERY)
		}

	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserHASHTAG {
		{
			p.SetState(87)
			p.Match(httpspecParserHASHTAG)
		}
		{
			p.SetState(88)
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
	p.EnterRule(localctx, 12, httpspecParserRULE_hierpart)
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
		p.SetState(91)
		p.Authority()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserABSOLUTEPATH {
		{
			p.SetState(92)
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
	p.EnterRule(localctx, 14, httpspecParserRULE_authority)
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
		p.SetState(95)
		p.Host()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserCOLON {
		{
			p.SetState(96)
			p.Match(httpspecParserCOLON)
		}
		{
			p.SetState(97)
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

func (s *HostContext) Ipv6address() IIpv6addressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpv6addressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpv6addressContext)
}

func (s *HostContext) RIGHTSQUAREBRACKET() antlr.TerminalNode {
	return s.GetToken(httpspecParserRIGHTSQUAREBRACKET, 0)
}

func (s *HostContext) Ipv4addressorregname() IIpv4addressorregnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpv4addressorregnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpv4addressorregnameContext)
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
	p.EnterRule(localctx, 16, httpspecParserRULE_host)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(httpspecParserLEFTSQUAREBRACKET)
		}
		{
			p.SetState(101)
			p.Ipv6address()
		}
		{
			p.SetState(102)
			p.Match(httpspecParserRIGHTSQUAREBRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Ipv4addressorregname()
		}

	}

	return localctx
}

// IIpv6addressContext is an interface to support dynamic dispatch.
type IIpv6addressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpv6addressContext differentiates from other interfaces.
	IsIpv6addressContext()
}

type Ipv6addressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv6addressContext() *Ipv6addressContext {
	var p = new(Ipv6addressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_ipv6address
	return p
}

func (*Ipv6addressContext) IsIpv6addressContext() {}

func NewIpv6addressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv6addressContext {
	var p = new(Ipv6addressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_ipv6address

	return p
}

func (s *Ipv6addressContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv6addressContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserSLASH)
}

func (s *Ipv6addressContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserSLASH, i)
}

func (s *Ipv6addressContext) AllRIGHTSQUAREBRACKET() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserRIGHTSQUAREBRACKET)
}

func (s *Ipv6addressContext) RIGHTSQUAREBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserRIGHTSQUAREBRACKET, i)
}

func (s *Ipv6addressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv6addressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv6addressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterIpv6address(s)
	}
}

func (s *Ipv6addressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitIpv6address(s)
	}
}

func (p *httpspecParser) Ipv6address() (localctx IIpv6addressContext) {
	localctx = NewIpv6addressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, httpspecParserRULE_ipv6address)
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
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpspecParserNEWLINE)|(1<<httpspecParserNEWLINEWITHINDENT)|(1<<httpspecParserLINETAIL)|(1<<httpspecParserINPUT)|(1<<httpspecParserALPHA)|(1<<httpspecParserDIGIT)|(1<<httpspecParserID)|(1<<httpspecParserWHITESPACE)|(1<<httpspecParserHASHTAG)|(1<<httpspecParserLINECOMMENT)|(1<<httpspecParserCOLON)|(1<<httpspecParserQUESTIONMARK)|(1<<httpspecParserLEFTSQUAREBRACKET)|(1<<httpspecParserREQUESTSEPARATOR)|(1<<httpspecParserMETHOD)|(1<<httpspecParserASTERISKFORM)|(1<<httpspecParserORIGINFORM)|(1<<httpspecParserSEGMENT)|(1<<httpspecParserABSOLUTEPATH)|(1<<httpspecParserSCHEME)|(1<<httpspecParserPORT)|(1<<httpspecParserPATHSEPARATOR)|(1<<httpspecParserHTTPVERSION)|(1<<httpspecParserQUERY)|(1<<httpspecParserREQUESTFRAGMENT)|(1<<httpspecParserFIELDNAME)|(1<<httpspecParserFIELDVALUE)|(1<<httpspecParserLEFTCURLYBRACKET)|(1<<httpspecParserRIGHTCURLYBRACKET))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(httpspecParserPERCENTSIGN-32))|(1<<(httpspecParserLESSTHANSIGN-32))|(1<<(httpspecParserGREATERTHANSIGN-32))|(1<<(httpspecParserFILEPATH-32))|(1<<(httpspecParserBOUNDARY-32))|(1<<(httpspecParserINPUTFILEREF-32))|(1<<(httpspecParserRESPONSEREF-32))|(1<<(httpspecParserENVVARIABLE-32)))) != 0) {
		{
			p.SetState(107)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == httpspecParserSLASH || _la == httpspecParserRIGHTSQUAREBRACKET {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpv4addressorregnameContext is an interface to support dynamic dispatch.
type IIpv4addressorregnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpv4addressorregnameContext differentiates from other interfaces.
	IsIpv4addressorregnameContext()
}

type Ipv4addressorregnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv4addressorregnameContext() *Ipv4addressorregnameContext {
	var p = new(Ipv4addressorregnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_ipv4addressorregname
	return p
}

func (*Ipv4addressorregnameContext) IsIpv4addressorregnameContext() {}

func NewIpv4addressorregnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv4addressorregnameContext {
	var p = new(Ipv4addressorregnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_ipv4addressorregname

	return p
}

func (s *Ipv4addressorregnameContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv4addressorregnameContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserSLASH)
}

func (s *Ipv4addressorregnameContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserSLASH, i)
}

func (s *Ipv4addressorregnameContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserCOLON)
}

func (s *Ipv4addressorregnameContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserCOLON, i)
}

func (s *Ipv4addressorregnameContext) AllQUESTIONMARK() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserQUESTIONMARK)
}

func (s *Ipv4addressorregnameContext) QUESTIONMARK(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserQUESTIONMARK, i)
}

func (s *Ipv4addressorregnameContext) AllHASHTAG() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserHASHTAG)
}

func (s *Ipv4addressorregnameContext) HASHTAG(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserHASHTAG, i)
}

func (s *Ipv4addressorregnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv4addressorregnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv4addressorregnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterIpv4addressorregname(s)
	}
}

func (s *Ipv4addressorregnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitIpv4addressorregname(s)
	}
}

func (p *httpspecParser) Ipv4addressorregname() (localctx IIpv4addressorregnameContext) {
	localctx = NewIpv4addressorregnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, httpspecParserRULE_ipv4addressorregname)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(112)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpspecParserSLASH)|(1<<httpspecParserHASHTAG)|(1<<httpspecParserCOLON)|(1<<httpspecParserQUESTIONMARK))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserFIELDNAME {
		{
			p.SetState(117)
			p.Headerfield()
		}
		{
			p.SetState(118)
			p.Match(httpspecParserNEWLINE)
		}

		p.SetState(124)
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

func (s *HeaderfieldContext) FIELDNAME() antlr.TerminalNode {
	return s.GetToken(httpspecParserFIELDNAME, 0)
}

func (s *HeaderfieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(httpspecParserCOLON, 0)
}

func (s *HeaderfieldContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserWHITESPACE)
}

func (s *HeaderfieldContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserWHITESPACE, i)
}

func (s *HeaderfieldContext) FIELDVALUE() antlr.TerminalNode {
	return s.GetToken(httpspecParserFIELDVALUE, 0)
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
		p.SetState(125)
		p.Match(httpspecParserFIELDNAME)
	}
	{
		p.SetState(126)
		p.Match(httpspecParserCOLON)
	}
	ignoreWs = true
	{
		p.SetState(128)
		p.Match(httpspecParserWHITESPACE)
	}
	ignoreWs = false
	{
		p.SetState(130)
		p.Match(httpspecParserFIELDVALUE)
	}
	ignoreWs = true
	{
		p.SetState(132)
		p.Match(httpspecParserWHITESPACE)
	}
	ignoreWs = false

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
	p.EnterRule(localctx, 26, httpspecParserRULE_messagebody)

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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpspecParserINPUT, httpspecParserINPUTFILEREF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Messages()
		}

	case httpspecParserBOUNDARY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Multipartformdata()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 28, httpspecParserRULE_messages)
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
		p.SetState(139)
		p.Messageline()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserNEWLINE {
		{
			p.SetState(140)
			p.Match(httpspecParserNEWLINE)
		}
		{
			p.SetState(141)
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

func (s *MessagelineContext) AllINPUT() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserINPUT)
}

func (s *MessagelineContext) INPUT(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, i)
}

func (s *MessagelineContext) LINETAIL() antlr.TerminalNode {
	return s.GetToken(httpspecParserLINETAIL, 0)
}

func (s *MessagelineContext) INPUTFILEREF() antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUTFILEREF, 0)
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
	p.EnterRule(localctx, 30, httpspecParserRULE_messageline)

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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpspecParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Match(httpspecParserINPUT)
		}
		{
			p.SetState(145)
			p.Match(httpspecParserINPUT)
		}
		{
			p.SetState(146)
			p.Match(httpspecParserINPUT)
		}
		{
			p.SetState(147)
			p.Match(httpspecParserLINETAIL)
		}

	case httpspecParserINPUTFILEREF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(httpspecParserINPUTFILEREF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 32, httpspecParserRULE_multipartformdata)

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
		p.SetState(151)
		p.Multipartfield()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(152)
			p.Multipartformdata()
		}

	}
	{
		p.SetState(155)
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
	p.EnterRule(localctx, 34, httpspecParserRULE_multipartfield)
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
		p.SetState(157)
		p.Match(httpspecParserBOUNDARY)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserFIELDNAME {
		{
			p.SetState(158)
			p.Headerfield()
		}
		{
			p.SetState(159)
			p.Match(httpspecParserNEWLINE)
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserINPUT || _la == httpspecParserINPUTFILEREF {
		{
			p.SetState(167)
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

func (s *HandlerscriptContext) LINETAIL() antlr.TerminalNode {
	return s.GetToken(httpspecParserLINETAIL, 0)
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
	p.EnterRule(localctx, 36, httpspecParserRULE_handlerscript)

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
		p.SetState(170)
		p.Match(httpspecParserLINETAIL)
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

func (s *ResponsehandlerContext) FILEPATH() antlr.TerminalNode {
	return s.GetToken(httpspecParserFILEPATH, 0)
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
	p.EnterRule(localctx, 38, httpspecParserRULE_responsehandler)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(httpspecParserGREATERTHANSIGN)
		}
		{
			p.SetState(173)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(174)
			p.Match(httpspecParserLEFTCURLYBRACKET)
		}
		{
			p.SetState(175)
			p.Match(httpspecParserPERCENTSIGN)
		}
		{
			p.SetState(176)
			p.Handlerscript()
		}
		{
			p.SetState(177)
			p.Match(httpspecParserPERCENTSIGN)
		}
		{
			p.SetState(178)
			p.Match(httpspecParserRIGHTCURLYBRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Match(httpspecParserGREATERTHANSIGN)
		}
		{
			p.SetState(181)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(182)
			p.Match(httpspecParserFILEPATH)
		}

	}

	return localctx
}
