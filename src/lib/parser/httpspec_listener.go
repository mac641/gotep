// Code generated from ./src/lib/parser/httpSpec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpSpec

import "github.com/antlr/antlr4/runtime/Go/antlr"

// httpSpecListener is a complete listener for a parse tree produced by httpSpecParser.
type httpSpecListener interface {
	antlr.ParseTreeListener

	// EnterRequestfile is called when entering the requestfile production.
	EnterRequestfile(c *RequestfileContext)

	// EnterRequestwithseparator is called when entering the requestwithseparator production.
	EnterRequestwithseparator(c *RequestwithseparatorContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterRequestline is called when entering the requestline production.
	EnterRequestline(c *RequestlineContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterHttpversion is called when entering the httpversion production.
	EnterHttpversion(c *HttpversionContext)

	// EnterRequesttarget is called when entering the requesttarget production.
	EnterRequesttarget(c *RequesttargetContext)

	// EnterOriginform is called when entering the originform production.
	EnterOriginform(c *OriginformContext)

	// EnterAbsoluteform is called when entering the absoluteform production.
	EnterAbsoluteform(c *AbsoluteformContext)

	// EnterScheme is called when entering the scheme production.
	EnterScheme(c *SchemeContext)

	// EnterHierpart is called when entering the hierpart production.
	EnterHierpart(c *HierpartContext)

	// EnterAsteriskform is called when entering the asteriskform production.
	EnterAsteriskform(c *AsteriskformContext)

	// EnterAuthority is called when entering the authority production.
	EnterAuthority(c *AuthorityContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterIpv6address is called when entering the ipv6address production.
	EnterIpv6address(c *Ipv6addressContext)

	// EnterIpv4addressorregname is called when entering the ipv4addressorregname production.
	EnterIpv4addressorregname(c *Ipv4addressorregnameContext)

	// EnterAbsolutepath is called when entering the absolutepath production.
	EnterAbsolutepath(c *AbsolutepathContext)

	// EnterPathseparator is called when entering the pathseparator production.
	EnterPathseparator(c *PathseparatorContext)

	// EnterSegment is called when entering the segment production.
	EnterSegment(c *SegmentContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterRequestfragment is called when entering the requestfragment production.
	EnterRequestfragment(c *RequestfragmentContext)

	// EnterHeaders is called when entering the headers production.
	EnterHeaders(c *HeadersContext)

	// EnterHeaderfield is called when entering the headerfield production.
	EnterHeaderfield(c *HeaderfieldContext)

	// EnterFieldname is called when entering the fieldname production.
	EnterFieldname(c *FieldnameContext)

	// EnterFieldvalue is called when entering the fieldvalue production.
	EnterFieldvalue(c *FieldvalueContext)

	// EnterMessagebody is called when entering the messagebody production.
	EnterMessagebody(c *MessagebodyContext)

	// EnterMessages is called when entering the messages production.
	EnterMessages(c *MessagesContext)

	// EnterMessageline is called when entering the messageline production.
	EnterMessageline(c *MessagelineContext)

	// EnterInputfileref is called when entering the inputfileref production.
	EnterInputfileref(c *InputfilerefContext)

	// EnterFilepath is called when entering the filepath production.
	EnterFilepath(c *FilepathContext)

	// EnterMultipartformdata is called when entering the multipartformdata production.
	EnterMultipartformdata(c *MultipartformdataContext)

	// EnterMultipartfield is called when entering the multipartfield production.
	EnterMultipartfield(c *MultipartfieldContext)

	// EnterBoundary is called when entering the boundary production.
	EnterBoundary(c *BoundaryContext)

	// EnterResponsehandler is called when entering the responsehandler production.
	EnterResponsehandler(c *ResponsehandlerContext)

	// EnterHandlerscript is called when entering the handlerscript production.
	EnterHandlerscript(c *HandlerscriptContext)

	// EnterResponseref is called when entering the responseref production.
	EnterResponseref(c *ResponserefContext)

	// EnterEnvvariable is called when entering the envvariable production.
	EnterEnvvariable(c *EnvvariableContext)

	// ExitRequestfile is called when exiting the requestfile production.
	ExitRequestfile(c *RequestfileContext)

	// ExitRequestwithseparator is called when exiting the requestwithseparator production.
	ExitRequestwithseparator(c *RequestwithseparatorContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitRequestline is called when exiting the requestline production.
	ExitRequestline(c *RequestlineContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitHttpversion is called when exiting the httpversion production.
	ExitHttpversion(c *HttpversionContext)

	// ExitRequesttarget is called when exiting the requesttarget production.
	ExitRequesttarget(c *RequesttargetContext)

	// ExitOriginform is called when exiting the originform production.
	ExitOriginform(c *OriginformContext)

	// ExitAbsoluteform is called when exiting the absoluteform production.
	ExitAbsoluteform(c *AbsoluteformContext)

	// ExitScheme is called when exiting the scheme production.
	ExitScheme(c *SchemeContext)

	// ExitHierpart is called when exiting the hierpart production.
	ExitHierpart(c *HierpartContext)

	// ExitAsteriskform is called when exiting the asteriskform production.
	ExitAsteriskform(c *AsteriskformContext)

	// ExitAuthority is called when exiting the authority production.
	ExitAuthority(c *AuthorityContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitIpv6address is called when exiting the ipv6address production.
	ExitIpv6address(c *Ipv6addressContext)

	// ExitIpv4addressorregname is called when exiting the ipv4addressorregname production.
	ExitIpv4addressorregname(c *Ipv4addressorregnameContext)

	// ExitAbsolutepath is called when exiting the absolutepath production.
	ExitAbsolutepath(c *AbsolutepathContext)

	// ExitPathseparator is called when exiting the pathseparator production.
	ExitPathseparator(c *PathseparatorContext)

	// ExitSegment is called when exiting the segment production.
	ExitSegment(c *SegmentContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitRequestfragment is called when exiting the requestfragment production.
	ExitRequestfragment(c *RequestfragmentContext)

	// ExitHeaders is called when exiting the headers production.
	ExitHeaders(c *HeadersContext)

	// ExitHeaderfield is called when exiting the headerfield production.
	ExitHeaderfield(c *HeaderfieldContext)

	// ExitFieldname is called when exiting the fieldname production.
	ExitFieldname(c *FieldnameContext)

	// ExitFieldvalue is called when exiting the fieldvalue production.
	ExitFieldvalue(c *FieldvalueContext)

	// ExitMessagebody is called when exiting the messagebody production.
	ExitMessagebody(c *MessagebodyContext)

	// ExitMessages is called when exiting the messages production.
	ExitMessages(c *MessagesContext)

	// ExitMessageline is called when exiting the messageline production.
	ExitMessageline(c *MessagelineContext)

	// ExitInputfileref is called when exiting the inputfileref production.
	ExitInputfileref(c *InputfilerefContext)

	// ExitFilepath is called when exiting the filepath production.
	ExitFilepath(c *FilepathContext)

	// ExitMultipartformdata is called when exiting the multipartformdata production.
	ExitMultipartformdata(c *MultipartformdataContext)

	// ExitMultipartfield is called when exiting the multipartfield production.
	ExitMultipartfield(c *MultipartfieldContext)

	// ExitBoundary is called when exiting the boundary production.
	ExitBoundary(c *BoundaryContext)

	// ExitResponsehandler is called when exiting the responsehandler production.
	ExitResponsehandler(c *ResponsehandlerContext)

	// ExitHandlerscript is called when exiting the handlerscript production.
	ExitHandlerscript(c *HandlerscriptContext)

	// ExitResponseref is called when exiting the responseref production.
	ExitResponseref(c *ResponserefContext)

	// ExitEnvvariable is called when exiting the envvariable production.
	ExitEnvvariable(c *EnvvariableContext)
}
