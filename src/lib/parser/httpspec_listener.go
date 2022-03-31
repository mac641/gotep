// Code generated from ./src/lib/parser/httpspec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpspec

import "github.com/antlr/antlr4/runtime/Go/antlr"

// httpspecListener is a complete listener for a parse tree produced by httpspecParser.
type httpspecListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterRequestwithseparator is called when entering the requestwithseparator production.
	EnterRequestwithseparator(c *RequestwithseparatorContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterRequestline is called when entering the requestline production.
	EnterRequestline(c *RequestlineContext)

	// EnterRequesttarget is called when entering the requesttarget production.
	EnterRequesttarget(c *RequesttargetContext)

	// EnterHeaders is called when entering the headers production.
	EnterHeaders(c *HeadersContext)

	// EnterFilepath is called when entering the filepath production.
	EnterFilepath(c *FilepathContext)

	// EnterInputfileref is called when entering the inputfileref production.
	EnterInputfileref(c *InputfilerefContext)

	// EnterMessagebody is called when entering the messagebody production.
	EnterMessagebody(c *MessagebodyContext)

	// EnterMessages is called when entering the messages production.
	EnterMessages(c *MessagesContext)

	// EnterMessageline is called when entering the messageline production.
	EnterMessageline(c *MessagelineContext)

	// EnterMultipartformdata is called when entering the multipartformdata production.
	EnterMultipartformdata(c *MultipartformdataContext)

	// EnterMultipartfield is called when entering the multipartfield production.
	EnterMultipartfield(c *MultipartfieldContext)

	// EnterHandlerscript is called when entering the handlerscript production.
	EnterHandlerscript(c *HandlerscriptContext)

	// EnterResponsehandler is called when entering the responsehandler production.
	EnterResponsehandler(c *ResponsehandlerContext)

	// EnterResponseref is called when entering the responseref production.
	EnterResponseref(c *ResponserefContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitRequestwithseparator is called when exiting the requestwithseparator production.
	ExitRequestwithseparator(c *RequestwithseparatorContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitRequestline is called when exiting the requestline production.
	ExitRequestline(c *RequestlineContext)

	// ExitRequesttarget is called when exiting the requesttarget production.
	ExitRequesttarget(c *RequesttargetContext)

	// ExitHeaders is called when exiting the headers production.
	ExitHeaders(c *HeadersContext)

	// ExitFilepath is called when exiting the filepath production.
	ExitFilepath(c *FilepathContext)

	// ExitInputfileref is called when exiting the inputfileref production.
	ExitInputfileref(c *InputfilerefContext)

	// ExitMessagebody is called when exiting the messagebody production.
	ExitMessagebody(c *MessagebodyContext)

	// ExitMessages is called when exiting the messages production.
	ExitMessages(c *MessagesContext)

	// ExitMessageline is called when exiting the messageline production.
	ExitMessageline(c *MessagelineContext)

	// ExitMultipartformdata is called when exiting the multipartformdata production.
	ExitMultipartformdata(c *MultipartformdataContext)

	// ExitMultipartfield is called when exiting the multipartfield production.
	ExitMultipartfield(c *MultipartfieldContext)

	// ExitHandlerscript is called when exiting the handlerscript production.
	ExitHandlerscript(c *HandlerscriptContext)

	// ExitResponsehandler is called when exiting the responsehandler production.
	ExitResponsehandler(c *ResponsehandlerContext)

	// ExitResponseref is called when exiting the responseref production.
	ExitResponseref(c *ResponserefContext)
}
