package lib_test

import (
	"testing"

	"github.com/mac641/gotep/src/lib"
)

func TestIsUrlValid(t *testing.T) {
	var (
		validUrls = []string{
			"https://192.168.178.2/index/https",
			"192.168.178.2/index/https",
			"http://www.example.com/test1",
			"[1111:2222:abcd:4444:5a6c:7777:9c9f:6789]/index",
			"/api/get",
			"*",
			"example.com/test?id=1",
			"example.com:80",
		}
		invalidUrls = []string{
			"httphallt",
			"",
			"a/relative/path",
			"htt://example.com",
			"htt//example.com",
		}
	)

	for i := range validUrls {
		valid := lib.IsUrlValid(validUrls[i])
		if !valid {
			t.Errorf("expected %s to be valid url.", validUrls[i])
		}
	}

	for i := range invalidUrls {
		invalid := lib.IsUrlValid(invalidUrls[i])
		if invalid {
			t.Errorf("expected %s to be invalid url.", invalidUrls[i])
		}
	}
}
