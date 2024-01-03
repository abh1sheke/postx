package client

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/abh1sheke/zing/args"
)

func constructBody(data *args.PostData, h http.Header) (io.Reader, error) {
	var body io.Reader
	var err error
	if data == nil {
		return nil, nil
	}
	switch data.Type {
	case args.Plain:
		body = strings.NewReader(data.Data)
		h.Set("Content-type", "text/plain")
	case args.JSON:
		body = strings.NewReader(data.Data)
		h.Set("Content-type", "application/json")
	case args.URLForm:
		body = strings.NewReader(data.Data)
		if err != nil {
			return nil, err
		}
		h.Set("Content-type", "application/x-www-form-urlencoded")
	case args.Multipart:
		var header string
		body, header, err = constructMutlipart(data.FileMap, data.DataMap)
		if err != nil {
			return nil, err
		}
		h.Set("Content-type", header)
	}
	return body, nil
}

func constructMutlipart(files, data map[string]string) (*bytes.Buffer, string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for param, file := range files {
		b, err := os.ReadFile(file)
		if err != nil {
			return nil, "", err
		}
		part, err := writer.CreateFormFile(param, filepath.Base(file))
		if err != nil {
			return nil, "", err
		}
		if _, err = part.Write(b); err != nil {
			return nil, "", err
		}
	}
	for k, v := range data {
		if err := writer.WriteField(k, v); err != nil {
			return nil, "", err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}

// func contructURLEncoded(data map[string]string) (*strings.Reader, error) {
// 	var encoded string
// 	form := &url.Values{}
// 	for k, v := range data {
// 		form.Set(k, v)
// 	}
// 	encoded := form.Encode()
// 	return strings.NewReader(encoded), nil
// }
