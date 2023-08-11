package main

import (
	"fmt"
	"os"

	"github.com/abh1sheke/postx/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %s", err.Error()))
	}
}
