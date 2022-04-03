// Code generated from ./src/lib/parser/httpSpec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpSpec

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 77, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 7, 2, 23, 10,
	2, 12, 2, 14, 2, 26, 11, 2, 3, 2, 7, 2, 29, 10, 2, 12, 2, 14, 2, 32, 11,
	2, 3, 2, 3, 2, 3, 3, 6, 3, 37, 10, 3, 13, 3, 14, 3, 38, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 47, 10, 4, 3, 5, 3, 5, 5, 5, 51, 10, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 56, 10, 5, 3, 6, 6, 6, 59, 10, 6, 13, 6, 14, 6, 60, 3,
	6, 3, 6, 5, 6, 65, 10, 6, 3, 6, 3, 6, 5, 6, 69, 10, 6, 3, 7, 7, 7, 72,
	10, 7, 12, 7, 14, 7, 75, 11, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2,
	2, 83, 2, 17, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 50,
	3, 2, 2, 2, 10, 68, 3, 2, 2, 2, 12, 73, 3, 2, 2, 2, 14, 16, 7, 10, 2, 2,
	15, 14, 3, 2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3,
	2, 2, 2, 18, 20, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 24, 5, 6, 4, 2, 21,
	23, 5, 4, 3, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3, 2, 2,
	2, 24, 25, 3, 2, 2, 2, 25, 30, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27, 29,
	7, 10, 2, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2,
	30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 34, 7,
	2, 2, 3, 34, 3, 3, 2, 2, 2, 35, 37, 7, 10, 2, 2, 36, 35, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2, 2,
	2, 40, 41, 5, 6, 4, 2, 41, 5, 3, 2, 2, 2, 42, 43, 5, 8, 5, 2, 43, 44, 7,
	3, 2, 2, 44, 46, 5, 12, 7, 2, 45, 47, 5, 10, 6, 2, 46, 45, 3, 2, 2, 2,
	46, 47, 3, 2, 2, 2, 47, 7, 3, 2, 2, 2, 48, 49, 7, 11, 2, 2, 49, 51, 7,
	6, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52,
	55, 7, 12, 2, 2, 53, 54, 7, 6, 2, 2, 54, 56, 7, 9, 2, 2, 55, 53, 3, 2,
	2, 2, 55, 56, 3, 2, 2, 2, 56, 9, 3, 2, 2, 2, 57, 59, 7, 16, 2, 2, 58, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2,
	61, 69, 3, 2, 2, 2, 62, 64, 7, 17, 2, 2, 63, 65, 7, 3, 2, 2, 64, 63, 3,
	2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 69, 3, 2, 2, 2, 66, 69, 7, 18, 2, 2, 67,
	69, 7, 19, 2, 2, 68, 58, 3, 2, 2, 2, 68, 62, 3, 2, 2, 2, 68, 66, 3, 2,
	2, 2, 68, 67, 3, 2, 2, 2, 69, 11, 3, 2, 2, 2, 70, 72, 7, 7, 2, 2, 71, 70,
	3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 13, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 13, 17, 24, 30, 38, 46, 50, 55,
	60, 64, 68, 73,
}
var literalNames []string

var symbolicNames = []string{
	"", "NEWLINE", "NEWLINEWITHINDENT", "LINECOMMENT", "WHITESPACE", "HEADERFIELD",
	"ENVVARIABLE", "HTTPVERSION", "REQUESTSEPARATOR", "METHOD", "REQUESTTARGET",
	"CLOSE_REQUESTSEPARATOR", "CLOSE_COMMENT", "CLOSE_HEADERFIELD", "MESSAGES",
	"INPUTFILEREF", "RESPONSEREF", "RESPONSEHANDLER",
}

var ruleNames = []string{
	"file", "requestwithseparator", "request", "requestline", "messagebody",
	"headers",
}

type httpSpecParser struct {
	*antlr.BaseParser
}

// NewhttpSpecParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *httpSpecParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewhttpSpecParser(input antlr.TokenStream) *httpSpecParser {
	this := new(httpSpecParser)
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
	this.GrammarFileName = "httpSpec.g4"

	return this
}

// httpSpecParser tokens.
const (
	httpSpecParserEOF                    = antlr.TokenEOF
	httpSpecParserNEWLINE                = 1
	httpSpecParserNEWLINEWITHINDENT      = 2
	httpSpecParserLINECOMMENT            = 3
	httpSpecParserWHITESPACE             = 4
	httpSpecParserHEADERFIELD            = 5
	httpSpecParserENVVARIABLE            = 6
	httpSpecParserHTTPVERSION            = 7
	httpSpecParserREQUESTSEPARATOR       = 8
	httpSpecParserMETHOD                 = 9
	httpSpecParserREQUESTTARGET          = 10
	httpSpecParserCLOSE_REQUESTSEPARATOR = 11
	httpSpecParserCLOSE_COMMENT          = 12
	httpSpecParserCLOSE_HEADERFIELD      = 13
	httpSpecParserMESSAGES               = 14
	httpSpecParserINPUTFILEREF           = 15
	httpSpecParserRESPONSEREF            = 16
	httpSpecParserRESPONSEHANDLER        = 17
)

// httpSpecParser rules.
const (
	httpSpecParserRULE_file                 = 0
	httpSpecParserRULE_requestwithseparator = 1
	httpSpecParserRULE_request              = 2
	httpSpecParserRULE_requestline          = 3
	httpSpecParserRULE_messagebody          = 4
	httpSpecParserRULE_headers              = 5
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
	p.RuleIndex = httpSpecParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_file

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
	return s.GetToken(httpSpecParserEOF, 0)
}

func (s *FileContext) AllREQUESTSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserREQUESTSEPARATOR)
}

func (s *FileContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUESTSEPARATOR, i)
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
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *httpSpecParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, httpSpecParserRULE_file)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(12)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Request()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(19)
				p.Requestwithseparator()
			}

		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(25)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(httpSpecParserEOF)
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
	p.RuleIndex = httpSpecParserRULE_requestwithseparator
	return p
}

func (*RequestwithseparatorContext) IsRequestwithseparatorContext() {}

func NewRequestwithseparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestwithseparatorContext {
	var p = new(RequestwithseparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requestwithseparator

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
	return s.GetTokens(httpSpecParserREQUESTSEPARATOR)
}

func (s *RequestwithseparatorContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUESTSEPARATOR, i)
}

func (s *RequestwithseparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestwithseparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestwithseparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequestwithseparator(s)
	}
}

func (s *RequestwithseparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequestwithseparator(s)
	}
}

func (p *httpSpecParser) Requestwithseparator() (localctx IRequestwithseparatorContext) {
	localctx = NewRequestwithseparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, httpSpecParserRULE_requestwithseparator)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(33)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
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
	p.RuleIndex = httpSpecParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_request

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

func (s *RequestContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, 0)
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

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (p *httpSpecParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, httpSpecParserRULE_request)
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
		p.SetState(40)
		p.Requestline()
	}
	{
		p.SetState(41)
		p.Match(httpSpecParserNEWLINE)
	}
	{
		p.SetState(42)
		p.Headers()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserMESSAGES)|(1<<httpSpecParserINPUTFILEREF)|(1<<httpSpecParserRESPONSEREF)|(1<<httpSpecParserRESPONSEHANDLER))) != 0 {
		{
			p.SetState(43)
			p.Messagebody()
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
	p.RuleIndex = httpSpecParserRULE_requestline
	return p
}

func (*RequestlineContext) IsRequestlineContext() {}

func NewRequestlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestlineContext {
	var p = new(RequestlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requestline

	return p
}

func (s *RequestlineContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestlineContext) REQUESTTARGET() antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUESTTARGET, 0)
}

