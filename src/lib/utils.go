package lib

import (
	"net"
	"net/url"
	"regexp"
	"strings"
)

func IsUrlValid(s string) bool {
	regIp := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|\[([A-Fa-f\d]{0,4}:?)*\]`)
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
