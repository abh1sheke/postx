package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/abh1sheke/zing/args"
	"github.com/spf13/cobra"
)

const (
	VERSION    = "0.1.0"
	USER_AGENT = "zing/" + VERSION
)

var _args *args.Args
var method, output, url, proxy, dataText, json, userAgent string
var files, data, dataForm, headers []string
var multi, include bool
var timeout int64

var rootCmd = &cobra.Command{
	Use:   "zing <url> [flags]",
	Short: "A cross-platform, fast and easy-to-use HTTP client for the command-line.",
	RunE: func(cmd *cobra.Command, a []string) error {
		if len(a) < 1 {
			return fmt.Errorf("no URL specified!")
		} else {
			url = a[0]
		}
		var _data, _files, _headers map[string]string
		var err error
		if _data, err = args.ParseKV(data, "data"); err != nil {
			return err
		}
		if _files, err = args.ParseKV(files, "files"); err != nil {
			return err
		}
		if _headers, err = args.ParseKV(headers, "headers"); err != nil {
			return err
		}
		if timeout > 60 {
			return fmt.Errorf("timeout value %q is too long", timeout)
		}
		_args = &args.Args{
			Method:  strings.ToUpper(method),
			Output:  output,
			URL:     url,
			Data:    _data,
			Files:   _files,
			Headers: _headers,
			Include: include,
			Proxy:   proxy,
			Timeout: time.Duration(timeout) * time.Second,
			Multi:   multi,
			Json:    json,
		}
		return nil
	},
}

func Execute() (*args.Args, error) {
	if err := rootCmd.Execute(); err != nil {
		return nil, err
	}
	return _args, nil
}

func init() {
	rootCmd.Version = VERSION

	rootCmd.Flags().StringVarP(&method, "method", "m", "get", "http request method")
	rootCmd.Flags().StringVarP(&userAgent, "user-agent", "A", USER_AGENT, "specify User-Agent to send")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "specify output file")
	rootCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "proxy url")
	rootCmd.Flags().StringVar(&dataText, "data-text", "", "set http POST data as text/plain")
	rootCmd.Flags().StringVarP(&json, "data-json", "j", "", "set http POST data as JSON")
	rootCmd.Flags().BoolVarP(&multi, "multipart", "M", false, "send request data as multipart/form-data")
	rootCmd.Flags().StringArrayVarP(&data, "data", "d", []string{}, "set http POST data")
	rootCmd.Flags().StringArrayVarP(&files, "file", "F", []string{}, "set MIME multipart MIME file (name=file)")
	rootCmd.Flags().StringArrayVar(&dataForm, "data-form", []string{}, "set http POST data as multipart/form-data (name=value)")
	rootCmd.Flags().StringArrayVarP(&headers, "headers", "H", []string{}, "set request headers (name=value)")
	rootCmd.Flags().Int64VarP(&timeout, "timeout", "t", 10, "set request timeout (in seconds)")
	rootCmd.Flags().BoolVarP(&include, "include", "i", false, "include request headers in output")
}
