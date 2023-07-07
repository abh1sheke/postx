package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

type Args struct {
	Method  *string
	URL     *string
	Data    *string
	Headers *[]string
	Repeat  *int
	Loop    *string
}

func (a *Args) Verify(parser *argparse.Parser) error {
	if strings.ToLower(*a.Method) == "post" && *a.Data == "" {
		example := strings.Join([]string{
			"postx -m POST -u http://127:0.0.1:8000 -d",
			"\"{\\\"id\\\": 1, \\\"hello\\\": \\\"world\\\"}\""},
			" ",
		)
		return fmt.
			Errorf("data required while using POST.\nexample: %v\n", example)
	}
	return nil
}

func Build(parser *argparse.Parser) (*Args, error) {
	method := parser.String(
		"m",
		"method",
		&argparse.Options{
			Required: true,
			Help:     "GET | POST; HTTP method",
			Validate: validateMethod,
		},
	)
	url := parser.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	data := parser.String(
		"d",
		"data",
		&argparse.Options{
			Required: false,
			Help:     "JSON; POST data",
			Validate: validateData,
		},
	)
	headers := parser.StringList(
		"H",
		"headers",
		&argparse.Options{
			Required: false,
			Help:     "key:value; Set request headers",
			Validate: validateHeaders,
		},
	)
	repeat := parser.Int(
		"r",
		"repeat",
		&argparse.Options{
			Required: false,
			Help:     "number; Repeat the request n number of times",
		},
	)
	loop := parser.String(
		"l", "loop",
		&argparse.Options{
			Required: false,
			Help:     "true | false; Perform n repitions forever (with a 1s timeout)",
			Validate: validateLoop,
		},
	)
	if err := parser.Parse(os.Args); err != nil {
		return nil, err
	}

	args := Args{
		Method:  method,
		URL:     url,
		Data:    data,
		Headers: headers,
		Repeat:  repeat,
		Loop:    loop,
	}
	err := args.Verify(parser)

	return &args, err
}
