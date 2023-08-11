package cmd

import (
	"github.com/abh1sheke/postx/cmd/del"
	"github.com/abh1sheke/postx/cmd/get"
	"github.com/abh1sheke/postx/cmd/head"
	"github.com/abh1sheke/postx/cmd/post"
	"github.com/abh1sheke/postx/cmd/put"
	"github.com/spf13/cobra"
)

type Args struct {
	Output   string
	Include  string
	Parallel int
	Loop     string
	Method   string
	Data     string
	Files    string
}

var rootCmd = &cobra.Command{
	Use:   "postx",
	Short: "A fast and feature-rich alternative to cURL.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP(
		"output",
		"o",
		"",
		"specify output file",
	)
	rootCmd.PersistentFlags().BoolP(
		"include",
		"i",
		false,
		"include request headers in output",
	)
	rootCmd.PersistentFlags().BoolP(
		"time",
		"t",
		false,
		"show time taken to make requests",
	)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(head.HeadCmd)
	rootCmd.AddCommand(post.PostCmd)
	rootCmd.AddCommand(put.PutCmd)
	rootCmd.AddCommand(del.DeleteCmd)
}
