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

// EnterRequestfile is called when production requestfile is entered.
func (s *BasehttpSpecListener) EnterRequestfile(ctx *RequestfileContext) {}

// ExitRequestfile is called when production requestfile is exited.
func (s *BasehttpSpecListener) ExitRequestfile(ctx *RequestfileContext) {}

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

// EnterMethod is called when production method is entered.
func (s *BasehttpSpecListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BasehttpSpecListener) ExitMethod(ctx *MethodContext) {}

// EnterHttpversion is called when production httpversion is entered.
func (s *BasehttpSpecListener) EnterHttpversion(ctx *HttpversionContext) {}

// ExitHttpversion is called when production httpversion is exited.
func (s *BasehttpSpecListener) ExitHttpversion(ctx *HttpversionContext) {}

// EnterRequesttarget is called when production requesttarget is entered.
func (s *BasehttpSpecListener) EnterRequesttarget(ctx *RequesttargetContext) {}

// ExitRequesttarget is called when production requesttarget is exited.
func (s *BasehttpSpecListener) ExitRequesttarget(ctx *RequesttargetContext) {}

// EnterOriginform is called when production originform is entered.
func (s *BasehttpSpecListener) EnterOriginform(ctx *OriginformContext) {}

// ExitOriginform is called when production originform is exited.
func (s *BasehttpSpecListener) ExitOriginform(ctx *OriginformContext) {}

// EnterAbsoluteform is called when production absoluteform is entered.
func (s *BasehttpSpecListener) EnterAbsoluteform(ctx *AbsoluteformContext) {}

// ExitAbsoluteform is called when production absoluteform is exited.
func (s *BasehttpSpecListener) ExitAbsoluteform(ctx *AbsoluteformContext) {}

// EnterScheme is called when production scheme is entered.
func (s *BasehttpSpecListener) EnterScheme(ctx *SchemeContext) {}

// ExitScheme is called when production scheme is exited.
func (s *BasehttpSpecListener) ExitScheme(ctx *SchemeContext) {}

// EnterHierpart is called when production hierpart is entered.
func (s *BasehttpSpecListener) EnterHierpart(ctx *HierpartContext) {}

// ExitHierpart is called when production hierpart is exited.
func (s *BasehttpSpecListener) ExitHierpart(ctx *HierpartContext) {}

// EnterAsteriskform is called when production asteriskform is entered.
func (s *BasehttpSpecListener) EnterAsteriskform(ctx *AsteriskformContext) {}

// ExitAsteriskform is called when production asteriskform is exited.
func (s *BasehttpSpecListener) ExitAsteriskform(ctx *AsteriskformContext) {}

// EnterAuthority is called when production authority is entered.
func (s *BasehttpSpecListener) EnterAuthority(ctx *AuthorityContext) {}

// ExitAuthority is called when production authority is exited.
func (s *BasehttpSpecListener) ExitAuthority(ctx *AuthorityContext) {}

// EnterPort is called when production port is entered.
func (s *BasehttpSpecListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BasehttpSpecListener) ExitPort(ctx *PortContext) {}

// EnterHost is called when production host is entered.
func (s *BasehttpSpecListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BasehttpSpecListener) ExitHost(ctx *HostContext) {}

// EnterIpv6address is called when production ipv6address is entered.
func (s *BasehttpSpecListener) EnterIpv6address(ctx *Ipv6addressContext) {}

// ExitIpv6address is called when production ipv6address is exited.
func (s *BasehttpSpecListener) ExitIpv6address(ctx *Ipv6addressContext) {}

// EnterIpv4addressorregname is called when production ipv4addressorregname is entered.
func (s *BasehttpSpecListener) EnterIpv4addressorregname(ctx *Ipv4addressorregnameContext) {}

// ExitIpv4addressorregname is called when production ipv4addressorregname is exited.
func (s *BasehttpSpecListener) ExitIpv4addressorregname(ctx *Ipv4addressorregnameContext) {}

// EnterAbsolutepath is called when production absolutepath is entered.
func (s *BasehttpSpecListener) EnterAbsolutepath(ctx *AbsolutepathContext) {}

// ExitAbsolutepath is called when production absolutepath is exited.
func (s *BasehttpSpecListener) ExitAbsolutepath(ctx *AbsolutepathContext) {}

// EnterPathseparator is called when production pathseparator is entered.
func (s *BasehttpSpecListener) EnterPathseparator(ctx *PathseparatorContext) {}

// ExitPathseparator is called when production pathseparator is exited.
func (s *BasehttpSpecListener) ExitPathseparator(ctx *PathseparatorContext) {}

// EnterSegment is called when production segment is entered.
func (s *BasehttpSpecListener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BasehttpSpecListener) ExitSegment(ctx *SegmentContext) {}

// EnterQuery is called when production query is entered.
func (s *BasehttpSpecListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasehttpSpecListener) ExitQuery(ctx *QueryContext) {}

// EnterRequestfragment is called when production requestfragment is entered.
func (s *BasehttpSpecListener) EnterRequestfragment(ctx *RequestfragmentContext) {}

