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
	// NOTE: multi-line mode flag (?m) enables ^ and $ to match line beginnings and endings, rather than string beginnings and endings
	regSeparator = regexp.MustCompile(`(?m)^(###[^\r\n]*(\r?\n|\r))`)
	regComments  = regexp.MustCompile(`(?m)^(//|#)([^\r\n]*(\r?\n|\r))`)
	regEnv       = regexp.MustCompile(`{{[ \t\f]*[\w\-]+[ \t\f]*}}`)
	// TODO: enhance response handler regex to be able to determine, if '###' and '%}' literals have been used inside the handler block
	regResponseHandler = regexp.MustCompile(`(?m)^>[ \t\f]+([^\r\n]*(\r?\n|\r)$|{%(.|(\r?\n|\r))+%})(\r?\n|\r)`)
	regResponseRef     = regexp.MustCompile(`(?m)^<>[ \t\f]+[^\r\n]*$(\r?\n|\r)`)
	regRequestline     = regexp.MustCompile(`(?m)^((GET|HEAD|POST|PUT|DELETE|CONNECT|PATCH|OPTIONS|TRACE)[ \t\f]+)?(([\d/\[*]|http|https)[^ \t\f]*)([ \t\f]+(HTTP/\d+(\.\d+)?))?\r?\n|\r$`)
	regEmptyNewline    = regexp.MustCompile(`(?m)^\r?\n|\r$`)
	regLineEnding      = regexp.MustCompile(`\r?\n|\r`)
	regHeaders         = regexp.MustCompile(`(?m)^(?P<Fieldname>[\w\-]+):[ \t\f]*(?P<Fieldvalue>[^\r\n]+[ \t\f]*)$`)
	regInputFileRef    = regexp.MustCompile(`(?m)^<[ \t\f]+(?P<Filepath>[^\r\n]+)$`)
	regIp              = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|\[([A-Fa-f\d]{0,4}:?)*\]`)
)
