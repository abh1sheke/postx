package http

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func FormRequest(
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

	form := url.Values{}
	for _, v := range args.Data {
		f := strings.Index(v, "=")
		form.Add(v[0:f], v[f+1:])
	}
	request, err = http.NewRequest("POST", args.URL, strings.NewReader(form.Encode()))
	if err != nil {
		logging.EFatalf(
			"Error: could not create http request.\nReason: %s",
			err.Error(),
		)
	}

	request.Header.Set("User-Agent", "postx/0.1")

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range args.Headers {
		values := strings.Split(v, "=")
		request.Header.Add(values[0], values[1])
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
