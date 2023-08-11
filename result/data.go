package result

import (
	"net/http"
	"strings"
)

type Data struct {
	Body     *[]byte        `json:"data"`
	Response *http.Response `json:"response"`
}

func (r *Data) GetData() *string {
	str := string(*r.Body) + "\n"
	return &str
}

func (r *Data) GetResponse() *string {
	res := r.Response
	resToStr := res.Proto + " " + res.Status + "\n"
	for i, v := range res.Header {
		var val = v[0]
		if len(v) > 1 {
			val = strings.Join(v, ", ")
		}
		resToStr += i + ": " + val + "\n"
	}
	resToStr += "\n"
	return &resToStr
}
