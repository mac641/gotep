grammar httpspec;

requests: | request REQUESTSEPARATOR requests;

// Line terminators
fragment CR: '\u000D';
fragment LF: '\u000A';
NEWLINE: CR? LF | CR;
NEWLINEWITHINDENT: NEWLINE REQUIREDWHITESPACE;
LINETAIL: INPUT* NEWLINE;

// Chars: every unicode character except NEWLINE is supPORTed
INPUT:
	~([CRLF] | '\u000D' | '\u000A'); // = NEWLINE because negated sets are not supported
ALPHA: [A-Za-z]; // TODO: Insert unicode chars
DIGIT: [0-9]; // TODO: Insert unicode chars
ID: (ALPHA | DIGIT | '-' | '_')+;

// Whitespaces
WHITESPACE: [\u0020\t\f]; // space, tab, form feed
OPTIONALWHITESPACE: WHITESPACE* -> skip;
REQUIREDWHITESPACE: WHITESPACE+;

// Comments
LINECOMMENT: ('#' LINETAIL | '//' LINETAIL) -> skip; // TODO: Insert unicode chars

// Requests
REQUESTSEPARATOR: '###' LINETAIL; // TODO: Insert unicode chars
METHOD:
	'GET'
	| 'HEAD'
	| 'POST'
	| 'PUT'
	| 'DELETE'
	| 'CONNECT'
	| 'PATCH'
	| 'OPTIONS'
	| 'TRACE'; // TODO: Insert unicode chars
ASTERISKFORM: '*'; // TODO: Insert unicode chars
ORIGINFORM:
	ABSOLUTEPATH ('?' QUERY)? ('#' REQUESTFRAGMENT)?; // TODO: Insert unicode chars
SEGMENT: ~('/' | '?' | '#'); // TODO: Insert unicode chars
ABSOLUTEPATH:
	'/' (PATHSEPARATOR SEGMENT)+; // TODO: Insert unicode chars
SCHEME: 'http' | 'https'; // TODO: Insert unicode chars
PORT: DIGIT+;
PATHSEPARATOR:
	'/' NEWLINEWITHINDENT; // TODO: Insert unicode chars
HTTPVERSION:
	'HTTP/' DIGIT+ '.' DIGIT+; // TODO: Insert unicode chars
QUERY:
	~'#' (NEWLINEWITHINDENT QUERY)?; // TODO: Insert unicode chars
REQUESTFRAGMENT:
	~'?' (NEWLINEWITHINDENT REQUESTFRAGMENT)?; // TODO: Insert unicode chars

request:
	requestline NEWLINE headers NEWLINE messagebody? RESPONSEHANDLER? responseref?;
requestline: (METHOD REQUIREDWHITESPACE)? requesttarget (
		REQUIREDWHITESPACE HTTPVERSION
	)?;
requesttarget: ORIGINFORM | absoluteform | ASTERISKFORM;
absoluteform: (SCHEME '://')? hierpart ('?' QUERY)? (
		'#' REQUESTFRAGMENT
	)?; // TODO: Insert unicode chars
hierpart: authority ABSOLUTEPATH?;
authority: host (':' PORT)?; // TODO: Insert unicode chars
host: '[' ipv6address ']' | ipv4addressorregname;
ipv6address: ~('/' | ']')+; // TODO: Insert unicode chars
ipv4addressorregname:
	~('/' | ':' | '?' | '#')+; // TODO: Insert unicode chars

// Headers
FIELDNAME: ~':'+; // TODO: Insert unicode chars

headers: (headerfield NEWLINE)*;
headerfield:
	FIELDNAME ':' OPTIONALWHITESPACE fieldvalue OPTIONALWHITESPACE; // TODO: Insert unicode chars
fieldvalue: LINETAIL (NEWLINEWITHINDENT fieldvalue)?;

// Message body
FILEPATH: LINETAIL;
BOUNDARY:
	'--'+ INPUT '--'? NEWLINE; // TODO: Insert unicode chars

messagebody: messages | multipartformdata;
messages: messageline (NEWLINE messageline)?;
messageline:
	~('<' | '<>' | '###') LINETAIL
	| inputfileref; // TODO: Insert unicode chars
inputfileref:
	'<' REQUIREDWHITESPACE FILEPATH; // TODO: Insert unicode chars
multipartformdata: multipartfield multipartformdata? BOUNDARY;
multipartfield:
	BOUNDARY (headerfield NEWLINE)* NEWLINE messages?;

// Response handler
HANDLERSCRIPT: (INPUT | NEWLINE)+ -> skip; // TODO: prompt error message if used anyways
RESPONSEHANDLER: // TODO: use parser rule when this will be implemented
	'>' REQUIREDWHITESPACE '{%' HANDLERSCRIPT '%}'
	| '>' REQUIREDWHITESPACE FILEPATH; // TODO: Insert unicode chars

// Response reference
responseref:
	'<>' REQUIREDWHITESPACE FILEPATH; // TODO: Insert unicode chars

// Environment variables
ENVVARIABLE: '{{' OPTIONALWHITESPACE ID OPTIONALWHITESPACE '}}';
// TODO: string or token based usage? // TODO: Insert unicode chars handle variables like so:
// https://www.jetbrains.com/help/idea/exploring-http-syntax.html#using_request_vars

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML
// multipart-form-data) parsing
