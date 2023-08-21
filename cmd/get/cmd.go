package get

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abh1sheke/postx/logging"
	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/runners"
	"github.com/spf13/cobra"
)

var url string
var headers []string
var parallel int
var loop bool
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Perform a GET request",
	Run: func(c *cobra.Command, a []string) {
		f, logger, err := logging.InitLogging()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		defer f.Close()

		method := "GET"
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
			runners.Looped(&args, startTime, logger)
		} else {
			runners.Single(&args, logger)
		}
	},
}

func init() {
	GetCmd.Flags().StringVarP(&url, "url", "u", "", "specify endpoint url")
	GetCmd.MarkFlagRequired("url")

	GetCmd.Flags().IntVarP(&parallel, "parallel", "p", 1, "specify number of requests to make in parallel")
	GetCmd.Flags().BoolVarP(&loop, "loop", "l", false, "loop over n requests")

	GetCmd.Flags().StringSliceVarP(&headers, "headers", "H", []string{}, "add request headers")
}
