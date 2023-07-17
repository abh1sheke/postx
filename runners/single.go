package runners

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	lochttp "github.com/abh1sheke/postx/http"
	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
)

func Single(args *parser.Args, logger *log.Logger) {
	r := lochttp.InitResultList(uint(*args.Parallel))
	startTime := time.Now()
	defer func() {
		lochttp.Data <- nil
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
		go lochttp.DefaultRequest(i, client, args, wg, logger)
	}
	wg.Wait()
}
