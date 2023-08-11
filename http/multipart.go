package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func MultipartRequest(
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

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	localWg := new(sync.WaitGroup)
	localWg.Add(2)
	go func() {
		defer localWg.Done()
		for _, v := range *args.Data {
			f := strings.Index(v, "=")
			key, value := v[0:f], v[f+1:]
			writer.WriteField(key, value)
		}
	}()

	go func() {
		defer localWg.Done()
		for _, v := range *args.Files {
			f := strings.Index(v, "=")
			key, path := v[0:f], v[f+1:]
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("could not open file: %v;\n", path)
				fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
				logger.Printf("could not open file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()

			part, err := writer.CreateFormFile(key, filepath.Base(path))
			if err != nil {
				fmt.Printf("could not process file: %v;", path)
				fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
				logger.Printf("could not process file: %v\n", err)
				os.Exit(1)
			}
			_, err = io.Copy(part, file)
			if err != nil {
				fmt.Printf("could not process file: %v;", path)
				fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
				logger.Printf("could not process file: %v\n", err)
				os.Exit(1)
			}
		}
	}()
	localWg.Wait()

	err = writer.Close()
	if err != nil {
		fmt.Println("could not close multipart writer;")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("could not close multipart writer: %v\n", err)
		os.Exit(1)
	}
	request, err = http.NewRequest(*args.Method, *args.URL, data)

	if err != nil {
		fmt.Println("could not create http request;")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("could not create http request: %v\n", err)
		os.Exit(1)
	}

	request.Header.Set("User-Agent", "postx/0.1")
	request.Header.Set("Content-type", writer.FormDataContentType())

	for _, v := range *args.Headers {
		f := strings.Index(v, "=")
		request.Header.Add(v[0:f], v[f+1:])
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
		c <- &result.Data{Body: &body, Response: response}
	}
}
