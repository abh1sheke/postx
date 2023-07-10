package http

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/abh1sheke/postx/parser"
)

func Single(args *parser.Args, mutex *ResMutex) {
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	for i := 1; i <= *args.Repeat; i++ {
		wg.Add(1)
		go makeRequest(i, client, args, mutex, wg)
	}
	wg.Wait()
}

func Looped(args *parser.Args, mutex *ResMutex) {
	wg := new(sync.WaitGroup)
	client := new(http.Client)
	iterCount := 0
	for {
		for i := 1; i <= *args.Repeat; i++ {
			wg.Add(1)
			go makeRequest(i, client, args, mutex, wg)
		}
		iterCount++
		fmt.Printf("%v iteration(s) done\n", iterCount)
		time.Sleep(1 * time.Second)
		wg.Wait()
	}
}
