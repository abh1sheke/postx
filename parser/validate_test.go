package parser

import (
	"testing"
)

func TestUrl(t *testing.T) {
	validUrls := []string{
		"https://example.com",
		"http://www.example.com",
		"https://subdomain.example.com",
		"https://example.com/path",
		"http://example.com?param=value",
		"https://example.com:8080",
		"http://localhost",
		"https://127.0.0.1",
		"https://[::1]",
		"ftp://example.com",
	}
	invalidUrls := []string{
		"http://example.com:abc",
		"://example.com",
		"not_a_url",
		"/foo/bar",
		"example.com",
		"http//example.com",
	}

	for i := range validUrls {
		err := validateUrl(validUrls[i : i+1])
		if err != nil {
			t.Fail()
		}
	}

	for i := range invalidUrls {
		err := validateUrl(invalidUrls[i : i+1])
		if err == nil {
			t.Fail()
		}
	}
}

func TestHeaders(t *testing.T) {
	invalidHeaders := []string{
		"key: value",
		"key:value ",
		"key: value ",
		" key:value",
		"key :value",
		" key :value",
		"key : value",
		"key :value ",
		"key : value",
		" key: value",
		" key:value ",
		" key: value ",
		" key :value ",
		" key : value",
		" key : value ",
	}

	for i := range invalidHeaders {
		err := validateHeaders(invalidHeaders[i : i+1])
		if err == nil {
			t.Fail()
		}
	}
}
