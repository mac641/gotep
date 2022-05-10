package lib

import "regexp"

const (
	// Terminal output colors
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

var (
	// Lib regex
	// NOTE: multi-line mode flag (?m) enables ^ and $ to match line beginnings and endings
	regComments     = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*$(\r?\n|\r))`)
	regEmptyNewline = regexp.MustCompile(`(?m)^\r?\n|\r$`)
	regEnv          = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
	regHeaders      = regexp.MustCompile(`(?m)^(?P<Fieldname>[\w\-]+):[ \t\f]*(?P<Fieldvalue>[^\r\n]+[ \t\f]*)$`)
	regInputFileRef = regexp.MustCompile(`(?m)^<[ \t\f]+(?P<Filepath>[^\r\n]+)$`)
	regIp           = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|\[([A-Fa-f\d%]{0,4}:?){0,8}\]`)
	regLineEnding   = regexp.MustCompile(`\r?\n|\r`)
	// Boundary chars:
	// https://stackoverflow.com/questions/147451/what-are-valid-characters-for-creating-a-multipart-form-boundary
	// Boundary length:
	// https://stackoverflow.com/questions/3508338/what-is-the-boundary-in-multipart-form-data
	regMultipartFormDataHeader = regexp.
					MustCompile(`(?m)^Content-Type: multipart\/form-data(; boundary=(?P<Boundary>[0-9a-zA-Z'()+_,\-.\/:=?]{1,70}))?$(\r?\n|\r)`)
	regRequestline = regexp.
			MustCompile(`(?m)^((GET|HEAD|POST|PUT|DELETE|CONNECT|PATCH|OPTIONS|TRACE)[ \t\f]+)?(([\d/\[*]|http|https)[^ \t\f]*)([ \t\f]+(HTTP/\d+(\.\d+)?))?$(\r?\n|\r)`)
	regResponseHandler = regexp.
				MustCompile(`(?m)^>[ \t\f]+([^\r\n]*(\r?\n|\r)$|{%(.|(\r?\n|\r))+%})$(\r?\n|\r)`)
	regResponseRef = regexp.MustCompile(`(?m)^<>[ \t\f]+[^\r\n]*$(\r?\n|\r)`)
	regSeparator   = regexp.MustCompile(`(?m)^###[^\r\n]*$(\r?\n|\r)`)
	regUrlScheme   = regexp.MustCompile(`(?m)^(http|https)://`)
)
