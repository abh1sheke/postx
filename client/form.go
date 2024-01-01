package client

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func constructMutlipart(files, data map[string]string) (io.Reader, string, error) {
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

func contructURLEncoded(data map[string]string) (io.Reader, int, error) {
	form := &url.Values{}
	for k, v := range data {
		form.Set(k, v)
	}
	encoded := form.Encode()
	return strings.NewReader(encoded), len(encoded), nil
}
