package del

import (
	"strconv"
	"time"

	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/runners"
	"github.com/spf13/cobra"
)

var url string
var headers []string
var parallel int
var loop bool
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Perform a DELETE request",
	Run: func(c *cobra.Command, a []string) {
		method := "HEAD"
		output := c.Parent().Flags().Lookup("output").Value.String()
		include, _ := strconv.ParseBool(c.Parent().Flags().Lookup("include").Value.String())
		benchTime, _ := strconv.ParseBool(c.Parent().Flags().Lookup("time").Value.String())
		args := parser.Args{
			Output:   &output,
			Include:  &include,
			Time:     &benchTime,
			URL:      &url,
			Method:   &method,
			Parallel: &parallel,
			Loop:     &loop,
			Headers:  &headers,
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
	DeleteCmd.Flags().StringVarP(&url, "url", "u", "", "specify endpoint url")
	DeleteCmd.MarkFlagRequired("url")

	DeleteCmd.Flags().IntVarP(&parallel, "parallel", "p", 1, "specify number of requests to make in parallel")
	DeleteCmd.Flags().BoolVarP(&loop, "loop", "l", false, "loop over n requests")

	DeleteCmd.Flags().StringSliceVarP(&headers, "headers", "H", []string{}, "add request headers")
}
