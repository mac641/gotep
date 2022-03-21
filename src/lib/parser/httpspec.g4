grammar httpspec;

// Line terminators
NEWLINE: '\u000D\u000A' | '\u000A' | '\u000D'; // \r\n | \r | \n
NEWLINEWITHINDENT: NEWLINE REQUIREDWHITESPACE;
LINETAIL: INPUT* NEWLINE;

// Chars: every unicode character except NEWLINE is supPORTed
INPUT: ~NEWLINE;
ALPHA: [A-Za-z]; // TODO: Insert unicode chars
DIGIT: [0-9]; // TODO: Insert unicode chars
ID: (ALPHA | DIGIT | '-' | '_')+;

// Whitespaces
WHITESPACE: [ \t\f]; // TODO: Insert unicode chars
OPTIONALWHITESPACE: WHITESPACE* -> skip;
REQUIREDWHITESPACE: WHITESPACE+;

// Comments
LINECOMMENT: ('#' LINETAIL | '//' LINETAIL) -> skip;

// Requests
REQUESTSEPARATOR: '###' LINETAIL;
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
REQUESTTARGET: ORIGINFORM | ABSOLUTEFORM | ASTERISKFORM;
ASTERISKFORM: '*';
ORIGINFORM: ABSOLUTEPATH ('?' QUERY)? ('#' REQUESTFRAGMENT)?;
// FIXME: figure out if it is necessary to parse ip addresses regarding hierpart - if not, make them lexer rules
ABSOLUTEFORM: (SCHEME '://')? hierpart ('?' QUERY)? (
		'#' REQUESTFRAGMENT
	)?;
SEGMENT: ~('/' | '?' | '#');
ABSOLUTEPATH: '/' (PATHSEPARATOR SEGMENT)+;
SCHEME: 'http' | 'https';
PORT: DIGIT+;
PATHSEPARATOR: '/' NEWLINEWITHINDENT;
HTTPVERSION: 'HTTP/' DIGIT+ '.' DIGIT+;
QUERY: ~'#' (NEWLINEWITHINDENT QUERY)?;
REQUESTFRAGMENT: ~'?' (NEWLINEWITHINDENT REQUESTFRAGMENT)?;

request:
	requestline NEWLINE headers NEWLINE messagebody? responsehandler? responseref?;
requestline: (METHOD REQUIREDWHITESPACE)? REQUESTTARGET (
		REQUIREDWHITESPACE HTTPVERSION
	)?;
hierpart: authority ABSOLUTEPATH?;
authority: host (':' PORT)?;
host: '[' ipv6address ']' | ipv4addressorregname;
ipv6address: ~('/' | ']')+;
ipv4addressorregname: ~('/' | ':' | '?' | '#')+;

// Headers
FIELDNAME: ~':'+;

headers: (headerfield NEWLINE)*;
headerfield:
	FIELDNAME ':' OPTIONALWHITESPACE fieldvalue OPTIONALWHITESPACE;
fieldvalue: LINETAIL (NEWLINEWITHINDENT fieldvalue)?;

// Message body
FILEPATH: LINETAIL;

messagebody: messages | multipartformdata;
messages: messageline (NEWLINE messageline)?;
messageline: ~('<' | '<>' | '###') LINETAIL | inputfileref;
inputfileref: '<' REQUIREDWHITESPACE FILEPATH;
multipartformdata: multipartfield multipartformdata? boundary;
multipartfield:
	boundary (headerfield NEWLINE)* NEWLINE messages?;

// Response handler
responsehandler:
	'>' REQUIREDWHITESPACE '{%' handlerscript '%}'
	| '>' REQUIREDWHITESPACE FILEPATH;

// Response reference
responseref: '<>' REQUIREDWHITESPACE FILEPATH;

// Environment variables
ENVVARIABLE: '{{' OPTIONALWHITESPACE ID OPTIONALWHITESPACE '}}';

// TODO: add rule for boundary and handlerscript

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML
// multipart-form-data) parsing
