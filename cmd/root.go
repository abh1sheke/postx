package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/abh1sheke/zing/args"
	"github.com/spf13/cobra"
)

const VERSION = "0.1.0"

var _args *args.Args
var method, output, url, proxy, json string
var files, data, headers []string
var multi, include bool
var timeout int64

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "A fast and feature-rich alternative to cURL.",
	// SilenceErrors: true,
	// SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, _ []string) error {
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
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "specify output file")
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "endpoint url")
	rootCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "proxy url")
	rootCmd.Flags().StringVar(&json, "json", "", "http POST data as JSON")
	rootCmd.Flags().BoolVarP(&multi, "multipart", "M", false, "send request data as multipart/form-data")
	rootCmd.Flags().StringArrayVarP(&data, "data", "d", []string{}, "set http POST data (name=value)")
	rootCmd.Flags().StringArrayVarP(&files, "file", "F", []string{}, "set MIME multipart MIME file (name=data)")
	rootCmd.Flags().StringArrayVarP(&headers, "headers", "H", []string{}, "set request headers (name=value)")
	rootCmd.Flags().Int64VarP(&timeout, "timeout", "t", 10, "set request timeout (in seconds)")
	rootCmd.Flags().BoolVarP(&include, "include", "i", false, "include request headers in output")

	rootCmd.MarkFlagRequired("url")
}
