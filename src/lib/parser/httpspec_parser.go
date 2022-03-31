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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 171,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 7,
	2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 2, 3, 2, 7, 2, 43, 10, 2, 12,
	2, 14, 2, 46, 11, 2, 3, 2, 7, 2, 49, 10, 2, 12, 2, 14, 2, 52, 11, 2, 3,
	2, 3, 2, 3, 3, 6, 3, 57, 10, 3, 13, 3, 14, 3, 58, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 68, 10, 4, 3, 4, 5, 4, 71, 10, 4, 3, 4, 5, 4,
	74, 10, 4, 3, 5, 3, 5, 5, 5, 78, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 83, 10,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 5, 7, 89, 10, 7, 7, 7, 91, 10, 7, 12, 7, 14,
	7, 94, 11, 7, 3, 8, 7, 8, 97, 10, 8, 12, 8, 14, 8, 100, 11, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 110, 10, 10, 3, 11, 3,
	11, 3, 11, 5, 11, 115, 10, 11, 3, 12, 3, 12, 7, 12, 119, 10, 12, 12, 12,
	14, 12, 122, 11, 12, 3, 12, 3, 12, 5, 12, 126, 10, 12, 3, 13, 3, 13, 5,
	13, 130, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 137, 10, 14,
	12, 14, 14, 14, 140, 11, 14, 3, 14, 3, 14, 5, 14, 144, 10, 14, 3, 15, 6,
	15, 147, 10, 15, 13, 15, 14, 15, 148, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 164, 10,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 4, 3, 2, 15, 17, 3, 2, 3,
	5, 2, 175, 2, 37, 3, 2, 2, 2, 4, 56, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8,
	77, 3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14, 98, 3, 2, 2,
	2, 16, 103, 3, 2, 2, 2, 18, 109, 3, 2, 2, 2, 20, 111, 3, 2, 2, 2, 22, 125,
	3, 2, 2, 2, 24, 127, 3, 2, 2, 2, 26, 133, 3, 2, 2, 2, 28, 146, 3, 2, 2,
	2, 30, 163, 3, 2, 2, 2, 32, 165, 3, 2, 2, 2, 34, 36, 7, 14, 2, 2, 35, 34,
	3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2,
	38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 44, 5, 6, 4, 2, 41, 43, 5,
	4, 3, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44,
	45, 3, 2, 2, 2, 45, 50, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 49, 7, 14,
	2, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 2, 2, 3,
	54, 3, 3, 2, 2, 2, 55, 57, 7, 14, 2, 2, 56, 55, 3, 2, 2, 2, 57, 58, 3,
	2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60,
	61, 5, 6, 4, 2, 61, 5, 3, 2, 2, 2, 62, 63, 5, 8, 5, 2, 63, 64, 7, 6, 2,
	2, 64, 65, 5, 12, 7, 2, 65, 67, 7, 6, 2, 2, 66, 68, 5, 18, 10, 2, 67, 66,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 71, 5, 30, 16,
	2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 74,
	5, 32, 17, 2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 7, 3, 2, 2, 2,
	75, 76, 7, 12, 2, 2, 76, 78, 7, 29, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 82, 5, 10, 6, 2, 80, 81, 7, 29, 2, 2,
	81, 83, 7, 13, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 9, 3,
	2, 2, 2, 84, 85, 9, 2, 2, 2, 85, 11, 3, 2, 2, 2, 86, 88, 7, 21, 2, 2, 87,
	89, 7, 6, 2, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 91, 3, 2, 2,
	2, 90, 86, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 13, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 97, 7, 28, 2, 2,
	96, 95, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3,
	2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 102, 7, 6, 2, 2,
	102, 15, 3, 2, 2, 2, 103, 104, 7, 25, 2, 2, 104, 105, 7, 29, 2, 2, 105,
	106, 5, 14, 8, 2, 106, 17, 3, 2, 2, 2, 107, 110, 5, 20, 11, 2, 108, 110,
	5, 24, 13, 2, 109, 107, 3, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 19, 3, 2,
	2, 2, 111, 114, 5, 22, 12, 2, 112, 113, 7, 6, 2, 2, 113, 115, 5, 22, 12,
	2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 21, 3, 2, 2, 2, 116,
	120, 10, 3, 2, 2, 117, 119, 7, 28, 2, 2, 118, 117, 3, 2, 2, 2, 119, 122,
	3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2,
	2, 2, 122, 120, 3, 2, 2, 2, 123, 126, 7, 6, 2, 2, 124, 126, 5, 16, 9, 2,
	125, 116, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 23, 3, 2, 2, 2, 127, 129,
	5, 26, 14, 2, 128, 130, 5, 24, 13, 2, 129, 128, 3, 2, 2, 2, 129, 130, 3,
	2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 27, 2, 2, 132, 25, 3, 2, 2,
	2, 133, 138, 7, 27, 2, 2, 134, 135, 7, 21, 2, 2, 135, 137, 7, 6, 2, 2,
	136, 134, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 141, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 143,
	7, 6, 2, 2, 142, 144, 5, 20, 11, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3,
	2, 2, 2, 144, 27, 3, 2, 2, 2, 145, 147, 7, 28, 2, 2, 146, 145, 3, 2, 2,
	2, 147, 148, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	150, 3, 2, 2, 2, 150, 151, 7, 6, 2, 2, 151, 29, 3, 2, 2, 2, 152, 153, 7,
	26, 2, 2, 153, 154, 7, 29, 2, 2, 154, 155, 7, 22, 2, 2, 155, 156, 7, 24,
	2, 2, 156, 157, 5, 28, 15, 2, 157, 158, 7, 24, 2, 2, 158, 159, 7, 23, 2,
	2, 159, 164, 3, 2, 2, 2, 160, 161, 7, 26, 2, 2, 161, 162, 7, 29, 2, 2,
	162, 164, 5, 14, 8, 2, 163, 152, 3, 2, 2, 2, 163, 160, 3, 2, 2, 2, 164,
	31, 3, 2, 2, 2, 165, 166, 7, 25, 2, 2, 166, 167, 7, 26, 2, 2, 167, 168,
	7, 29, 2, 2, 168, 169, 5, 14, 8, 2, 169, 33, 3, 2, 2, 2, 23, 37, 44, 50,
	58, 67, 70, 73, 77, 82, 88, 92, 98, 109, 114, 120, 125, 129, 138, 143,
	148, 163,
}
var literalNames = []string{
	"", "'<'", "'<>'", "'###'", "", "", "'\u003A'", "'\u003F'", "'\u005B'",
	"'\u005D'", "", "", "", "", "", "", "'\u002F'", "'\u0023'", "", "", "'\u007B'",
	"'\u007D'", "'\u0025'", "'\u003C'", "'\u003E'",
}
var symbolicNames = []string{
	"", "", "", "", "NEWLINE", "NEWLINEWITHINDENT", "COLON", "QUESTIONMARK",
	"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "METHOD", "HTTPVERSION", "REQUESTSEPARATOR",
	"ASTERISKFORM", "ABSOLUTEFORM", "ORIGINFORM", "SLASH", "HASHTAG", "LINECOMMENT",
	"HEADERFIELD", "LEFTCURLYBRACKET", "RIGHTCURLYBRACKET", "PERCENTSIGN",
	"LESSTHANSIGN", "GREATERTHANSIGN", "BOUNDARY", "INPUT", "WHITESPACE", "ENVVARIABLE",
}

