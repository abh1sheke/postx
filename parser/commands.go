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
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
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
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)

	return command, url, headers
}

func Post(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *string,
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
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)
	data := command.String(
		"d",
		"data",
		&argparse.Options{
			Required: true,
			Help:     "JSON; POST data",
			Validate: validateData,
		},
	)

	return command, url, headers, data
}

func Put(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *string,
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
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)
	data := command.String(
		"d",
		"data",
		&argparse.Options{
			Required: true,
			Help:     "JSON; PUT data",
			Validate: validateData,
		},
	)

	return command, url, headers, data
}

func Delete(parser *argparse.Parser) (
	*argparse.Command, *string, *[]string, *string,
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
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)
	data := command.String(
		"d",
		"data",
		&argparse.Options{
			Required: false,
			Help:     "JSON; PUT data",
			Validate: validateData,
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
			Help:     "key:value; Set request header",
			Validate: validateHeaders,
		},
	)
	data := command.StringList(
		"D",
		"data",
		&argparse.Options{
			Required: true,
			Help:     "field=value; Set form data fields",
			Validate: validateFormData,
		},
	)

	return command, url, headers, data
}
