package http

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func DefaultRequest(
	id int,
	c chan *result.Data,
	client *http.Client,
	args *parser.Args,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var request *http.Request
	var response *http.Response
	var err error
	if args.Method == "GET" || args.Method == "HEAD" {
		request, err = http.NewRequest(args.Method, args.URL, nil)
	} else {
		data := extractData(args)
		request, err = http.NewRequest(args.Method, args.URL, data)
	}

	if err != nil {
		logging.EFatalf(
			"Error: could not create http request.\nReason: %s",
			err.Error(),
		)
	}

	request.Header.Set("User-Agent", "postx/0.1")

	for _, v := range args.Headers {
		f := strings.Index(v, "=")
		request.Header.Add(v[0:f], v[f+1:])
	}
	response, err = client.Do(request)
	if err != nil {
		logging.EFatalf(
			"Error: could not perform http request.\nReason: %s",
			err.Error(),
		)
	}
	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		logging.EFatalf(
			"Error: could not read response body.\nReason: %s",
			err.Error(),
		)
	} else {
		c <- &result.Data{Body: body, Response: response}
	}
}
