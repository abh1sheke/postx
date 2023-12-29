package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func MultipartRequest(
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

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	localWg := new(sync.WaitGroup)
	localWg.Add(2)
	go func() {
		defer localWg.Done()
		for _, v := range args.Data {
			f := strings.Index(v, "=")
			key, value := v[0:f], v[f+1:]
			writer.WriteField(key, value)
		}
	}()

	go func() {
		defer localWg.Done()
		for _, v := range args.Files {
			f := strings.Index(v, "=")
			key, path := v[0:f], v[f+1:]
			file, err := os.Open(path)
			if err != nil {
				logging.EFatalf(
					"Error: could not open file '%s'.\nReason: %s",
					path,
					err.Error(),
				)
			}
			defer file.Close()

			part, err := writer.CreateFormFile(key, filepath.Base(path))
			if err != nil {
				logging.EFatalf(
					"Error: could not create form file.\nReason: %s",
					err.Error(),
				)
			}
			_, err = io.Copy(part, file)
			if err != nil {
				logging.EFatalf(
					"Error: failed to copy form file.\nReason: %s",
					err.Error(),
				)
			}
		}
	}()
	localWg.Wait()

	err = writer.Close()
	if err != nil {
		logging.EFatalf(
			"Error: failed to close file writer.\nReason: %s",
			err.Error(),
		)
	}
	request, err = http.NewRequest(args.Method, args.URL, data)

	if err != nil {
		logging.EFatalf(
			"Error: could not create http request.\nReason: %s",
			err.Error(),
		)
	}

	request.Header.Set("User-Agent", "postx/0.1")
	request.Header.Set("Content-type", writer.FormDataContentType())

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
