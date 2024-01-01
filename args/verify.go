package args

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
	"unicode"
)

func verifyMethod(method string) error {
	valid := []string{
		"GET", "POST", "PATCH", "PUT", "DELETE", "HEAD", "TRACE", "CONNECT", "OPTIONS",
	}
	method = strings.TrimFunc(method, unicode.IsSpace)
	if !slices.Contains(valid, strings.ToUpper(method)) {
		return fmt.Errorf("invalid value %q received for flag %q", method, "method")
	}
	return nil
}

func verifyURL(_url string) error {
	_, err := url.ParseRequestURI(_url)
	if err != nil {
		return fmt.Errorf("invalid value %q received for flag %q", _url, "url")
	}
	return nil
}

func ParseKV(data []string, _type string) (map[string]string, error) {
	m := make(map[string]string)
	_type = fmt.Sprintf("--%s", _type)
	for _, i := range data {
		i = strings.TrimFunc(i, unicode.IsSpace)
		split := strings.Split(i, "=")
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid value %q for %s; expected values in \"k=v\" form", i, _type)
		}
		k, v := split[0], split[1]
		_, ok := m[k]
		if ok {
			return nil, fmt.Errorf("duplicate key/value pair %q received for %s", i, _type)
		}
		m[k] = v
	}
	return m, nil
}
