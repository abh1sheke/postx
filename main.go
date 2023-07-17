package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/runners"
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

	startTime := time.Now()
	if *args.Loop {
		runners.Looped(args, startTime, logger)
	} else {
		runners.Single(args, logger)
	}
}
