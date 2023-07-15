package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

type Args struct {
	Method   *string
	URL      *string
	Data     *string
	Headers  *[]string
	Parallel *int
	Loop     *bool
	Output   *string
}

func (a *Args) Verify(parser *argparse.Parser) error {
	if strings.ToLower(*a.Method) == "post" && *a.Data == "" {
		example := strings.Join([]string{
			"postx -m POST -u http://127:0.0.1:8000 -d",
			`"{\"id\": 1, \"hello\": \"world\"}"`,
		},
			" ",
		)
		return fmt.
			Errorf("data required while using POST.\nexample: %v\n", example)
	}
	return nil
}

func Build(parser *argparse.Parser) (*Args, error) {
	get, getUrl, getHeaders := Get(parser)
	head, headUrl, headHeaders := Head(parser)
	post, postUrl, postHeaders, postData := Post(parser)
	put, putUrl, putHeaders, putData := Put(parser)
	del, delUrl, delHeaders := Delete(parser)

	// Common args
	parallel := parser.Int(
		"p",
		"parallel",
		&argparse.Options{
			Required: false,
			Help:     "number; Perform n requests in parallel",
		},
	)
	loop := parser.Flag(
		"l", "loop",
		&argparse.Options{
			Required: false,
			Help:     "Loop request forever (with a 1s timeout)",
		},
	)
	output := parser.String(
		"o",
		"output",
		&argparse.Options{
			Required: false,
			Help:     "Specify output file",
		},
	)

	if err := parser.Parse(os.Args); err != nil {
		return nil, err
	}

	var method string
	var url, data *string
	var header *[]string
	if get.Happened() {
		method = "GET"
		url = getUrl
		header = getHeaders
	} else if post.Happened() {
		method = "POST"
		url = postUrl
		header = postHeaders
		data = postData
	} else if put.Happened() {
		method = "PUT"
		url = putUrl
		header = putHeaders
		data = putData
	} else if del.Happened() {
		method = "DELETE"
		url = delUrl
		header = delHeaders
	} else if head.Happened() {
		method = "HEAD"
		url = headUrl
		header = headHeaders
	}

	if *parallel <= 0 {
		*parallel = 1
	}

	args := Args{
		Method:   &method,
		URL:      url,
		Data:     data,
		Headers:  header,
		Parallel: parallel,
		Loop:     loop,
		Output:   output,
	}
	err := args.Verify(parser)

	return &args, err
}
