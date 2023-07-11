package http

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/abh1sheke/postx/parser"
)

func Single(args *parser.Args, logger *log.Logger) {
	r := InitResultList(uint(*args.Parallel))
	startTime := time.Now()
	defer func() {
		Data <- nil
		r.SaveToFile(args.Output, logger)
		fmt.Printf(
			"took %vms for %v requests.\n",
			time.Since(startTime).Milliseconds(),
			*args.Parallel,
		)
	}()
	go Consumer(r)
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	for i := 1; i <= *args.Parallel; i++ {
		wg.Add(1)
		go makeRequest(i, client, args, wg, logger)
	}
	wg.Wait()
}
