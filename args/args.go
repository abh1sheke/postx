package args

import (
	"time"
)

// Args holds all the necessary information to perfrom a HTTP request
// as specified by the user
type Args struct {
	Method, Output, URL, Proxy string
	Data, Files, Headers       map[string]string
	Multi, Include             bool
	Timeout                    time.Duration
}

func (a *Args) Verify() error {
	if err := verifyMethod(a.Method); err != nil {
		return err
	}
	if err := verifyURL(a.URL); err != nil {
		return err
	}
	return nil
}
