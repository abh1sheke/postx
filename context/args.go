package context

import (
	"strings"
)

type Args struct {
	Method, Output, URL string
	Data, Files         map[string]string
	Include             bool
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
