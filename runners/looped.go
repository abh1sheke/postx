package runners

import (
	"fmt"
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

func Looped(args *parser.Args, startTime time.Time) {
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
					"\nkeyboard interrup; exiting process.\n%v",
					fmt.Sprintf("looped %v request(s) for %v iterations. (total = %v)\n",
						*args.Parallel, iterCount, *args.Parallel*iterCount),
				)
				fmt.Println()
				logging.SaveOutput(args, r)
				c <- nil
				os.Exit(1)
			}
		}
	}()

	go r.Consumer(c)
	for {
		for i := 1; i <= *args.Parallel; i++ {
			wg.Add(1)
			go method(i, c, client, args, wg)
		}
		time.Sleep(1 * time.Second)
		wg.Wait()
		iterCount++
		fmt.Printf("\r%v iteration(s) done", iterCount)
	}
}
