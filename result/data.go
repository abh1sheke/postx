package result

import "net/http"

type Data struct {
	Body     *[]byte        `json:"data"`
	Request  *http.Request  `json:"request"`
	Response *http.Response `json:"response"`
}

func (r *Data) GetData() *string {
	str := string(*r.Body)
	return &str
}

func (r *Data) GetRequest() *http.Request {
	return r.Request
}

func (r *Data) GetResponse() *http.Response {
	return r.Response
}
