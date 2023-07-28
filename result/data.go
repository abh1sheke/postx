package result

import "net/http"

type Data struct {
	Body     *[]byte        `json:"data"`
	Response *http.Response `json:"response"`
}

func (r *Data) GetData() *string {
	str := string(*r.Body)
	return &str
}

func (r *Data) GetResponse() *http.Response {
	return r.Response
}
