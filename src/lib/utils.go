package lib

import (
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mac641/gotep/src/lib/context"
	"golang.org/x/term"
)

var (
	ctx = context.GetContext()
)

// Converts any file name or relative path to absolute path using context' pathPrefix
func ConvertToAbsolutePath(p string) (string, error) {
	if filepath.IsAbs(p) {
		return p, nil
	}

	prefix := ctx.GetPathPrefix()
	if !regexp.MustCompile(`^\.\/?$`).MatchString(prefix) && prefix != "" {
		p = path.Join(prefix, p)
	}
	absPath, err := filepath.Abs(p)

	return absPath, err
}

// Returns the current dimensions of file descriptor (os.Stdout.Fd()) (no scrollback buffer)
func GetTerminalSize() (width int, height int, err error) {
	width, height, err = term.GetSize(int(os.Stdout.Fd()))
	return width, height, err
}

func IsUrlValid(s string) bool {
	// Return false on empty string
	if s == "" {
		return false
	}

	// Check asterisk form
	if s == "*" {
		return true
	}

	// Check absolute / origin form
	_, err := url.ParseRequestURI(s) // returns true if absolute URI or absolute path
	// Check for misspelled url schemes since ParseRequestUri does accept these
	if err == nil && (RegUrlScheme.MatchString(s) || filepath.IsAbs(s)) {
		return true
	}

	// Check if relative url (not relative path) by prepending "http://"
	// Exclude misspelled urls schemes, urls containing ip addresses and relative paths
	if !(strings.Contains(s, ":/") || strings.Contains(s, "//") || RegIp.MatchString(s)) &&
		len(strings.Split(s, ".")) > 1 {
		absoluteUrl := "http://" + s
		_, err = url.ParseRequestURI(absoluteUrl)
		if err == nil {
			return true
		}
	}

	// Check if ip is valid
	ip := RegIp.FindString(s)
	ip = strings.Trim(ip, "[]")
	if ip == "" {
		return false
	}

	address := net.ParseIP(ip)
	if address == nil {
		return false
	}

	// Check if remaining path segment is valid
	s = RegIp.ReplaceAllLiteralString(s, "")
	_, err = url.ParseRequestURI(s)
	return err == nil
}

func TrimEmptyStringsFromSlice(s []string) []string {
	s = TrimLeftEmptyStringsFromSlice(s)
	s = TrimRightEmptyStringsFromSlice(s)
	return s
}

func TrimLeftEmptyStringsFromSlice(s []string) []string {
	firstNonEmptyStringIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			firstNonEmptyStringIndex++
		} else {
			break
		}
	}
	return s[firstNonEmptyStringIndex:]
}

func TrimRightEmptyStringsFromSlice(s []string) []string {
	lastNonEmptyStringIndex := len(s) - 1
	for i := lastNonEmptyStringIndex; i >= 0; i-- {
		if s[i] == "" {
			lastNonEmptyStringIndex--
		} else {
			break
		}
	}
	return s[0 : lastNonEmptyStringIndex+1]
}
