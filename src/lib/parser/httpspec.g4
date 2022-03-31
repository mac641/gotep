grammar httpspec;

// Add additional lexer member to ignoreWs whitespace if necessary
@lexer::members {
   var ignoreWs = false
}

// Line terminators
fragment CR: '\r'; // \r
fragment LF: '\n'; // \n

NEWLINE: CR? LF | CR;
NEWLINEWITHINDENT: NEWLINE WHITESPACE;

// List of multiple requests
file:
	REQUESTSEPARATOR* request requestwithseparator* REQUESTSEPARATOR* EOF;
requestwithseparator: REQUESTSEPARATOR+ request;

// Requests
fragment ASTERISK: '\u002A'; // *
fragment FULLSTOP: '\u002E'; // .
fragment SEGMENT:
	~(
		'\u002F'
		| '\u003F'
		| '\u0023'
		| [\u000D\u000A\u0020\u0009\u000C]
	)+; // SLASH | QUESTIONMARK | HASHTAG | [\r\n \t\f]
fragment SCHEME: 'http' | 'https';
fragment PATHSEPARATOR: SLASH | NEWLINEWITHINDENT;
fragment ABSOLUTEPATH: SLASH | (PATHSEPARATOR SEGMENT)+;
fragment HIERPART: AUTHORITY ABSOLUTEPATH?;
fragment AUTHORITY: HOST (COLON PORT)?;
fragment PORT: DIGIT DIGIT DIGIT DIGIT DIGIT;
fragment HOST:
	LEFTSQUAREBRACKET IPV6ADDRESS RIGHTSQUAREBRACKET
	| IPV4ADDRESSORREGNAME;
// TODO: invetigate regname and specify it further
fragment IPV4ADDRESSORREGNAME:
	DIGIT DIGIT? DIGIT? FULLSTOP DIGIT DIGIT? DIGIT? FULLSTOP DIGIT DIGIT? DIGIT? FULLSTOP DIGIT
		DIGIT? DIGIT?;
// | ~('\u002F' | '\u003A' | '\u003F' | '\u0023')+; // SLASH | COLON | QUESTIONMARK | HASHTAG
fragment IPV6ADDRESS: (HEX | COLON)*;
// ((HEX | COLON)* HEX+)* COLON HEX* COLON (HEX+ (HEX | COLON)*)*; TODO: figure out if this matches
// all ipv6 examples
fragment QUERY: ~'\u0023' (NEWLINEWITHINDENT QUERY)?; // HASHTAG
fragment REQUESTFRAGMENT:
	~'\u003F' (NEWLINEWITHINDENT REQUESTFRAGMENT)?; // QUESTIONMARK

COLON: '\u003A'; // :
QUESTIONMARK: '\u003F'; // ?
LEFTSQUAREBRACKET: '\u005B'; // [
RIGHTSQUAREBRACKET: '\u005D'; // ]
METHOD:
	'GET'
	| 'HEAD'
	| 'POST'
	| 'PUT'
	| 'DELETE'
	| 'CONNECT'
	| 'PATCH'
	| 'OPTIONS'
	| 'TRACE';
HTTPVERSION: 'HTTP/' DIGIT+ '.' DIGIT+;
REQUESTSEPARATOR: '###' ~('\r' | '\n')* NEWLINE;
ASTERISKFORM: ASTERISK;
ABSOLUTEFORM: (SCHEME '://')? HIERPART ('?' QUERY+)? (
		'#' REQUESTFRAGMENT
	)?;
ORIGINFORM: ABSOLUTEPATH ('?' QUERY+)? ('#' REQUESTFRAGMENT)?;

request:
	requestline NEWLINE headers NEWLINE messagebody? responsehandler? responseref?;
requestline: (METHOD WHITESPACE)? requesttarget (
		WHITESPACE HTTPVERSION
	)?;
requesttarget: ORIGINFORM | ABSOLUTEFORM | ASTERISKFORM;

// Comments
SLASH: '\u002F'; // /
HASHTAG: '\u0023'; // #
LINECOMMENT: (
		'#' ~[\u000D\u000A]* NEWLINE
		| '//' ~[\u000D\u000A]* NEWLINE
	) -> skip;
// NOTE: needs to be specified after REQUESTSEPARATOR, to ensure it not being skipped as regular comment

// Headers
fragment FIELDVALUE: (ID SLASH?)* NEWLINE? (
		NEWLINEWITHINDENT FIELDVALUE
	)?;
fragment FIELDNAME:
	~('\u003A' | [\u000D\u000A\u0020\u0009\u000C])+; // COLON | [\r\n \t\f]

HEADERFIELD:
	FIELDNAME COLON {ignoreWs = true } WHITESPACE* {ignoreWs = false } FIELDVALUE {ignoreWs = true
			} WHITESPACE* {ignoreWs = false };

headers: (HEADERFIELD NEWLINE?)*;

// Message body
LEFTCURLYBRACKET: '\u007B'; // {
RIGHTCURLYBRACKET: '\u007D'; // }
PERCENTSIGN: '\u0025'; // %
LESSTHANSIGN: '\u003C'; // <
GREATERTHANSIGN: '\u003E'; // >
BOUNDARY: (HYPHEN HYPHEN)+ ID (HYPHEN HYPHEN)? NEWLINE;

filepath: INPUT* NEWLINE;
inputfileref: LESSTHANSIGN WHITESPACE filepath;
messagebody: messages | multipartformdata;
messages: messageline (NEWLINE messageline)?;
messageline:
	~('<' | '<>' | '###') INPUT* NEWLINE
	| inputfileref;
// suport unicode characters as part of message
multipartformdata: multipartfield multipartformdata? BOUNDARY;
multipartfield:
	BOUNDARY (HEADERFIELD NEWLINE)* NEWLINE messages?;

// Chars: every unicode character except NEWLINE is supported
fragment ALPHA: [\u0041-\u005A\u0061-\u007A]; // A-Z, a-z
fragment DIGIT: [\u0030-\u0039]; // 0-9
fragment HEX: [0-9A-Fa-f];
fragment HYPHEN: '\u002D'; // -
fragment UNDERSCORE: '\u005F'; // _
fragment ID: (ALPHA | DIGIT | HYPHEN | UNDERSCORE);

INPUT: ~[\u000D\u000A\u0020\u0009\u000C]; // [\r\n \t\f]

// Whitespaces
WHITESPACE:
	[ \t\f]+ { if ignoreWs {
        l.Skip()
    } }; // space, horizontal tab (\t), form feed (\f)

// Response handler
handlerscript: INPUT+ NEWLINE;
// TODO: prompt error message if used anyways
responsehandler: // TODO: use parser rule when this will be implemented
	GREATERTHANSIGN WHITESPACE LEFTCURLYBRACKET PERCENTSIGN handlerscript PERCENTSIGN
		RIGHTCURLYBRACKET
	| GREATERTHANSIGN WHITESPACE filepath;

// Response reference
responseref: LESSTHANSIGN GREATERTHANSIGN WHITESPACE filepath;

// Environment variables
ENVVARIABLE:
	LEFTCURLYBRACKET LEFTCURLYBRACKET {ignoreWs = true} WHITESPACE* {ignoreWs = false } ID+ {ignoreWs = true
		} WHITESPACE* {ignoreWs = false } RIGHTCURLYBRACKET RIGHTCURLYBRACKET;
// TODO: string or token based usage? handle variables like so:
// https://www.jetbrains.com/help/idea/exploring-http-syntax.html#using_request_vars

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML
// multipart-form-data) parsing
