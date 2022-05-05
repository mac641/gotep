package lib

import (
	"net"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

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
	if err == nil && (regUrlScheme.MatchString(s) || filepath.IsAbs(s)) {
		return true
	}

	// Check if relative url (not relative path) by prepending "http://"
	// Exclude misspelled urls schemes, urls containing ip addresses and relative paths
	if !(strings.Contains(s, ":/") || strings.Contains(s, "//") || regIp.MatchString(s)) &&
		len(strings.Split(s, ".")) > 1 {
		absoluteUrl := "http://" + s
		_, err = url.ParseRequestURI(absoluteUrl)
		if err == nil {
			return true
		}
	}

	// Check if ip is valid
	ip := regIp.FindString(s)
	ip = strings.Trim(ip, "[]")
	if ip == "" {
		return false
	}

	address := net.ParseIP(ip)
	if address == nil {
		return false
	}

	// Check if remaining path segment is valid
	s = regIp.ReplaceAllLiteralString(s, "")
	_, err = url.ParseRequestURI(s)
	return err == nil
}

func ConvertToAbsolutePath(p string, prefix string) string {
	if filepath.IsAbs(p) {
		return p
	}

	if prefix != "." && prefix != "" {
		p = path.Join(prefix, p)
	}
	absPath, err := filepath.Abs(p)
	cobra.CheckErr(err)

	return absPath
}
