package main

import (
	"github.com/abh1sheke/postx/cmd"
	"github.com/abh1sheke/postx/logging"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logging.EPrintf("Error: %s", err.Error())
	}
}