var ruleNames = []string{
	"file", "requestwithseparator", "request", "requestline", "requesttarget",
	"headers", "filepath", "inputfileref", "messagebody", "messages", "messageline",
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
	httpspecParserEOF                = antlr.TokenEOF
	httpspecParserT__0               = 1
	httpspecParserT__1               = 2
	httpspecParserT__2               = 3
	httpspecParserNEWLINE            = 4
	httpspecParserNEWLINEWITHINDENT  = 5
	httpspecParserCOLON              = 6
	httpspecParserQUESTIONMARK       = 7
	httpspecParserLEFTSQUAREBRACKET  = 8
	httpspecParserRIGHTSQUAREBRACKET = 9
	httpspecParserMETHOD             = 10
	httpspecParserHTTPVERSION        = 11
	httpspecParserREQUESTSEPARATOR   = 12
	httpspecParserASTERISKFORM       = 13
	httpspecParserABSOLUTEFORM       = 14
	httpspecParserORIGINFORM         = 15
	httpspecParserSLASH              = 16
	httpspecParserHASHTAG            = 17
	httpspecParserLINECOMMENT        = 18
	httpspecParserHEADERFIELD        = 19
	httpspecParserLEFTCURLYBRACKET   = 20
	httpspecParserRIGHTCURLYBRACKET  = 21
	httpspecParserPERCENTSIGN        = 22
	httpspecParserLESSTHANSIGN       = 23
	httpspecParserGREATERTHANSIGN    = 24
	httpspecParserBOUNDARY           = 25
	httpspecParserINPUT              = 26
	httpspecParserWHITESPACE         = 27
	httpspecParserENVVARIABLE        = 28
)

