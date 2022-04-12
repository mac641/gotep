lexer grammar httpSpecLexer;

// Add additional lexer member to ignoreWs whitespace if necessary
@lexer::members {
   var ignoreWs = false
}

// Define everything needed in default mode
//  Line terminators
NEWLINE				: CR? LF | CR;
NEWLINEWITHINDENT	: NEWLINE WHITESPACE+;

// Comments
LINECOMMENT : ('#' ~'#' ~'#' ~[\u000D\u000A]* NEWLINE | '//' ~[\u000D\u000A]* NEWLINE) -> skip;

// Whitespaces
WHITESPACE:
	[ \t\f]+ { if ignoreWs {
        l.Skip()
    } }
	; // space, horizontal tab (\t), form feed (\f)

// Line terminators
fragment CR	: '\r'; // \r
fragment LF	: '\n'; // \n

// Comments
fragment SLASH		: '\u002F'; // /
fragment HASHTAG	: '\u0023'; // #

// Headers
fragment FIELDVALUE	: (ID SLASH?)+ NEWLINE? (NEWLINEWITHINDENT FIELDVALUE)?;
fragment FIELDNAME	: ~('\u003A' | [\u000D\u000A\u0020\u0009\u000C])+ COLON; // COLON | [\r\n \t\f]

// Chars: every unicode character except NEWLINE is supported
fragment DIGIT	: [\u0030-\u0039]; // 0-9
fragment HEX	: [0-9A-Fa-f];
fragment ID		: [A-Za-z0-9\-_]; //ALPHA | DIGIT | HYPHEN | UNDERSCORE

// Headers
HEADERFIELD : FIELDNAME {ignoreWs = true} WHITESPACE* {ignoreWs = false} FIELDVALUE {ignoreWs = true} WHITESPACE* {ignoreWs = false };

// Environment variables
ENVVARIABLE:
	'{{' {ignoreWs = true} WHITESPACE* {ignoreWs = false } ID+ {ignoreWs = true
		} WHITESPACE* {ignoreWs = false } '}}'
	;
// TODO: string or token based usage? handle variables like so: https://www.jetbrains.com/help/idea/exploring-http-syntax.html#using_request_vars

// Requests
fragment QUESTIONMARK : '\u003F'; // ?
fragment SEGMENT:
	~('\u002F' | '\u003F' | '\u0023' | '\u0022' | [\u000D\u000A\u0020\u0009\u000C])+
	; // SLASH | QUESTIONMARK | HASHTAG | QUOTATION MARK | [\r\n \t\f]
fragment PATHSEPARATOR	: SLASH | NEWLINEWITHINDENT;
fragment ABSOLUTEPATH	: SLASH | (PATHSEPARATOR SEGMENT)+;
fragment PORT			: DIGIT DIGIT DIGIT DIGIT DIGIT;
// TODO: investigate regname and specify it further
fragment IPV4ADDRESSORREGNAME:
	DIGIT DIGIT? DIGIT? '.' DIGIT DIGIT? DIGIT? '.' DIGIT DIGIT? DIGIT? '.' DIGIT DIGIT? DIGIT?
	| (ID+ | '.')+
	; // SLASH | COLON | QUESTIONMARK | HASHTAG
fragment IPV6ADDRESS : (HEX | COLON)*;
// ((HEX | COLON)* HEX+)* COLON HEX* COLON (HEX+ (HEX | COLON)*)*; // TODO: figure out if this matches all ipv6 examples
fragment HOST				: '[' IPV6ADDRESS ']' | IPV4ADDRESSORREGNAME;
fragment QUERY				: ~'\u0023' (NEWLINEWITHINDENT QUERY)?; // HASHTAG
fragment ASTERISKFORM		: '*';
fragment ABSOLUTEFORM		: (('http' | 'https') '://')? (HOST (COLON PORT)?) ABSOLUTEPATH? ('?' QUERY+)? ( '#' REQUESTFRAGMENT)?;
fragment ORIGINFORM			: ABSOLUTEPATH ('?' QUERY+)? ('#' REQUESTFRAGMENT)?;
fragment REQUESTFRAGMENT	: ~'\u003F' (NEWLINEWITHINDENT REQUESTFRAGMENT)?; // QUESTIONMARK
fragment COLON				: '\u003A'; // :

fragment HTTPVERSION	: 'HTTP/' DIGIT+ '.' DIGIT+;
fragment METHOD			: 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'CONNECT' | 'PATCH' | 'OPTIONS' | 'TRACE';
fragment REQUESTTARGET	: ORIGINFORM | ABSOLUTEFORM | ASTERISKFORM;
REQUESTSEPARATOR		: '###' ~('\r' | '\n')* NEWLINE;
REQUESTLINE				: (METHOD WHITESPACE)? REQUESTTARGET (WHITESPACE HTTPVERSION)?;

// Everything inside of a messagebody
// MESSAGE_ENTRY : NEWLINE? NEWLINE MESSAGELINE -> more, pushMode(MESSAGE);
// mode MESSAGE;
// CLOSE_REQUESTSEPARATOR	: REQUESTSEPARATOR -> popMode;
// CLOSE_COMMENT			: LINECOMMENT -> skip, popMode;
// CLOSE_HEADERFIELD		: HEADERFIELD -> popMode;
// CLOSE_REQUESTTARGET			: REQUESTTARGET -> popMode;
MESSAGE : MESSAGELINE | MULTIPARTFORMDATA;

// References
fragment FILEPATH: (
		'.' SLASH?
		| (SLASH ~[\u000D\u000A\u0020\u0009\u000C\u002F]+)
		| (~[\u000D\u000A\u0020\u0009\u000C\u002F]+ SLASH)+
	) NEWLINE
	; // [\r\n \t\f/]
INPUTFILEREF : '<' WHITESPACE FILEPATH;
// Response handler
fragment HANDLERSCRIPT	: ID+ (NEWLINE | NEWLINEWITHINDENT);
RESPONSEHANDLER			: '>' WHITESPACE '{%' HANDLERSCRIPT '%}' | '>' WHITESPACE FILEPATH; // TODO: prompt error message if used anyways

// Response reference
RESPONSEREF : '<>' WHITESPACE FILEPATH;

// Messages
fragment MULTIPARTFORMDATA	: MULTIPARTFIELD MULTIPARTFORMDATA? BOUNDARY;
fragment MULTIPARTFIELD		: BOUNDARY (HEADERFIELD NEWLINE)* NEWLINE MESSAGE?;
fragment BOUNDARY			: '--'+ ID '--'? NEWLINE;
fragment MESSAGELINE		: ~[\u000D\u000A]+ WHITESPACE*;
// TODO: support unicode characters as part of message and use semantic predicate to validate first couple of chars