// ExitRequestfragment is called when production requestfragment is exited.
func (s *BasehttpSpecListener) ExitRequestfragment(ctx *RequestfragmentContext) {}

// EnterHeaders is called when production headers is entered.
func (s *BasehttpSpecListener) EnterHeaders(ctx *HeadersContext) {}

// ExitHeaders is called when production headers is exited.
func (s *BasehttpSpecListener) ExitHeaders(ctx *HeadersContext) {}

// EnterHeaderfield is called when production headerfield is entered.
func (s *BasehttpSpecListener) EnterHeaderfield(ctx *HeaderfieldContext) {}

// ExitHeaderfield is called when production headerfield is exited.
func (s *BasehttpSpecListener) ExitHeaderfield(ctx *HeaderfieldContext) {}

// EnterFieldname is called when production fieldname is entered.
func (s *BasehttpSpecListener) EnterFieldname(ctx *FieldnameContext) {}

// ExitFieldname is called when production fieldname is exited.
func (s *BasehttpSpecListener) ExitFieldname(ctx *FieldnameContext) {}

// EnterFieldvalue is called when production fieldvalue is entered.
func (s *BasehttpSpecListener) EnterFieldvalue(ctx *FieldvalueContext) {}

// ExitFieldvalue is called when production fieldvalue is exited.
func (s *BasehttpSpecListener) ExitFieldvalue(ctx *FieldvalueContext) {}

// EnterMessagebody is called when production messagebody is entered.
func (s *BasehttpSpecListener) EnterMessagebody(ctx *MessagebodyContext) {}

// ExitMessagebody is called when production messagebody is exited.
func (s *BasehttpSpecListener) ExitMessagebody(ctx *MessagebodyContext) {}

// EnterMessages is called when production messages is entered.
func (s *BasehttpSpecListener) EnterMessages(ctx *MessagesContext) {}

// ExitMessages is called when production messages is exited.
func (s *BasehttpSpecListener) ExitMessages(ctx *MessagesContext) {}

// EnterMessageline is called when production messageline is entered.
func (s *BasehttpSpecListener) EnterMessageline(ctx *MessagelineContext) {}

// ExitMessageline is called when production messageline is exited.
func (s *BasehttpSpecListener) ExitMessageline(ctx *MessagelineContext) {}

// EnterInputfileref is called when production inputfileref is entered.
func (s *BasehttpSpecListener) EnterInputfileref(ctx *InputfilerefContext) {}

// ExitInputfileref is called when production inputfileref is exited.
func (s *BasehttpSpecListener) ExitInputfileref(ctx *InputfilerefContext) {}

// EnterFilepath is called when production filepath is entered.
func (s *BasehttpSpecListener) EnterFilepath(ctx *FilepathContext) {}

// ExitFilepath is called when production filepath is exited.
func (s *BasehttpSpecListener) ExitFilepath(ctx *FilepathContext) {}

// EnterMultipartformdata is called when production multipartformdata is entered.
func (s *BasehttpSpecListener) EnterMultipartformdata(ctx *MultipartformdataContext) {}

// ExitMultipartformdata is called when production multipartformdata is exited.
func (s *BasehttpSpecListener) ExitMultipartformdata(ctx *MultipartformdataContext) {}

// EnterMultipartfield is called when production multipartfield is entered.
func (s *BasehttpSpecListener) EnterMultipartfield(ctx *MultipartfieldContext) {}

// ExitMultipartfield is called when production multipartfield is exited.
func (s *BasehttpSpecListener) ExitMultipartfield(ctx *MultipartfieldContext) {}

// EnterBoundary is called when production boundary is entered.
func (s *BasehttpSpecListener) EnterBoundary(ctx *BoundaryContext) {}

// ExitBoundary is called when production boundary is exited.
func (s *BasehttpSpecListener) ExitBoundary(ctx *BoundaryContext) {}

// EnterResponsehandler is called when production responsehandler is entered.
func (s *BasehttpSpecListener) EnterResponsehandler(ctx *ResponsehandlerContext) {}

// ExitResponsehandler is called when production responsehandler is exited.
func (s *BasehttpSpecListener) ExitResponsehandler(ctx *ResponsehandlerContext) {}

// EnterHandlerscript is called when production handlerscript is entered.
func (s *BasehttpSpecListener) EnterHandlerscript(ctx *HandlerscriptContext) {}

// ExitHandlerscript is called when production handlerscript is exited.
func (s *BasehttpSpecListener) ExitHandlerscript(ctx *HandlerscriptContext) {}

// EnterResponseref is called when production responseref is entered.
func (s *BasehttpSpecListener) EnterResponseref(ctx *ResponserefContext) {}

// ExitResponseref is called when production responseref is exited.
func (s *BasehttpSpecListener) ExitResponseref(ctx *ResponserefContext) {}

// EnterEnvvariable is called when production envvariable is entered.
func (s *BasehttpSpecListener) EnterEnvvariable(ctx *EnvvariableContext) {}

// ExitEnvvariable is called when production envvariable is exited.
func (s *BasehttpSpecListener) ExitEnvvariable(ctx *EnvvariableContext) {}
