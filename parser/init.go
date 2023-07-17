package parser

import (
	"os"

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
	FormData *[]string
}

func Build(parser *argparse.Parser) (*Args, error) {
	get, getUrl, getHeaders := Get(parser)
	head, headUrl, headHeaders := Head(parser)
	post, postUrl, postHeaders, postData := Post(parser)
	put, putUrl, putHeaders, putData := Put(parser)
	del, delUrl, delHeaders, delData := Delete(parser)
	form, formUrl, formHeaders, formData := Form(parser)

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

	err := parser.Parse(os.Args)
	if err != nil {
		return nil, err
	}

	var method string
	var url, data *string
	var header, fdata *[]string
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
		data = delData
	} else if head.Happened() {
		method = "HEAD"
		url = headUrl
		header = headHeaders
	} else if form.Happened() {
		method = "FORM"
		url = formUrl
		header = formHeaders
		fdata = formData
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
		FormData: fdata,
	}

	return &args, err
}
