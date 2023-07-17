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

func Looped(
	args *parser.Args,
	startTime time.Time,
	logger *log.Logger,
) {
	r := InitResultList(uint(*args.Parallel))
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
				r.SaveToFile(args.Output, logger)
				Data <- nil
				os.Exit(1)
			}
		}
	}()

	go Consumer(r)
	for {
		for i := 1; i <= *args.Parallel; i++ {
			wg.Add(1)
			go makeRequest(i, client, args, wg, logger)
		}
		iterCount++
		fmt.Printf("%v iteration(s) done\n", iterCount)
		time.Sleep(1 * time.Second)
		wg.Wait()
	}
}