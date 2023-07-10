package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/abh1sheke/postx/parser"
)

func Single(args *parser.Args, mutex *ResMutex, logger *log.Logger) {
	defer mutex.SaveToFile(args.Output, logger)
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	for i := 1; i <= *args.Parallel; i++ {
		wg.Add(1)
		go makeRequest(i, client, args, mutex, wg, logger)
	}
	wg.Wait()
}

func Looped(
	args *parser.Args,
	mutex *ResMutex,
	startTime time.Time,
	logger *log.Logger,
) {
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	iterCount := 0

	sig := make(chan os.Signal)
	signal.Notify(sig)
	go func() {
		for {
			switch <-sig {
			case syscall.SIGINT:
				fmt.Printf(
					"\nkeyboard interrup; exiting process.\n%v %vms\n",
					"took: ",
					time.Since(startTime).Milliseconds(),
				)
				mutex.SaveToFile(args.Output, logger)
				os.Exit(1)
			}
		}
	}()

	for {
		for i := 1; i <= *args.Parallel; i++ {
			wg.Add(1)
			go makeRequest(i, client, args, mutex, wg, logger)
		}
		iterCount++
		fmt.Printf("%v iteration(s) done\n", iterCount)
		time.Sleep(1 * time.Second)
		wg.Wait()
	}
}
