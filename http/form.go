package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func FormRequest(
	id int,
	c chan *result.Data,
	client *http.Client,
	args *parser.Args,
	wg *sync.WaitGroup,
	logger *log.Logger,
) {
	defer wg.Done()
	var request *http.Request
	var response *http.Response
	var err error

	form := url.Values{}
	for _, v := range *args.FormData {
		values := strings.Split(v, "=")
		form.Add(values[0], values[1])
	}
	request, err = http.NewRequest("POST", *args.URL, strings.NewReader(form.Encode()))

	if err != nil {
		fmt.Println("could not create http request;")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("could not create http request: %v\n", err)
		os.Exit(1)
	}

	request.Header.Set("User-Agent", "postx/0.1")

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range *args.Headers {
		values := strings.Split(v, ":")
		request.Header.Add(values[0], values[1])
	}

	response, err = client.Do(request)
	if err != nil {
		fmt.Println("could not perform http request;")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("could not perform http request: %v\n", err)
		os.Exit(1)
	}
	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("could not read response body.")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("could not perform http request: %v\n", err)
	} else {
		c <- &result.Data{Body: &body, Request: request, Response: response}
	}
}
