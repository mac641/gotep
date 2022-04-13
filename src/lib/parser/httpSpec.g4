grammar httpSpec;

options {
	tokenVocab = httpSpecLexer;
}

// List of multiple requests
file					: REQUESTSEPARATOR* request requestwithseparator* REQUESTSEPARATOR* EOF;
requestwithseparator	: REQUESTSEPARATOR+ request;
request					: lines+;

// Lines
lines:
	REQUESTLINE NEWLINE
	| (HEADERFIELD NEWLINE?)+
	| INPUTFILEREF NEWLINE
	| RESPONSEREF NEWLINE
	| RESPONSEHANDLER NEWLINE
	| (MESSAGE (NEWLINE | NEWLINEWITHINDENT)?)+
	| NEWLINEWITHINDENT
	| NEWLINE
	;

// TODO: Investigate Island Grammars further, especially regarding JSON, ECMAScript (maybe HTML multipart-form-data) parsing
