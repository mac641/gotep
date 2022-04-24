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
	_, err := url.Parse(s)
	if err != nil {
		ip := regIp.FindString(s)
		ip = strings.Trim(ip, "[]")
		if ip == "" {
			return false
		}

		address := net.ParseIP(ip)
		if address == nil {
			return false
		}

		s = regIp.ReplaceAllLiteralString(s, "")
		_, err := url.Parse(s)
		return err == nil
	}
	return true
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
