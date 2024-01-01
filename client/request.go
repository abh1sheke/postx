package client

import (
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/abh1sheke/zing/args"
)

func buildRequest(a *args.Args) (*http.Request, error) {
	var body io.Reader
	var size int
	var err error
	var multipartHeader string
	if len(a.Files) > 0 || a.Multi {
		body, multipartHeader, err = constructMutlipart(a.Files, a.Data)
	} else if len(a.Files) == 0 && len(a.Data) > 0 {
		body, size, err = contructURLEncoded(a.Data)
	}
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(a.Method, a.URL, body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		switch reflect.TypeOf(body).String() {
		case "*bytes.Buffer":
			r.Header.Add("Content-Type", multipartHeader)
		case "*strings.Reader":
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Add("Content-Length", strconv.Itoa(size))
		}
	}
	for k, v := range a.Headers {
		r.Header.Add(k, v)
	}
	return r, nil
}
