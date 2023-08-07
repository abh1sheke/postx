package parser

import (
	"os"

	"github.com/akamensky/argparse"
)

type Args struct {
	Method   *string
	URL      *string
	Data     *[]string
	Headers  *[]string
	Parallel *int
	Loop     *bool
	Files    *[]string
	Include  *bool
	Time     *bool
	Output   *string
}

func Build(parser *argparse.Parser) (*Args, error) {
	get, getUrl, getHeaders := Get(parser)
	head, headUrl, headHeaders := Head(parser)
	post, postUrl, postHeaders, postBody, postFiles := Post(parser)
	put, putUrl, putHeaders, putBody := Put(parser)
	del, delUrl, delHeaders, delBody := Delete(parser)
	form, formUrl, formHeaders, formBody := Form(parser)

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
	include := parser.Flag(
		"i", "include",
		&argparse.Options{
			Required: false,
			Help:     "Include response headers in the output",
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
	time := parser.Flag(
		"t",
		"time",
		&argparse.Options{
			Required: false,
			Help:     "Print time taken for performing request(s)",
		},
	)

	err := parser.Parse(os.Args)
	if err != nil {
		return nil, err
	}

	var method string
	var url *string
	var data, header, files *[]string
	if get.Happened() {
		method = "GET"
		url = getUrl
		header = getHeaders
	} else if post.Happened() {
		method = "POST"
		url = postUrl
		header = postHeaders
		data = postBody
		files = postFiles
	} else if put.Happened() {
		method = "PUT"
		url = putUrl
		header = putHeaders
		data = putBody
	} else if del.Happened() {
		method = "DELETE"
		url = delUrl
		header = delHeaders
		data = delBody
	} else if head.Happened() {
		method = "HEAD"
		url = headUrl
		header = headHeaders
	} else if form.Happened() {
		method = "FORM"
		url = formUrl
		header = formHeaders
		data = formBody
	}

	if *parallel <= 0 {
		*parallel = 1
	}

	args := Args{
		Method:   &method,
		URL:      url,
		Data:     data,
		Files:    files,
		Headers:  header,
		Parallel: parallel,
		Loop:     loop,
		Include:  include,
		Time:     time,
		Output:   output,
	}

	return &args, err
}
