package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abh1sheke/postx/http"
	"github.com/abh1sheke/postx/parser"
	"github.com/akamensky/argparse"
)

func main() {
	argparser := argparse.NewParser(
		"postx",
		"A CLI tool to help you test RESTful endpoints",
	)
	args, err := parser.Build(argparser)
	if err != nil {
		fmt.Print(argparser.Usage(err))
		os.Exit(1)
	}

	list := make([]http.Res, 0)
	resultMutex := http.ResMutex{Result: list}

	startTime := time.Now().UnixMilli()
	if *args.Loop == "true" {
		http.Looped(args, &resultMutex)
	} else {
		http.Single(args, &resultMutex)
	}
	endTime := time.Now().UnixMilli()

	fmt.Printf(
		"took %vms to make %v requests.\n",
		endTime-startTime,
		len(resultMutex.Result),
	)
}
