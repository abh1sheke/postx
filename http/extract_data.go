package http

import (
	"fmt"
	"strings"
)

func extractData(d *[]string) *strings.Reader {
	data := "{"
	for i, v := range *d {
		split := strings.Split(v, "=")
		data += fmt.Sprintf(`"%s":"%s"`, split[0], split[1])
		if i < len(*d)-1 {
			data += ","
		}
	}
	data += "}"
	return strings.NewReader(data)
}
