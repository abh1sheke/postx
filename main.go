package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	resultMutex := http.InitResMutex(uint(*args.Repeat))

	startTime := time.Now()
	if *args.Loop == "true" {
		sig := make(chan os.Signal)
		signal.Notify(sig)
		go func() {
			for {
				switch <-sig {
				case syscall.SIGINT:
					fmt.Printf(
						"keyboard interrup; exiting process.\n%v %vms\n",
                        "took: ",
						time.Since(startTime).Milliseconds(),
					)
					os.Exit(1)
				}
			}
		}()
		http.Looped(args, resultMutex)
	} else {
		http.Single(args, resultMutex)
	}

	fmt.Printf(
		"took %vms to make %v requests.\n",
		time.Since(startTime).Milliseconds(),
		len(*resultMutex.Result),
	)
}
