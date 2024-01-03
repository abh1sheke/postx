package client

import (
	"net/http"

	"github.com/abh1sheke/zing/args"
)

func buildRequest(a *args.Args) (*http.Request, error) {
	h := make(http.Header)
	h.Set("User-Agent", a.Agent)
	for k, v := range a.Headers {
		h.Set(k, v)
	}
	b, err := constructBody(a.Data, h)
	if err != nil {
		return nil, err
	}
	for k, v := range a.Headers {
		h.Set(k, v)
	}
	r, err := http.NewRequest(a.Method, a.URL, b)
	r.Header = h
	if err != nil {
		return nil, err
	}
	return r, nil
}
