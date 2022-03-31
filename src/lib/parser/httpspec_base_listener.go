// Code generated from ./src/lib/parser/httpspec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpspec

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasehttpspecListener is a complete listener for a parse tree produced by httpspecParser.
type BasehttpspecListener struct{}

var _ httpspecListener = &BasehttpspecListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasehttpspecListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasehttpspecListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasehttpspecListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasehttpspecListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BasehttpspecListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BasehttpspecListener) ExitFile(ctx *FileContext) {}

// EnterRequestwithseparator is called when production requestwithseparator is entered.
func (s *BasehttpspecListener) EnterRequestwithseparator(ctx *RequestwithseparatorContext) {}

// ExitRequestwithseparator is called when production requestwithseparator is exited.
func (s *BasehttpspecListener) ExitRequestwithseparator(ctx *RequestwithseparatorContext) {}

// EnterRequest is called when production request is entered.
func (s *BasehttpspecListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production request is exited.
func (s *BasehttpspecListener) ExitRequest(ctx *RequestContext) {}

// EnterRequestline is called when production requestline is entered.
func (s *BasehttpspecListener) EnterRequestline(ctx *RequestlineContext) {}

// ExitRequestline is called when production requestline is exited.
func (s *BasehttpspecListener) ExitRequestline(ctx *RequestlineContext) {}

// EnterRequesttarget is called when production requesttarget is entered.
func (s *BasehttpspecListener) EnterRequesttarget(ctx *RequesttargetContext) {}

// ExitRequesttarget is called when production requesttarget is exited.
func (s *BasehttpspecListener) ExitRequesttarget(ctx *RequesttargetContext) {}

// EnterHeaders is called when production headers is entered.
func (s *BasehttpspecListener) EnterHeaders(ctx *HeadersContext) {}

// ExitHeaders is called when production headers is exited.
func (s *BasehttpspecListener) ExitHeaders(ctx *HeadersContext) {}

// EnterFilepath is called when production filepath is entered.
func (s *BasehttpspecListener) EnterFilepath(ctx *FilepathContext) {}

// ExitFilepath is called when production filepath is exited.
func (s *BasehttpspecListener) ExitFilepath(ctx *FilepathContext) {}

// EnterInputfileref is called when production inputfileref is entered.
func (s *BasehttpspecListener) EnterInputfileref(ctx *InputfilerefContext) {}

// ExitInputfileref is called when production inputfileref is exited.
func (s *BasehttpspecListener) ExitInputfileref(ctx *InputfilerefContext) {}

// EnterMessagebody is called when production messagebody is entered.
func (s *BasehttpspecListener) EnterMessagebody(ctx *MessagebodyContext) {}

// ExitMessagebody is called when production messagebody is exited.
func (s *BasehttpspecListener) ExitMessagebody(ctx *MessagebodyContext) {}

// EnterMessages is called when production messages is entered.
func (s *BasehttpspecListener) EnterMessages(ctx *MessagesContext) {}

// ExitMessages is called when production messages is exited.
func (s *BasehttpspecListener) ExitMessages(ctx *MessagesContext) {}

// EnterMessageline is called when production messageline is entered.
func (s *BasehttpspecListener) EnterMessageline(ctx *MessagelineContext) {}

// ExitMessageline is called when production messageline is exited.
func (s *BasehttpspecListener) ExitMessageline(ctx *MessagelineContext) {}

// EnterMultipartformdata is called when production multipartformdata is entered.
func (s *BasehttpspecListener) EnterMultipartformdata(ctx *MultipartformdataContext) {}

// ExitMultipartformdata is called when production multipartformdata is exited.
func (s *BasehttpspecListener) ExitMultipartformdata(ctx *MultipartformdataContext) {}

// EnterMultipartfield is called when production multipartfield is entered.
func (s *BasehttpspecListener) EnterMultipartfield(ctx *MultipartfieldContext) {}

// ExitMultipartfield is called when production multipartfield is exited.
func (s *BasehttpspecListener) ExitMultipartfield(ctx *MultipartfieldContext) {}

// EnterHandlerscript is called when production handlerscript is entered.
func (s *BasehttpspecListener) EnterHandlerscript(ctx *HandlerscriptContext) {}

// ExitHandlerscript is called when production handlerscript is exited.
func (s *BasehttpspecListener) ExitHandlerscript(ctx *HandlerscriptContext) {}

// EnterResponsehandler is called when production responsehandler is entered.
func (s *BasehttpspecListener) EnterResponsehandler(ctx *ResponsehandlerContext) {}

// ExitResponsehandler is called when production responsehandler is exited.
func (s *BasehttpspecListener) ExitResponsehandler(ctx *ResponsehandlerContext) {}

// EnterResponseref is called when production responseref is entered.
func (s *BasehttpspecListener) EnterResponseref(ctx *ResponserefContext) {}

// ExitResponseref is called when production responseref is exited.
func (s *BasehttpspecListener) ExitResponseref(ctx *ResponserefContext) {}
