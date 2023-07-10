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
	get := parser.NewCommand("get", "Perform a GET request")
	post := parser.NewCommand("post", "Perform a POST request")
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
	// GET Args
	getUrl := get.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	getHeaders := get.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)
	// POST args
	postUrl := post.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	postHeaders := post.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)
	postData := post.String(
		"d",
		"data",
		&argparse.Options{
			Required: false,
			Help:     "JSON; POST data",
			Validate: validateData,
		},
	)

	if err := parser.Parse(os.Args); err != nil {
		return nil, err
	}

	var method string
	var url, data *string
	var header *[]string
	if get.Happened() {
		method = strings.ToUpper(get.GetName())
		url = getUrl
		header = getHeaders
	} else {
		method = strings.ToUpper(post.GetName())
		url = postUrl
		header = postHeaders
		data = postData
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
