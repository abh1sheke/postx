package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/abh1sheke/postx/parser"
)

func makeRequest(
	id int,
	client *http.Client,
	args *parser.Args,
	mutex *ResMutex,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var request *http.Request
	var response *http.Response
	var err error
	if *args.Method == "get" {
		request, err = http.NewRequest(*args.Method, *args.URL, nil)
	} else {
		data := strings.NewReader(*args.Data)
		request, err = http.NewRequest(*args.Method, *args.URL, data)
	}

	if err != nil {
		fmt.Println("could not create http request")
		fmt.Printf("error reason: %v\n", err)
		os.Exit(1)
	}

	for _, v := range *args.Headers {
		values := strings.Split(v, ":")
		request.Header.Add(values[0], values[1])
	}
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("could not create http request")
		fmt.Printf("error reason: %v\n", err)
		os.Exit(1)
	}
	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading response body for request id: %v\n", id)
		fmt.Printf("error reason: %v\n", err)
	} else {
		mutex.M.Lock()
		mutex.Add(&Res{Data: string(body), Status: response.Status})
		mutex.M.Unlock()
	}
}
