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
	Method, Output, URL, Proxy, Agent string
	Headers                           map[string]string
	Multi, Include                    bool
	Timeout                           time.Duration
	Data                              *PostData
}

func trimWhitespace(s ...*string) {
	for _, i := range s {
		strings.TrimFunc(*i, unicode.IsSpace)
	}
}

// ParseKV converts a slice with data in the form of "$1$delim$2" into a map
func ParseKV(data []string, delim, _type string) (map[string]string, error) {
	m := make(map[string]string)
	for _, i := range data {
		split := strings.SplitN(i, delim, 2)
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid value %q for %s; expected values in \"k%sv\" form", i, _type, delim)
		}
		k, v := split[0], split[1]
		trimWhitespace(&k, &v)
		_, ok := m[k]
		if ok {
			fmt.Printf("warning: duplicate key/value pair %q received for %s\n", i, _type)
		}
		m[k] = v
	}
	return m, nil
}
