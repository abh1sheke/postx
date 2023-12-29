package main

import (
	"github.com/abh1sheke/postx/cmd"
	"github.com/abh1sheke/postx/print"
)

func main() {
	if err := cmd.Execute(); err != nil {
		print.Efatalf("Error: %v", err)
	}
}
