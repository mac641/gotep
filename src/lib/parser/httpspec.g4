grammar httpspec;

// Add additional lexer member to ignoreWs whitespace if necessary
@lexer::members {
   var ignoreWs = false
}

// List of multiple requests
file: requests EOF;
requests: (requestseparator request)+;

// Line terminators
fragment CR: '\u000D'; // \r
fragment LF: '\u000A'; // \n

NEWLINE: CR? LF | CR;
NEWLINEWITHINDENT: NEWLINE WHITESPACE;

// Chars: every unicode character except NEWLINE is supported
fragment HYPHEN: '\u002D'; // -
fragment UNDERSCORE: '\u005F'; // _

INPUT: ~('\u000D' | '\u000A'); // \r | \n
ALPHA: [\u0041-\u005A\u0061-\u007A]; // A-Z, a-z
DIGIT: [\u0030-\u0039]; // 0-9
ID: (ALPHA | DIGIT | HYPHEN | UNDERSCORE);

// Whitespaces
WHITESPACE:
	[\u0020\u0009\u000C]+ { if ignoreWs {
        l.Skip()
    } }; // space, horizontal tab (\t), form feed (\f)

// Comments
SLASH: '\u002F'; // /
HASHTAG: '\u0023'; // #
// FIXME: figure out way to ignore comments with single hashtag and double slash but not request separators
LINECOMMENT: (
		// HASHTAG INPUT NEWLINE |
		SLASH SLASH INPUT NEWLINE
	);

// Requests
fragment ASTERISK: '\u002A'; // *
fragment FULLSTOP: '\u002E'; // .

COLON: '\u003A'; // :
QUESTIONMARK: '\u003F'; // ?
LEFTSQUAREBRACKET: '\u005B'; // [
RIGHTSQUAREBRACKET: '\u005D'; // ]
METHOD:
	'\u0047' '\u0045' '\u0054' // GET
	| '\u0048' '\u0045' '\u0041' '\u0044' // HEAD
	| '\u0050' '\u004F' '\u0053' '\u0054' // POST
	| '\u0050' '\u0055' '\u0054' // PUT
	| '\u0044' '\u0045' '\u004C' '\u0045' '\u0054' '\u0045' // DELETE
	| '\u0043' '\u004F' '\u004E' '\u004E' '\u0045' '\u0043' '\u0054' // CONNECT
	| '\u0050' '\u0041' '\u0054' '\u0043' '\u0048' // PATCH
	| '\u004F' '\u0050' '\u0054' '\u0049' '\u004F' '\u004E' '\u0053' // OPTIONS
	| '\u0054' '\u0052' '\u0041' '\u0043' '\u0045'; // TRACE
ASTERISKFORM: ASTERISK;
ORIGINFORM:
	ABSOLUTEPATH (QUESTIONMARK QUERY)? (HASHTAG REQUESTFRAGMENT)?;
SEGMENT:
	~('\u002F' | '\u003F' | '\u0023'); // SLASH, QUESTIONMARK, HASHTAG
ABSOLUTEPATH: SLASH (PATHSEPARATOR SEGMENT)+;
SCHEME: '\u0068' '\u0074' '\u0074' '\u0070' '\u0073'?; // http(s)
PORT: DIGIT+;
PATHSEPARATOR: SLASH NEWLINEWITHINDENT;
HTTPVERSION:
	'\u0048' '\u0054' '\u0054' '\u0050' SLASH DIGIT+ FULLSTOP DIGIT+; // HTTP
QUERY: ~'\u0023' (NEWLINEWITHINDENT QUERY)?; // HASHTAG
REQUESTFRAGMENT:
	~'\u003F' (NEWLINEWITHINDENT REQUESTFRAGMENT)?; // QUESTIONMARK
IPV6ADDRESS:
	~('\u002F' | '\u005D'); // SLASH | RIGHTSQUAREBRACKET
IPV4ADDRESSORREGNAME:
	~('\u002F' | '\u003A' | '\u003F' | '\u0023'); // SLASH | COLON | QUESTIONMARK | HASHTAG

requestseparator:
	'###' ~(HASHTAG | NEWLINE)* NEWLINE; // TODO: why can HASHTAG HASHTAG HASHTAG not be matched?
request:
	requestline NEWLINE headers NEWLINE messagebody? responsehandler? responseref?;
requestline: (METHOD WHITESPACE)? requesttarget (
		WHITESPACE HTTPVERSION
	)?;
requesttarget: ORIGINFORM | absoluteform | ASTERISKFORM;
absoluteform: (SCHEME COLON SLASH SLASH)? hierpart (
		QUESTIONMARK QUERY
	)? (HASHTAG REQUESTFRAGMENT)?;
hierpart: authority ABSOLUTEPATH?;
authority: host (COLON PORT)?;
host:
	LEFTSQUAREBRACKET IPV6ADDRESS+ RIGHTSQUAREBRACKET
	| IPV4ADDRESSORREGNAME+;

// Headers
FIELDNAME: ~'\u003A'; // COLON

fieldvalue: INPUT NEWLINE (NEWLINEWITHINDENT fieldvalue)?;
headers: (headerfield NEWLINE)*;
headerfield:
	FIELDNAME+ COLON {ignoreWs = true } WHITESPACE* {ignoreWs = false } fieldvalue {ignoreWs = true}
		WHITESPACE* {ignoreWs = false };

// Message body
LEFTCURLYBRACKET: '\u007B'; // {
RIGHTCURLYBRACKET: '\u007D'; // }
PERCENTSIGN: '\u0025'; // %
LESSTHANSIGN: '\u003C'; // <
GREATERTHANSIGN: '\u003E'; // >
BOUNDARY: (HYPHEN HYPHEN)+ INPUT (HYPHEN HYPHEN)? NEWLINE;

filepath: INPUT NEWLINE;
inputfileref: LESSTHANSIGN WHITESPACE filepath;
messagebody: messages | multipartformdata;
messages: messageline (NEWLINE messageline)?;
messageline: ~('<' | '<>' | '###') INPUT NEWLINE | inputfileref;
// TODO: test if this works with unicode signs
multipartformdata: multipartfield multipartformdata? BOUNDARY;
multipartfield:
	BOUNDARY (headerfield NEWLINE)* NEWLINE messages?;

// Response handler
handlerscript:
	INPUT NEWLINE; // TODO: prompt error message if used anyways
responsehandler: // TODO: use parser rule when this will be implemented
	GREATERTHANSIGN WHITESPACE LEFTCURLYBRACKET PERCENTSIGN handlerscript PERCENTSIGN
		RIGHTCURLYBRACKET
	| GREATERTHANSIGN WHITESPACE filepath;

// Response reference
responseref: LESSTHANSIGN GREATERTHANSIGN WHITESPACE filepath;

// Environment variables

// ENVVARIABLE: LEFTCURLYBRACKET LEFTCURLYBRACKET {ignoreWs = true} WHITESPACE* {ignoreWs = false }
// ID+ {ignoreWs = true } WHITESPACE* {ignoreWs = false } RIGHTCURLYBRACKET RIGHTCURLYBRACKET;

// TODO: string or token based usage? handle variables like so:
// https://www.jetbrains.com/help/idea/exploring-http-syntax.html#using_request_vars

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML
// multipart-form-data) parsing