// httpspecParser rules.
const (
	httpspecParserRULE_file                 = 0
	httpspecParserRULE_requestwithseparator = 1
	httpspecParserRULE_request              = 2
	httpspecParserRULE_requestline          = 3
	httpspecParserRULE_requesttarget        = 4
	httpspecParserRULE_headers              = 5
	httpspecParserRULE_filepath             = 6
	httpspecParserRULE_inputfileref         = 7
	httpspecParserRULE_messagebody          = 8
	httpspecParserRULE_messages             = 9
	httpspecParserRULE_messageline          = 10
	httpspecParserRULE_multipartformdata    = 11
	httpspecParserRULE_multipartfield       = 12
	httpspecParserRULE_handlerscript        = 13
	httpspecParserRULE_responsehandler      = 14
	httpspecParserRULE_responseref          = 15
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

func (s *FileContext) Request() IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(httpspecParserEOF, 0)
}

func (s *FileContext) AllREQUESTSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserREQUESTSEPARATOR)
}

func (s *FileContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserREQUESTSEPARATOR, i)
}

func (s *FileContext) AllRequestwithseparator() []IRequestwithseparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRequestwithseparatorContext)(nil)).Elem())
	var tst = make([]IRequestwithseparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRequestwithseparatorContext)
		}
	}

	return tst
}

func (s *FileContext) Requestwithseparator(i int) IRequestwithseparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestwithseparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRequestwithseparatorContext)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserREQUESTSEPARATOR {
		{
			p.SetState(32)
			p.Match(httpspecParserREQUESTSEPARATOR)
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Request()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(39)
				p.Requestwithseparator()
			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserREQUESTSEPARATOR {
		{
			p.SetState(45)
			p.Match(httpspecParserREQUESTSEPARATOR)
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(httpspecParserEOF)
	}

	return localctx
}

// IRequestwithseparatorContext is an interface to support dynamic dispatch.
type IRequestwithseparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestwithseparatorContext differentiates from other interfaces.
	IsRequestwithseparatorContext()
}

type RequestwithseparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestwithseparatorContext() *RequestwithseparatorContext {
	var p = new(RequestwithseparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpspecParserRULE_requestwithseparator
	return p
}

func (*RequestwithseparatorContext) IsRequestwithseparatorContext() {}

func NewRequestwithseparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestwithseparatorContext {
	var p = new(RequestwithseparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpspecParserRULE_requestwithseparator

	return p
}

func (s *RequestwithseparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestwithseparatorContext) Request() IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RequestwithseparatorContext) AllREQUESTSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserREQUESTSEPARATOR)
}

func (s *RequestwithseparatorContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserREQUESTSEPARATOR, i)
}

func (s *RequestwithseparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestwithseparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestwithseparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.EnterRequestwithseparator(s)
	}
}

func (s *RequestwithseparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpspecListener); ok {
		listenerT.ExitRequestwithseparator(s)
	}
}

