package args

import (
	"strings"
	"time"
)

type Args struct {
	Method, Output, URL, Proxy string
	Data, Files, Headers       map[string]string
	Multi, Include             bool
	Timeout                    time.Duration
}

func (a *Args) GetMethod() string {
	return strings.ToUpper(a.Method)
}

func (a *Args) Verify() error {
	if err := verifyMethod(a.Method); err != nil {
		return nil
	}
	if err := verifyURL(a.URL); err != nil {
		return nil
	}
	return nil
}
