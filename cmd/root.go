package cmd

import (
	"fmt"
	"time"

	"github.com/abh1sheke/postx/args"
	"github.com/spf13/cobra"
)

var method, output, url, proxy string
var files, data, headers []string
var multi, include bool
var timeout int64

var rootCmd = &cobra.Command{
	Use:   "postx",
	Short: "A fast and feature-rich alternative to cURL.",
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
		a := &args.Args{
			Method:  method,
			Output:  output,
			URL:     url,
			Data:    _data,
			Files:   _files,
			Headers: _headers,
			Include: include,
			Proxy:   proxy,
			Timeout: time.Duration(timeout),
		}
		if len(_files) > 0 {
			a.Multi = true
		}
		if err = a.Verify(); err != nil {
			return err
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&method, "method", "m", "get", "specify http request method")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "specify output file")
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "endpoint to which request is to be sent")
	rootCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "specify proxy url")
	rootCmd.Flags().StringArrayVarP(&data, "data", "d", []string{}, "form data to be sent")
	rootCmd.Flags().BoolVarP(&multi, "multipart", "M", false, "send request data as multipart/form")
	rootCmd.Flags().StringArrayVarP(&files, "file", "f", []string{}, "path of files to send")
	rootCmd.Flags().StringArrayVarP(&headers, "headers", "H", []string{}, "set request headers")
	rootCmd.Flags().Int64VarP(&timeout, "timeout", "t", 10, "request timeout (in seconds)")
	rootCmd.Flags().BoolVarP(&include, "include", "i", false, "include request headers in output")

	rootCmd.MarkFlagRequired("url")
}
