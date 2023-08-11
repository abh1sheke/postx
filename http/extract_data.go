package http

import (
	"strings"
	"sync"

	"github.com/abh1sheke/postx/parser"
)

func extractStringData(body *[]string, c chan *string, wg *sync.WaitGroup) {
	defer wg.Done()
	if body != nil {
		for _, v := range *body {
			f := strings.IndexAny(v, "=")
			key := v[0:f]
			val := v[f+1:]
			data := `"` + key + `":"` + val + `"`
			c <- &data
		}
	}
}

func extractNumericalData(body *[]string, c chan *string, wg *sync.WaitGroup) {
	defer wg.Done()
	if body != nil {
		for _, v := range *body {
			f := strings.IndexAny(v, "=")
			key := v[0:f]
			val := v[f+1:]
			data := `"` + key + `":` + val
			c <- &data
		}
	}
}

func extractData(args *parser.Args) *strings.Reader {
  wg := new(sync.WaitGroup)
	data := make(chan *string)
	var max int
	if args.Data != nil && args.Numerical != nil {
		max = len(*args.Data) + len(*args.Numerical)
	} else if args.Numerical == nil {
		max = len(*args.Data)
	} else {
		max = len(*args.Numerical)
	}
	body := "{"
	count := 0
	wg.Add(1)
	go func() {
		for {
			select {
			case item := <-data:
				body += *item
				count++
				if count == max {
					body += "}"
					wg.Done()
					break
				} else if count < max {
					body += ","
				}
			}
		}
	}()
	wg.Add(2)
	go extractStringData(args.Data, data, wg)
	go extractNumericalData(args.Numerical, data, wg)
	wg.Wait()
	return strings.NewReader(body)
}
