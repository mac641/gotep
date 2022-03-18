grammar httpspec;

// Chars: every unicode character except line breaks are supported
INPUT: .;
ALPHA: [A-Za-z];
DIGIT: [0-9];
identifier: (ALPHA | DIGIT | '-' | '_')+;

// Line terminators
newline: 'r\n' | '\n' | '\r';
newlinewithindent: newline*;
linetail: INPUT* newline;

// Whitespaces
WHITESPACE: [ \t\f];
optionalwhitespace: WHITESPACE*;
requiredwhitespace: WHITESPACE+;

// Comments
linecomment: '#' linetail | '//' linetail;
// TODO: Put them in hidden channel because there not necessary for executing http requests

// Requests
requestseparator: '###' linetail;
request:
	requestline newline headers newline messagebody? responsehandler? responseref?;
requestline: (method requiredwhitespace)? requesttarget (
		requiredwhitespace httpversion
	)?;
method:
	'GET'
	| 'HEAD'
	| 'POST'
	| 'PUT'
	| 'DELETE'
	| 'CONNECT'
	| 'PATCH'
	| 'OPTIONS'
	| 'TRACE';
httpversion: 'HTTP/' DIGIT+ '.' DIGIT+;
requesttarget: originform | absoluteform | asteriskform;
originform: absolutepath ('?' query)? ('#' requestfragment)?;
absoluteform: (scheme '://')? hierpart ('?' query)? (
		'#' requestfragment
	)?;
scheme: 'http' | 'https';
hierpart: authority absolutepath?;
asteriskform: '*';
authority: host (':' port)?;
port: DIGIT+;
host: '[' ipv6address ']' | ipv4addressorregname;
ipv6address: ~('/' | ']')+;
ipv4addressorregname: ~('/' | ':' | '?' | '#')+;
absolutepath: '/' (pathseparator segment)+;
pathseparator: '/' newlinewithindent;
segment: ~('/' | '?' | '#');
query: ~'#' (newlinewithindent query)?;
requestfragment: ~'?' (newlinewithindent requestfragment)?;

// Headers
headers: (headerfield newline)*;
headerfield:
	fieldname ':' optionalwhitespace fieldvalue optionalwhitespace;
fieldname: ~':'+;
fieldvalue: linetail (newlinewithindent fieldvalue)?;

// Message body
messagebody: messages | multipartformdata;
messages: messageline (newline messageline)?;
messageline: ~('<' | '<>' | '###') linetail | inputfileref;
inputfileref: '<' requiredwhitespace filepath;
filepath: linetail;
multipartformdata: multipartfield multipartformdata? boundary;
multipartfield:
	boundary (headerfield newline)* newline messages?;

// Response handler
responsehandler:
	'>' requiredwhitespace '{%' handlerscript '%}'
	| '>' requiredwhitespace filepath;

// Response reference
responseref: '<>' requiredwhitespace filepath;

// Environment variables
envvariable:
	'{{' optionalwhitespace identifier optionalwhitespace '}}';

// TODO: add rule for boundary and handlerscript

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML
// multipart-form-data) parsing