func (p *httpspecParser) Requestwithseparator() (localctx IRequestwithseparatorContext) {
	localctx = NewRequestwithseparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, httpspecParserRULE_requestwithseparator)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpspecParserREQUESTSEPARATOR {
		{
			p.SetState(53)
			p.Match(httpspecParserREQUESTSEPARATOR)
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Request()
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
		p.SetState(60)
		p.Requestline()
	}
	{
		p.SetState(61)
		p.Match(httpspecParserNEWLINE)
	}
	{
		p.SetState(62)
		p.Headers()
	}
	{
		p.SetState(63)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(64)
			p.Messagebody()
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserGREATERTHANSIGN {
		{
			p.SetState(67)
			p.Responsehandler()
		}

	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserLESSTHANSIGN {
		{
			p.SetState(70)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserMETHOD {
		{
			p.SetState(73)
			p.Match(httpspecParserMETHOD)
		}
		{
			p.SetState(74)
			p.Match(httpspecParserWHITESPACE)
		}

	}
	{
		p.SetState(77)
		p.Requesttarget()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserWHITESPACE {
		{
			p.SetState(78)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(79)
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

func (s *RequesttargetContext) ABSOLUTEFORM() antlr.TerminalNode {
	return s.GetToken(httpspecParserABSOLUTEFORM, 0)
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
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpspecParserASTERISKFORM)|(1<<httpspecParserABSOLUTEFORM)|(1<<httpspecParserORIGINFORM))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

func (s *HeadersContext) AllHEADERFIELD() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserHEADERFIELD)
}

func (s *HeadersContext) HEADERFIELD(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserHEADERFIELD, i)
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
	p.EnterRule(localctx, 10, httpspecParserRULE_headers)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserHEADERFIELD {
		{
			p.SetState(84)
			p.Match(httpspecParserHEADERFIELD)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(85)
				p.Match(httpspecParserNEWLINE)
			}

		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

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

func (s *FilepathContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *FilepathContext) AllINPUT() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserINPUT)
}

func (s *FilepathContext) INPUT(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, i)
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
	p.EnterRule(localctx, 12, httpspecParserRULE_filepath)
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

	for _la == httpspecParserINPUT {
		{
			p.SetState(93)
			p.Match(httpspecParserINPUT)
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
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
	p.EnterRule(localctx, 14, httpspecParserRULE_inputfileref)

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
		p.SetState(101)
		p.Match(httpspecParserLESSTHANSIGN)
	}
	{
		p.SetState(102)
		p.Match(httpspecParserWHITESPACE)
	}
	{
		p.SetState(103)
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
	p.EnterRule(localctx, 16, httpspecParserRULE_messagebody)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Messages()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
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
	p.EnterRule(localctx, 18, httpspecParserRULE_messages)
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
		p.SetState(109)
		p.Messageline()
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpspecParserNEWLINE {
		{
			p.SetState(110)
			p.Match(httpspecParserNEWLINE)
		}
		{
			p.SetState(111)
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

func (s *MessagelineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *MessagelineContext) AllINPUT() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserINPUT)
}

func (s *MessagelineContext) INPUT(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, i)
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
	p.EnterRule(localctx, 20, httpspecParserRULE_messageline)
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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpspecParserT__0)|(1<<httpspecParserT__1)|(1<<httpspecParserT__2))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == httpspecParserINPUT {
			{
				p.SetState(115)
				p.Match(httpspecParserINPUT)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(121)
			p.Match(httpspecParserNEWLINE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
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
	p.EnterRule(localctx, 22, httpspecParserRULE_multipartformdata)

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
		p.Multipartfield()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(126)
			p.Multipartformdata()
		}

	}
	{
		p.SetState(129)
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

func (s *MultipartfieldContext) AllHEADERFIELD() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserHEADERFIELD)
}

func (s *MultipartfieldContext) HEADERFIELD(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserHEADERFIELD, i)
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
	p.EnterRule(localctx, 24, httpspecParserRULE_multipartfield)
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
		p.SetState(131)
		p.Match(httpspecParserBOUNDARY)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpspecParserHEADERFIELD {
		{
			p.SetState(132)
			p.Match(httpspecParserHEADERFIELD)
		}
		{
			p.SetState(133)
			p.Match(httpspecParserNEWLINE)
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(139)
		p.Match(httpspecParserNEWLINE)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(140)
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

func (s *HandlerscriptContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpspecParserNEWLINE, 0)
}

func (s *HandlerscriptContext) AllINPUT() []antlr.TerminalNode {
	return s.GetTokens(httpspecParserINPUT)
}

func (s *HandlerscriptContext) INPUT(i int) antlr.TerminalNode {
	return s.GetToken(httpspecParserINPUT, i)
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
	p.EnterRule(localctx, 26, httpspecParserRULE_handlerscript)
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
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpspecParserINPUT {
		{
			p.SetState(143)
			p.Match(httpspecParserINPUT)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
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
	p.EnterRule(localctx, 28, httpspecParserRULE_responsehandler)

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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(httpspecParserGREATERTHANSIGN)
		}
		{
			p.SetState(151)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(152)
			p.Match(httpspecParserLEFTCURLYBRACKET)
		}
		{
			p.SetState(153)
			p.Match(httpspecParserPERCENTSIGN)
		}
		{
			p.SetState(154)
			p.Handlerscript()
		}
		{
			p.SetState(155)
			p.Match(httpspecParserPERCENTSIGN)
		}
		{
			p.SetState(156)
			p.Match(httpspecParserRIGHTCURLYBRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(httpspecParserGREATERTHANSIGN)
		}
		{
			p.SetState(159)
			p.Match(httpspecParserWHITESPACE)
		}
		{
			p.SetState(160)
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
	p.EnterRule(localctx, 30, httpspecParserRULE_responseref)

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
		p.SetState(163)
		p.Match(httpspecParserLESSTHANSIGN)
	}
	{
		p.SetState(164)
		p.Match(httpspecParserGREATERTHANSIGN)
	}
	{
		p.SetState(165)
		p.Match(httpspecParserWHITESPACE)
	}
	{
		p.SetState(166)
		p.Filepath()
	}

	return localctx
}
