package main

import (
	"github.com/abh1sheke/postx/client"
	"github.com/abh1sheke/postx/cmd"
	"github.com/abh1sheke/postx/print"
)

func main() {
	args, err := cmd.Execute()
	if err != nil {
		return
	}
	_, err = client.Do(args)
	if err != nil {
		print.Eprintln("Error: ", err)
    return
	}
}
