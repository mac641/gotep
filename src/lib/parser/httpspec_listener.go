// Code generated from ./src/lib/parser/httpSpec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpSpec

import "github.com/antlr/antlr4/runtime/Go/antlr"

// httpSpecListener is a complete listener for a parse tree produced by httpSpecParser.
type httpSpecListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterRequestwithseparator is called when entering the requestwithseparator production.
	EnterRequestwithseparator(c *RequestwithseparatorContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterRequestline is called when entering the requestline production.
	EnterRequestline(c *RequestlineContext)

	// EnterMessagebody is called when entering the messagebody production.
	EnterMessagebody(c *MessagebodyContext)

	// EnterHeaders is called when entering the headers production.
	EnterHeaders(c *HeadersContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitRequestwithseparator is called when exiting the requestwithseparator production.
	ExitRequestwithseparator(c *RequestwithseparatorContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitRequestline is called when exiting the requestline production.
	ExitRequestline(c *RequestlineContext)

	// ExitMessagebody is called when exiting the messagebody production.
	ExitMessagebody(c *MessagebodyContext)

	// ExitHeaders is called when exiting the headers production.
	ExitHeaders(c *HeadersContext)
}
