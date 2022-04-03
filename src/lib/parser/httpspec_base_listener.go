// Code generated from ./src/lib/parser/httpSpec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpSpec

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasehttpSpecListener is a complete listener for a parse tree produced by httpSpecParser.
type BasehttpSpecListener struct{}

var _ httpSpecListener = &BasehttpSpecListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasehttpSpecListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasehttpSpecListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasehttpSpecListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasehttpSpecListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BasehttpSpecListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BasehttpSpecListener) ExitFile(ctx *FileContext) {}

// EnterRequestwithseparator is called when production requestwithseparator is entered.
func (s *BasehttpSpecListener) EnterRequestwithseparator(ctx *RequestwithseparatorContext) {}

// ExitRequestwithseparator is called when production requestwithseparator is exited.
func (s *BasehttpSpecListener) ExitRequestwithseparator(ctx *RequestwithseparatorContext) {}

// EnterRequest is called when production request is entered.
func (s *BasehttpSpecListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production request is exited.
func (s *BasehttpSpecListener) ExitRequest(ctx *RequestContext) {}

// EnterRequestline is called when production requestline is entered.
func (s *BasehttpSpecListener) EnterRequestline(ctx *RequestlineContext) {}

// ExitRequestline is called when production requestline is exited.
func (s *BasehttpSpecListener) ExitRequestline(ctx *RequestlineContext) {}

// EnterMessagebody is called when production messagebody is entered.
func (s *BasehttpSpecListener) EnterMessagebody(ctx *MessagebodyContext) {}

// ExitMessagebody is called when production messagebody is exited.
func (s *BasehttpSpecListener) ExitMessagebody(ctx *MessagebodyContext) {}

// EnterHeaders is called when production headers is entered.
func (s *BasehttpSpecListener) EnterHeaders(ctx *HeadersContext) {}

// ExitHeaders is called when production headers is exited.
func (s *BasehttpSpecListener) ExitHeaders(ctx *HeadersContext) {}
