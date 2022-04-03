grammar httpSpec;

options {
	tokenVocab = httpSpecLexer;
}

// List of multiple requests
file					: REQUESTSEPARATOR* request requestwithseparator* REQUESTSEPARATOR* EOF;
requestwithseparator	: REQUESTSEPARATOR+ request;

// Requests
request:
	requestline NEWLINE headers messagebody?
	; // MESSAGEBODY contains RESPONSEREF and RESPONSEHANDLER as part of MESSAGE lexer mode

// Lines
requestline	: (METHOD WHITESPACE)? REQUESTTARGET (WHITESPACE HTTPVERSION)?;
messagebody	: MESSAGES+ | INPUTFILEREF NEWLINE? | RESPONSEREF | RESPONSEHANDLER;

// Headers
headers : HEADERFIELD*;

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML multipart-form-data) parsing
