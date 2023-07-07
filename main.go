package main

import (
	"fmt"

	"github.com/abh1sheke/postx/parser"
	"github.com/akamensky/argparse"
)

func main() {
	argparser := argparse.NewParser("postx", "A CLI tool to help you test RESTful endpoints")
	_, err := parser.Build(argparser)
	if err != nil {
		fmt.Print(argparser.Usage(err))
	}
}