func (s *RequestlineContext) METHOD() antlr.TerminalNode {
	return s.GetToken(httpSpecParserMETHOD, 0)
}

func (s *RequestlineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserWHITESPACE)
}

func (s *RequestlineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserWHITESPACE, i)
}

func (s *RequestlineContext) HTTPVERSION() antlr.TerminalNode {
	return s.GetToken(httpSpecParserHTTPVERSION, 0)
}

func (s *RequestlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequestline(s)
	}
}

func (s *RequestlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequestline(s)
	}
}

func (p *httpSpecParser) Requestline() (localctx IRequestlineContext) {
	localctx = NewRequestlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, httpSpecParserRULE_requestline)
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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserMETHOD {
		{
			p.SetState(46)
			p.Match(httpSpecParserMETHOD)
		}
		{
			p.SetState(47)
			p.Match(httpSpecParserWHITESPACE)
		}

	}
	{
		p.SetState(50)
		p.Match(httpSpecParserREQUESTTARGET)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserWHITESPACE {
		{
			p.SetState(51)
			p.Match(httpSpecParserWHITESPACE)
		}
		{
			p.SetState(52)
			p.Match(httpSpecParserHTTPVERSION)
		}

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
	p.RuleIndex = httpSpecParserRULE_messagebody
	return p
}

func (*MessagebodyContext) IsMessagebodyContext() {}

func NewMessagebodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagebodyContext {
	var p = new(MessagebodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_messagebody

	return p
}

func (s *MessagebodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagebodyContext) AllMESSAGES() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserMESSAGES)
}

func (s *MessagebodyContext) MESSAGES(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserMESSAGES, i)
}

func (s *MessagebodyContext) INPUTFILEREF() antlr.TerminalNode {
	return s.GetToken(httpSpecParserINPUTFILEREF, 0)
}

func (s *MessagebodyContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, 0)
}

func (s *MessagebodyContext) RESPONSEREF() antlr.TerminalNode {
	return s.GetToken(httpSpecParserRESPONSEREF, 0)
}

func (s *MessagebodyContext) RESPONSEHANDLER() antlr.TerminalNode {
	return s.GetToken(httpSpecParserRESPONSEHANDLER, 0)
}

func (s *MessagebodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagebodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagebodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMessagebody(s)
	}
}

func (s *MessagebodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMessagebody(s)
	}
}

func (p *httpSpecParser) Messagebody() (localctx IMessagebodyContext) {
	localctx = NewMessagebodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, httpSpecParserRULE_messagebody)
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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpSpecParserMESSAGES:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == httpSpecParserMESSAGES {
			{
				p.SetState(55)
				p.Match(httpSpecParserMESSAGES)
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case httpSpecParserINPUTFILEREF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(httpSpecParserINPUTFILEREF)
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == httpSpecParserNEWLINE {
			{
				p.SetState(61)
				p.Match(httpSpecParserNEWLINE)
			}

		}

	case httpSpecParserRESPONSEREF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(httpSpecParserRESPONSEREF)
		}

	case httpSpecParserRESPONSEHANDLER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Match(httpSpecParserRESPONSEHANDLER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = httpSpecParserRULE_headers
	return p
}

func (*HeadersContext) IsHeadersContext() {}

func NewHeadersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadersContext {
	var p = new(HeadersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_headers

	return p
}

func (s *HeadersContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadersContext) AllHEADERFIELD() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserHEADERFIELD)
}

func (s *HeadersContext) HEADERFIELD(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserHEADERFIELD, i)
}

func (s *HeadersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHeaders(s)
	}
}

func (s *HeadersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHeaders(s)
	}
}

func (p *httpSpecParser) Headers() (localctx IHeadersContext) {
	localctx = NewHeadersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, httpSpecParserRULE_headers)
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpSpecParserHEADERFIELD {
		{
			p.SetState(68)
			p.Match(httpSpecParserHEADERFIELD)
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
