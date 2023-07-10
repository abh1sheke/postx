package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abh1sheke/postx/http"
	"github.com/abh1sheke/postx/logging"
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
	logFile, logger, err := logging.InitLogging()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logFile.Close()

	resultMutex := http.InitResMutex(uint(*args.Parallel))

	startTime := time.Now()
	if *args.Loop {
		http.Looped(args, resultMutex, startTime, logger)
	} else {
		http.Single(args, resultMutex, logger)
	}

	fmt.Printf(
		"took %vms to make %v requests.\n",
		time.Since(startTime).Milliseconds(),
		len(*resultMutex.Result),
	)
}
