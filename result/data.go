package result

import (
	"fmt"
	"net/http"
	"strings"
)

type Data struct {
	Body     *[]byte        `json:"data"`
	Response *http.Response `json:"response"`
}

func (r *Data) GetData() *string {
	str := fmt.Sprintf("%v\n", string(*r.Body))
	return &str
}

func (r *Data) GetResponse() *string {
	res := r.Response
	resToStr := fmt.Sprintf("%s %s\n", res.Proto, res.Status)
	for i, v := range res.Header {
		var val = v[0]
		if len(v) > 1 {
			val = strings.Join(v, ",")
		}
		resToStr += fmt.Sprintf("%v: %v\n", i, val)
	}
	resToStr += "\n"
	return &resToStr
}
