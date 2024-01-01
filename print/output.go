package print

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func includeStatus(buf *bytes.Buffer, res *http.Response) error {
	status := strings.Join(
		[]string{"\033[1m" + res.Proto, res.Request.Method, res.Status, "\033[0m", "\n"},
		" ",
	)
	if _, err := buf.WriteString(status); err != nil {
		return err
	}
	return nil
}

func includeHeaders(buf *bytes.Buffer, res *http.Response) error {
	for k, v := range res.Header {
		val := "\033[1m" + k + "\033[0m" + ": " + strings.Join(v, "; ")
		if _, err := buf.WriteString(val); err != nil {
			return err
		}
		if err := buf.WriteByte(byte('\n')); err != nil {
			return err
		}
	}
	return nil
}

func Output(outfile string, include bool, res *http.Response) error {
	var out io.Writer
	buf := new(bytes.Buffer)
	if include {
		if err := includeStatus(buf, res); err != nil {
			return err
		}
		if err := includeHeaders(buf, res); err != nil {
			return err
		}
		if err := buf.WriteByte(byte('\n')); err != nil {
			return err
		}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if _, err = buf.Write(body); err != nil {
		return err
	}
	if len(outfile) == 0 {
		out = os.Stdout
	} else {
		var err error
		if out, err = os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			return err
		}
	}
	out.Write(buf.Bytes())
	return nil
}
