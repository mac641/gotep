grammar httpSpec;

// TODO: add unicode support

// Lexical structure
// Line terminators
fragment CR	: '\r'; // \r
fragment LF	: '\n'; // \n

NEWLINE				: CR? LF | CR;
NEWLINEWITHINDENT	: NEWLINE WHITESPACE+;
LINETAIL: INPUT* NEWLINE;

// Whitespaces
WHITESPACE: [ \t\f]; // space, horizontal tab (\t), form feed (\f)
OPTIONALWHITESPACE: WHITESPACE* -> skip;
REQUIREDWHITESPACE: WHITESPACE+;

// Comments
LINECOMMENT: ('#' | '//') LINETAIL -> skip;

// Request separators
REQUESTSEPARATOR: '###' LINETAIL;

// Base symbols
INPUT: ~[\r\n];
ALPHA: [\u0041-\u005A\u0061-\u007A]; // A-Z, a-z
DIGIT: [\u0030-\u0039]; // 0-9
ID: (ALPHA | DIGIT | '-' | '_')+;


// Grammar structure
// Request file
requestfile: REQUESTSEPARATOR* request requestwithseparator* REQUESTSEPARATOR* EOF;
requestwithseparator: REQUESTSEPARATOR+ request;

// Request
request: requestline NEWLINE headers NEWLINE messagebody? responsehandler? responseref?;

// Request line
requestline: (method REQUIREDWHITESPACE)? requesttarget (REQUIREDWHITESPACE httpversion)?;
method: 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'CONNECT' | 'PATCH' | 'OPTIONS' | 'TRACE';
httpversion: 'HTTP/' DIGIT+ '.' DIGIT+;

// Request target
requesttarget: originform | absoluteform | asteriskform;
originform: absolutepath ('?' query)? ('#' requestfragment)?;
absoluteform: (scheme '//')? hierpart ('?' query)? ('#' requestfragment)?;
scheme: 'http' | 'https';
hierpart: authority absolutepath?;
asteriskform: '*';

// Authority
authority: host (':' port)?;
port: DIGIT+;
host: '[' ipv6address ']' | ipv4addressorregname;
ipv6address: ~('/' | ']')+;
ipv4addressorregname: ~('/' | ':'| '?' | '#')+;

// Resource path
absolutepath: '/' | (pathseparator segment)+;
pathseparator: '/' | NEWLINEWITHINDENT;
segment: ~('/' | '?' | '#')*;

// Query and requestfragment
query: ~'#'* (NEWLINEWITHINDENT query)?;
requestfragment: ~'?'* (NEWLINEWITHINDENT requestfragment)?;

// Headers
headers: (headerfield NEWLINE)*;
headerfield: fieldname ':' OPTIONALWHITESPACE fieldvalue OPTIONALWHITESPACE;
fieldname: ~':'+;
fieldvalue: LINETAIL (NEWLINEWITHINDENT fieldvalue);

// Message body
messagebody: messages | multipartformdata;
messages: messageline (NEWLINE messageline)?;
messageline: ~('<' | '<>' | '###') LINETAIL | inputfileref;
inputfileref: '<' REQUIREDWHITESPACE filepath;
filepath: LINETAIL;

// Mulitpart-form-data
multipartformdata: multipartfield multipartformdata? boundary;
multipartfield: boundary (headerfield NEWLINE)* NEWLINE messages?;

// Source: https://stackoverflow.com/questions/147451/what-are-valid-characters-for-creating-a-multipart-form-boundary
boundary: ( DIGIT | ALPHA | '\'' | '(' | ')' | '+' | '_' | ',' | '-' | '.' | '/' | ':' | '=' | '?' )+;

// Reponse handler
responsehandler: '>' REQUIREDWHITESPACE ('{%' handlerscript '%}' | filepath);
handlerscript: INPUT+;

// Response reference
responseref: '<>' REQUIREDWHITESPACE filepath;

// Environment variables
envvariable: '{{' OPTIONALWHITESPACE ID OPTIONALWHITESPACE '}}';

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML multipart-form-data) parsing
