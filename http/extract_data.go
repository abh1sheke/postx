package http

import (
	"strings"
)

func extractData(body *[]string) *strings.Reader {
	data := "{"
	for i, v := range *body {
		f := strings.IndexAny(v, "=")
		key := v[0:f]
		val := v[f+1:]
		data += `"` + key + `":"` + val + `"`
		if i < len(*body)-1 {
			data += ","
		}
	}
	data += "}"
	return strings.NewReader(data)
}
