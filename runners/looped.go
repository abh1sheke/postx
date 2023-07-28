package runners

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	lhttp "github.com/abh1sheke/postx/http"
	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func Looped(
	args *parser.Args,
	startTime time.Time,
	logger *log.Logger,
) {
	r := result.InitResultList(uint(*args.Parallel))
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	iterCount := 0

	var method RequestFunc
	switch *args.Method {
	case "FORM":
		method = lhttp.FormRequest
	default:
		method = lhttp.DefaultRequest
	}

	c := make(chan *result.Data)
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
				logging.HandleLogging(args, r, logger)
				c <- nil
				os.Exit(1)
			}
		}
	}()

	go r.Consumer(c)
	for {
		for i := 1; i <= *args.Parallel; i++ {
			wg.Add(1)
			go method(i, c, client, args, wg, logger)
		}
		iterCount++
		fmt.Printf("%v iteration(s) done\n", iterCount)
		time.Sleep(1 * time.Second)
		wg.Wait()
	}
}
