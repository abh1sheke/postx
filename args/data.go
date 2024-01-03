package args

import "strings"

type PostDataType uint8

const (
	Plain PostDataType = iota
	URLForm
	Multipart
	JSON
)

type PostData struct {
	Type    PostDataType
	DataMap map[string]string
	FileMap map[string]string
	Data    string
}

func (a *Args) Extract(text, json string, files, data, form []string, multipart bool) error {
	var postData *PostData
	var err error

	if multipart || len(files) > 0 || len(form) > 0 {
		if a.Method == "GET" {
			a.Method = "POST"
		}
		var _data, _files map[string]string
		data = append(data, form...)
		if _data, err = ParseKV(data, "=", "form"); err != nil {
			return err
		}
		if _files, err = ParseKV(files, "=", "files"); err != nil {
			return err
		}
		postData = &PostData{Type: Multipart, DataMap: _data, FileMap: _files}
		a.Data = postData
		return nil
	}

	if len(text) > 0 {
		if a.Method == "GET" {
			a.Method = "POST"
		}
		postData = &PostData{Type: Plain, Data: text}
		a.Data = postData
		return nil
	}

	if len(json) > 0 {
		if a.Method == "GET" {
			a.Method = "POST"
		}
		postData = &PostData{Type: JSON, Data: json}
		a.Data = postData
		return nil
	}

	if len(data) > 1 {
		if a.Method == "GET" {
			a.Method = "POST"
		}
		var strData string
		strData = strings.Join(data, "&")
		postData = &PostData{Type: URLForm, Data: strData}
		a.Data = postData
		return nil
	} else if len(data) == 1 {
		if a.Method == "GET" {
			a.Method = "POST"
		}
		postData = &PostData{Type: URLForm, Data: data[0]}
		a.Data = postData
		return nil
	}
	return nil
}
