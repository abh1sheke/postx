package client

import (
	"net/http"
	"net/url"

	"github.com/abh1sheke/postx/args"
)

func newClient(a *args.Args) (*http.Client, error) {
	c := &http.Client{Timeout: a.Timeout}
	if len(a.Proxy) > 0 {
		proxyURL, err := url.Parse(a.Proxy)
		if err != nil {
			return nil, err
		}
		c.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	return c, nil
}

// Do creates a http respnse based on the user input
func Do(a *args.Args) (*http.Response, error) {
	c, err := newClient(a)
	if err != nil {
		return nil, err
	}
	r, err := buildRequest(a)
	if err != nil {
		return nil, err
	}
	return c.Do(r)
}
