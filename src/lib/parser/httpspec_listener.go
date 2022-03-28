// Code generated from ./src/lib/parser/httpspec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpspec

import "github.com/antlr/antlr4/runtime/Go/antlr"

// httpspecListener is a complete listener for a parse tree produced by httpspecParser.
type httpspecListener interface {
	antlr.ParseTreeListener

	// EnterRequests is called when entering the requests production.
	EnterRequests(c *RequestsContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterRequestline is called when entering the requestline production.
	EnterRequestline(c *RequestlineContext)

	// EnterRequesttarget is called when entering the requesttarget production.
	EnterRequesttarget(c *RequesttargetContext)

	// EnterAbsoluteform is called when entering the absoluteform production.
	EnterAbsoluteform(c *AbsoluteformContext)

	// EnterHierpart is called when entering the hierpart production.
	EnterHierpart(c *HierpartContext)

	// EnterAuthority is called when entering the authority production.
	EnterAuthority(c *AuthorityContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterIpv6address is called when entering the ipv6address production.
	EnterIpv6address(c *Ipv6addressContext)

	// EnterIpv4addressorregname is called when entering the ipv4addressorregname production.
	EnterIpv4addressorregname(c *Ipv4addressorregnameContext)

	// EnterHeaders is called when entering the headers production.
	EnterHeaders(c *HeadersContext)

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

	// ExitRequests is called when exiting the requests production.
	ExitRequests(c *RequestsContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitRequestline is called when exiting the requestline production.
	ExitRequestline(c *RequestlineContext)

	// ExitRequesttarget is called when exiting the requesttarget production.
	ExitRequesttarget(c *RequesttargetContext)

	// ExitAbsoluteform is called when exiting the absoluteform production.
	ExitAbsoluteform(c *AbsoluteformContext)

	// ExitHierpart is called when exiting the hierpart production.
	ExitHierpart(c *HierpartContext)

	// ExitAuthority is called when exiting the authority production.
	ExitAuthority(c *AuthorityContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitIpv6address is called when exiting the ipv6address production.
	ExitIpv6address(c *Ipv6addressContext)

	// ExitIpv4addressorregname is called when exiting the ipv4addressorregname production.
	ExitIpv4addressorregname(c *Ipv4addressorregnameContext)

	// ExitHeaders is called when exiting the headers production.
	ExitHeaders(c *HeadersContext)

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
}
