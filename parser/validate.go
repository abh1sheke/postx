package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

func validateUrl(args []string) error {
	arg := args[0]
	_, err := url.Parse(arg)
	if err != nil {
		return fmt.
			Errorf("invalid value: %v; expected valid URL", arg)
	}
	return nil
}

func validateMethod(args []string) error {
	arg := strings.ToLower(args[0])
	if arg == "get" || arg == "post" {
		return nil
	}
	return fmt.
		Errorf("invalid value %v; expected get | post\n", args[0])
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

func validateLoop(args []string) error {
	arg := strings.ToLower(args[0])
	_, err := strconv.ParseBool(arg)
	if err != nil {
		return fmt.
			Errorf("invalid value: %v; expected true | false\n", args[0])
	}
	return nil
}

func validateHeaders(args []string) error {
	arg := args[0]
	split := strings.Split(arg, ":")
	err := fmt.Errorf("invalid value: %v; expected key:pair\n", arg)
	if len(split) != 2 {
		return err
	}
	for _, v := range split {
		if len(v) < 1 {
			return err
		}
	}
	return nil
}
