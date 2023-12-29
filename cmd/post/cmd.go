package post

import (
	"strconv"
	"time"

	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/runners"
	"github.com/spf13/cobra"
)

var url string
var headers, body, files, numerical []string
var parallel int
var loop, form bool

var PostCmd = &cobra.Command{
	Use:   "post",
	Short: "Perform a POST request",
	Run: func(c *cobra.Command, a []string) {
		method := "POST"
		if form {
			method = "FORM"
		}
		output := c.Parent().Flags().Lookup("output").Value.String()
		include, _ := strconv.ParseBool(c.Parent().Flags().Lookup("include").Value.String())
		benchTime, _ := strconv.ParseBool(c.Parent().Flags().Lookup("time").Value.String())
		args := parser.Args{
			Output:   output,
			Include:  include,
			Time:     benchTime,
			URL:      url,
			Method:   method,
			Parallel: parallel,
			Loop:     loop,
			Headers:  headers,
			Files:    files,
			Data:     body,
		}
		if loop {
			startTime := time.Now()
			runners.Looped(&args, startTime)
		} else {
			runners.Single(&args)
		}
	},
}

func init() {
	PostCmd.Flags().StringVarP(&url, "url", "u", "", "specify endpoint url")
	PostCmd.MarkFlagRequired("url")

	PostCmd.Flags().IntVarP(&parallel, "parallel", "p", 1, "specify number of requests to make in parallel")
	PostCmd.Flags().BoolVarP(&loop, "loop", "l", false, "loop over n requests")
	PostCmd.Flags().BoolVarP(&form, "form", "F", false, "send form encoded data")

	PostCmd.Flags().StringSliceVarP(&body, "body", "b", []string{}, "key=value; specify request body values")
	PostCmd.Flags().StringSliceVarP(&numerical, "numerical", "n", []string{}, "key=value; specify request body (numerical) values")

	PostCmd.Flags().StringSliceVarP(&files, "file", "f", []string{}, "key=filepath; specify multipart request files")

	PostCmd.Flags().StringSliceVarP(&headers, "header", "H", []string{}, "key=value; add request headers")
}
