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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 61, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11,
	2, 3, 2, 7, 2, 25, 10, 2, 12, 2, 14, 2, 28, 11, 2, 3, 2, 3, 2, 3, 3, 6,
	3, 33, 10, 3, 13, 3, 14, 3, 34, 3, 3, 3, 3, 3, 4, 6, 4, 40, 10, 4, 13,
	4, 14, 4, 41, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 48, 10, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 59, 10, 5, 3, 5, 2, 2, 6, 2,
	4, 6, 8, 2, 3, 3, 2, 3, 4, 2, 68, 2, 13, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2,
	6, 39, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 12, 7, 9, 2, 2, 11, 10, 3, 2,
	2, 2, 12, 15, 3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 16,
	3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 16, 20, 5, 6, 4, 2, 17, 19, 5, 4, 3, 2,
	18, 17, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3,
	2, 2, 2, 21, 26, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 25, 7, 9, 2, 2, 24,
	23, 3, 2, 2, 2, 25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2,
	2, 27, 29, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 29, 30, 7, 2, 2, 3, 30, 3, 3,
	2, 2, 2, 31, 33, 7, 9, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34,
	32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 5, 6, 4,
	2, 37, 5, 3, 2, 2, 2, 38, 40, 5, 8, 5, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3,
	2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 7, 3, 2, 2, 2, 43,
	44, 7, 10, 2, 2, 44, 59, 7, 3, 2, 2, 45, 47, 7, 7, 2, 2, 46, 48, 7, 3,
	2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 59, 3, 2, 2, 2, 49, 50,
	7, 12, 2, 2, 50, 59, 7, 3, 2, 2, 51, 52, 7, 14, 2, 2, 52, 59, 7, 3, 2,
	2, 53, 54, 7, 13, 2, 2, 54, 59, 7, 3, 2, 2, 55, 56, 7, 11, 2, 2, 56, 59,
	9, 2, 2, 2, 57, 59, 7, 3, 2, 2, 58, 43, 3, 2, 2, 2, 58, 45, 3, 2, 2, 2,
	58, 49, 3, 2, 2, 2, 58, 51, 3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 58, 55, 3,
	2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 9, 3, 2, 2, 2, 9, 13, 20, 26, 34, 41,
	47, 58,
}
var literalNames []string

var symbolicNames = []string{
	"", "NEWLINE", "NEWLINEWITHINDENT", "LINECOMMENT", "WHITESPACE", "HEADERFIELD",
	"ENVVARIABLE", "REQUESTSEPARATOR", "REQUESTLINE", "MESSAGE", "INPUTFILEREF",
	"RESPONSEHANDLER", "RESPONSEREF",
}

var ruleNames = []string{
	"file", "requestwithseparator", "request", "lines",
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
	httpSpecParserEOF               = antlr.TokenEOF
	httpSpecParserNEWLINE           = 1
	httpSpecParserNEWLINEWITHINDENT = 2
	httpSpecParserLINECOMMENT       = 3
	httpSpecParserWHITESPACE        = 4
	httpSpecParserHEADERFIELD       = 5
	httpSpecParserENVVARIABLE       = 6
	httpSpecParserREQUESTSEPARATOR  = 7
	httpSpecParserREQUESTLINE       = 8
	httpSpecParserMESSAGE           = 9
	httpSpecParserINPUTFILEREF      = 10
	httpSpecParserRESPONSEHANDLER   = 11
	httpSpecParserRESPONSEREF       = 12
)

// httpSpecParser rules.
const (
	httpSpecParserRULE_file                 = 0
	httpSpecParserRULE_requestwithseparator = 1
	httpSpecParserRULE_request              = 2
	httpSpecParserRULE_lines                = 3
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(8)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Request()
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(15)
				p.Requestwithseparator()
			}

		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(21)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(27)
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
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(29)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
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

func (s *RequestContext) AllLines() []ILinesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILinesContext)(nil)).Elem())
	var tst = make([]ILinesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILinesContext)
		}
	}

	return tst
}

func (s *RequestContext) Lines(i int) ILinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILinesContext)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserNEWLINE)|(1<<httpSpecParserHEADERFIELD)|(1<<httpSpecParserREQUESTLINE)|(1<<httpSpecParserMESSAGE)|(1<<httpSpecParserINPUTFILEREF)|(1<<httpSpecParserRESPONSEHANDLER)|(1<<httpSpecParserRESPONSEREF))) != 0) {
		{
			p.SetState(36)
			p.Lines()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILinesContext is an interface to support dynamic dispatch.
type ILinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinesContext differentiates from other interfaces.
	IsLinesContext()
}

type LinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinesContext() *LinesContext {
	var p = new(LinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_lines
	return p
}

func (*LinesContext) IsLinesContext() {}

func NewLinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinesContext {
	var p = new(LinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_lines

	return p
}

func (s *LinesContext) GetParser() antlr.Parser { return s.parser }

func (s *LinesContext) REQUESTLINE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUESTLINE, 0)
}

func (s *LinesContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, 0)
}

func (s *LinesContext) HEADERFIELD() antlr.TerminalNode {
	return s.GetToken(httpSpecParserHEADERFIELD, 0)
}

func (s *LinesContext) INPUTFILEREF() antlr.TerminalNode {
	return s.GetToken(httpSpecParserINPUTFILEREF, 0)
}

func (s *LinesContext) RESPONSEREF() antlr.TerminalNode {
	return s.GetToken(httpSpecParserRESPONSEREF, 0)
}

func (s *LinesContext) RESPONSEHANDLER() antlr.TerminalNode {
	return s.GetToken(httpSpecParserRESPONSEHANDLER, 0)
}

func (s *LinesContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserMESSAGE, 0)
}

func (s *LinesContext) NEWLINEWITHINDENT() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINEWITHINDENT, 0)
}

func (s *LinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterLines(s)
	}
}

func (s *LinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitLines(s)
	}
}

func (p *httpSpecParser) Lines() (localctx ILinesContext) {
	localctx = NewLinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, httpSpecParserRULE_lines)
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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpSpecParserREQUESTLINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(httpSpecParserREQUESTLINE)
		}
		{
			p.SetState(42)
			p.Match(httpSpecParserNEWLINE)
		}

	case httpSpecParserHEADERFIELD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Match(httpSpecParserHEADERFIELD)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(44)
				p.Match(httpSpecParserNEWLINE)
			}

		}

	case httpSpecParserINPUTFILEREF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.Match(httpSpecParserINPUTFILEREF)
		}
		{
			p.SetState(48)
			p.Match(httpSpecParserNEWLINE)
		}

	case httpSpecParserRESPONSEREF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(httpSpecParserRESPONSEREF)
		}
		{
			p.SetState(50)
			p.Match(httpSpecParserNEWLINE)
		}

	case httpSpecParserRESPONSEHANDLER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(51)
			p.Match(httpSpecParserRESPONSEHANDLER)
		}
		{
			p.SetState(52)
			p.Match(httpSpecParserNEWLINE)
		}

	case httpSpecParserMESSAGE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.Match(httpSpecParserMESSAGE)
		}
		{
			p.SetState(54)
			_la = p.GetTokenStream().LA(1)

			if !(_la == httpSpecParserNEWLINE || _la == httpSpecParserNEWLINEWITHINDENT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case httpSpecParserNEWLINE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(55)
			p.Match(httpSpecParserNEWLINE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
