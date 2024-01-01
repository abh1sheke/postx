package args

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// Args holds all the necessary information to perfrom a HTTP request
// as specified by the user
type Args struct {
	Method, Output, URL, Proxy string
	Data, Files, Headers       map[string]string
	Multi, Include             bool
	Timeout                    time.Duration
}

// ParseKV converts a slice with data in the form of "$1=$2" into a map
func ParseKV(data []string, _type string) (map[string]string, error) {
	m := make(map[string]string)
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
