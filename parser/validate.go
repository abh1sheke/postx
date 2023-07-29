package parser

import (
	"fmt"
	"net/url"
	"strings"
)

func validateUrl(args []string) error {
	arg := args[0]
	u, err := url.ParseRequestURI(arg)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return nil
	}
	return fmt.
		Errorf("invalid value: %v; expected valid URL", arg)
}

func validateFields(args []string) error {
	arg := args[0]
	split := strings.Split(arg, "=")
	err := fmt.Errorf("invalid value: %v; expected key:value\n", arg)
	if len(split) != 2 {
		return err
	}
	for _, v := range split {
		if len(v) > len(strings.Trim(v, " ")) {
			return fmt.Errorf(
				"invalid value: \"%v\"; please remove trailing whitespaces",
				v,
			)
		}
	}
	return nil
}
