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

	// EnterLines is called when entering the lines production.
	EnterLines(c *LinesContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitRequestwithseparator is called when exiting the requestwithseparator production.
	ExitRequestwithseparator(c *RequestwithseparatorContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitLines is called when exiting the lines production.
	ExitLines(c *LinesContext)
}
