package lib

import (
	"regexp"
)

var (
	// Lib regex
	// NOTE: multi-line mode flag (?m) enables ^ and $ to match line beginnings and endings
	RegComments     = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*$(\r?\n|\r))`)
	RegEmptyNewLine = regexp.MustCompile(`(?m)^\r?\n|\r$`)
	RegEnv          = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
	RegHeaders      = regexp.MustCompile(`(?m)^(?P<Fieldname>[\w\-]+):[ \t\f]*(?P<Fieldvalue>[^\r\n]+[ \t\f]*)$`)
	RegInputFileRef = regexp.MustCompile(`(?m)^<[ \t\f]+(?P<Filepath>[^\r\n]+)$`)
	RegIp           = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|\[([A-Fa-f\d%]{0,4}:?){0,8}\]`)
	RegLineEnding   = regexp.MustCompile(`\r?\n|\r`)
	// Boundary chars:
	// https://stackoverflow.com/questions/147451/what-are-valid-characters-for-creating-a-multipart-form-boundary
	// Boundary length:
	// https://stackoverflow.com/questions/3508338/what-is-the-boundary-in-multipart-form-data
	RegMultipartFormDataHeader = regexp.
					MustCompile(`(?m)^Content-Type: multipart\/form-data(; boundary=(?P<Boundary>[0-9a-zA-Z'()+_,\-.\/:=?]{1,70}))?$(\r?\n|\r)`)
	// Non Ascii regex source: https://stackoverflow.com/questions/2124010/grep-regex-to-match-non-ascii-characters
	RegNonAscii    = regexp.MustCompile(`[^\x00-\x7F]`)
	RegRequestLine = regexp.
			MustCompile(`(?m)^((GET|HEAD|POST|PUT|DELETE|CONNECT|PATCH|OPTIONS|TRACE)[ \t\f]+)?(([\d/\[*]|http|https)\S*)((\r?\n|\r)^\s+[\w\-\/?#~:.%+]+(\r?\n|\r)?)*([ \t\f]+(HTTP/\d+(\.\d+)?))?$(\r?\n|\r)`)
	RegRequestLineQuery       = regexp.MustCompile(`(?m)\?[^\s?]+$`)
	RegRequestLineFragment    = regexp.MustCompile(`(?m)#[^\s#]+$`)
	RegRequestLinePathSegment = regexp.MustCompile(`(?m)\/[^\/][^\s.]+$`)
	RegResponseHandler        = regexp.
					MustCompile(`(?m)^>[ \t\f]+([^\r\n]*(\r?\n|\r)$|{%(.|(\r?\n|\r))+%})$(\r?\n|\r)`)
	RegResponseRef = regexp.MustCompile(`(?m)^<>[ \t\f]+[^\r\n]*$(\r?\n|\r)`)
	RegSeparator   = regexp.MustCompile(`(?m)^###[^\r\n]*$(\r?\n|\r)`)
	RegUrlScheme   = regexp.MustCompile(`(?m)^(http|https)://`)
)
