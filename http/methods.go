package http

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/abh1sheke/postx/parser"
)

func Single(args *parser.Args, mutex *ResMutex, logger *log.Logger) {
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	log.Printf("starting %v single request(s).", *args.Repeat)
	for i := 1; i <= *args.Repeat; i++ {
		wg.Add(1)
		go makeRequest(i, client, args, mutex, wg, logger)
	}
	wg.Wait()
}

func Looped(args *parser.Args, mutex *ResMutex, logger *log.Logger) {
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	iterCount := 0
	log.Printf("looping %v request(s).", *args.Repeat)
	for {
		for i := 1; i <= *args.Repeat; i++ {
			wg.Add(1)
			go makeRequest(i, client, args, mutex, wg, logger)
		}
		iterCount++
		fmt.Printf("%v iteration(s) done\n", iterCount)
		time.Sleep(1 * time.Second)
		wg.Wait()
	}
}
