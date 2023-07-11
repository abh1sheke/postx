package http

import (
	"log"
	"net/http"
	"sync"

	"github.com/abh1sheke/postx/parser"
)

func Single(args *parser.Args, logger *log.Logger) {
	r := InitResultList(uint(*args.Parallel))
	defer func() {
		Data <- nil
		r.SaveToFile(args.Output, logger)
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
