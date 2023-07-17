package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/goccy/go-json"
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

func validateData(args []string) error {
	arg := []byte(args[0])
	values := make(map[string]interface{})
	err := json.Unmarshal(arg, &values)
	if err != nil {
		return fmt.
			Errorf("invalid value: %v; expected JSON value\n", args[0])
	}
	return nil
}

func validateHeaders(args []string) error {
	arg := args[0]
	split := strings.Split(arg, ":")
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

func validateFormData(args []string) error {
	arg := args[0]
	split := strings.Split(arg, "=")
	err := fmt.Errorf("invalid value: %v; expected field=value\n", arg)
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

