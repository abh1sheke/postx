package main

import (
	"github.com/abh1sheke/postx/client"
	"github.com/abh1sheke/postx/cmd"
	"github.com/abh1sheke/postx/print"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			print.Eprintln("Error: ", err)
		}
	}()
	args, err := cmd.Execute()
	if err != nil || args == nil {
		return
	}
	res, err := client.Do(args)
	if err != nil {
		panic(err)
	}
	if err = print.Output(args.Output, args.Include, res); err != nil {
		panic(err)
	}
}
