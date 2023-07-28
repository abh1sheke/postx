package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

func extractOutput(d *result.Data, i *bool) *string {
	var output string = ""
	if *i {
		output += *d.GetResponse()
	}
	output += *d.GetData()
	return &output
}

func HandleLogging(args *parser.Args, r *result.Result, logger *log.Logger) {
	output := ""
	if len(*r.List) > 1 {
		for _, v := range *r.List {
			output += *extractOutput(v, args.Include)
			output += fmt.Sprintf("==========\n")
		}
	} else {
		output += *extractOutput((*r.List)[0], args.Include)
	}

	if len(*r.List) > 100 && len(*args.Output) < 1 {
		cwd, _ := os.Getwd()
		filename := filepath.Join(
			cwd, fmt.Sprintf("postx-out-%v.txt", time.Now().Format(time.RFC3339)))
		fmt.Printf("postx: saving output to %v; (out > 100)\n", filename)
		SaveToFile(&output, &filename, logger)
	} else if len(*args.Output) > 0 {
		SaveToFile(&output, args.Output, logger)
	} else {
		fmt.Println(output)
	}
}
