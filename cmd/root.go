package cmd

import (
	"github.com/abh1sheke/postx/args"
	"github.com/spf13/cobra"
)

var method, output, url string
var files, data []string
var multi, include, time bool

var rootCmd = &cobra.Command{
	Use:   "postx",
	Short: "A fast and feature-rich alternative to cURL.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var _data, _files map[string]string
		var err error
		if _data, err = args.ParseKV(data, "data"); err != nil {
			return err
		}
		if _files, err = args.ParseKV(files, "files"); err != nil {
			return err
		}
		a := &args.Args{
			Method: method, Output: output, URL: url, Data: _data, Files: _files, Include: include,
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
	rootCmd.Flags().StringArrayVarP(&data, "data", "d", []string{}, "form data to be sent")
	rootCmd.Flags().BoolVar(&multi, "multi", false, "send request data as multipart/form")
	rootCmd.Flags().StringArrayVarP(&files, "file", "f", []string{}, "path to files meant for sending")
	rootCmd.Flags().BoolVarP(&include, "include", "i", false, "include request headers in output")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagsRequiredTogether("method", "url")
}
