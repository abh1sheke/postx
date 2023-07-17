package runners

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	lhttp "github.com/abh1sheke/postx/http"
	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
)

type RequestFunc = func(id int, client *http.Client, args *parser.Args, wg *sync.WaitGroup, logger *log.Logger)

func Single(args *parser.Args, logger *log.Logger) {
	r := lhttp.InitResultList(uint(*args.Parallel))
	var method RequestFunc

	switch *args.Method {
	case "FORM":
		method = lhttp.FormRequest
	default:
		method = lhttp.DefaultRequest
	}

	startTime := time.Now()
	defer func() {
		lhttp.Data <- nil
		logging.SaveToFile(r, args.Output, logger)
		fmt.Printf(
			"took %vms for %v requests.\n",
			time.Since(startTime).Milliseconds(),
			*args.Parallel,
		)
	}()
	go r.Consumer()
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	for i := 1; i <= *args.Parallel; i++ {
		wg.Add(1)
		go method(i, client, args, wg, logger)
	}
	wg.Wait()
}