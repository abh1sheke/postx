package parser

import "github.com/akamensky/argparse"

func Get(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string,
) {
	command := parser.
		NewCommand("get", "Perform a GET request")
	url := command.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	headers := command.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set request header",
			Validate: validateFields,
		},
	)

	return command, url, headers
}

func Head(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string,
) {
	command := parser.
		NewCommand("head", "Perform a HEAD request")
	url := command.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	headers := command.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set request header",
			Validate: validateFields,
		},
	)

	return command, url, headers
}

func Post(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *[]string, *[]string,
) {
	command := parser.
		NewCommand("post", "Perform a POST request")
	url := command.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	headers := command.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set request header",
			Validate: validateFields,
		},
	)
	data := command.StringList(
		"d",
		"data",
		&argparse.Options{
			Required: true,
			Help:     "field=value; Set POST data body",
			Validate: validateFields,
		},
	)
	files := command.StringList(
		"f",
		"file",
		&argparse.Options{
			Required: false,
			Help:     "field=filepath; Set POST file data",
			Validate: validateFields,
		},
	)

	return command, url, headers, data, files
}

func Put(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *[]string,
) {
	command := parser.
		NewCommand("put", "Perform a PUT request")
	url := command.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	headers := command.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set request header",
			Validate: validateFields,
		},
	)
	data := command.StringList(
		"d",
		"data",
		&argparse.Options{
			Required: true,
			Help:     "field=value; Set PUT data body",
			Validate: validateFields,
		},
	)

	return command, url, headers, data
}

func Delete(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *[]string,
) {
	command := parser.
		NewCommand("delete", "Perform a DELETE request")
	url := command.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	headers := command.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set request header",
			Validate: validateFields,
		},
	)
	data := command.StringList(
		"d",
		"data",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set DELETE body data",
			Validate: validateFields,
		},
	)

	return command, url, headers, data
}

func Form(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *[]string,
) {
	command := parser.
		NewCommand("form", "Submit a HTML form")
	url := command.String(
		"u",
		"url",
		&argparse.Options{
			Required: true,
			Help:     "URL of endpoint",
			Validate: validateUrl,
		},
	)
	headers := command.StringList(
		"H",
		"header",
		&argparse.Options{
			Required: false,
			Help:     "field=value; Set request header",
			Validate: validateFields,
		},
	)
	data := command.StringList(
		"D",
		"data",
		&argparse.Options{
			Required: true,
			Help:     "field=value; Set form data",
			Validate: validateFields,
		},
	)

	return command, url, headers, data
}
