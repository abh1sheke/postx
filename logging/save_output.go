package logging

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/abh1sheke/postx/colors"
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

func SaveOutput(args *parser.Args, r *result.Result) {
	output := ""
	if len(*r.List) > 1 {
		for _, v := range *r.List {
			output += *extractOutput(v, args.Include)
			output += "==========\n"
		}
	} else {
		output += *extractOutput((*r.List)[0], args.Include)
	}

	if len(*r.List) > 5 && len(*args.Output) < 1 {
		cwd, _ := os.Getwd()
		filename := filepath.Join(
			cwd, "postx-out-"+time.Now().Format(time.RFC3339)+".txt")
		fmt.Printf("postx: saving output to %v; (out > 100)\n", filename)
		writeFile(&output, &filename)
	} else if len(*args.Output) > 0 || *args.Loop {
		writeFile(&output, args.Output)
	} else {
		colorizedOut := colors.ColorizeOutput(&output)
		fmt.Println(*colorizedOut)
	}
}

func writeFile(r *string, outfile *string) {
	if len(*outfile) < 1 {
		return
	}
	err := os.WriteFile(
		*outfile, []byte(*r), 0644,
	)
	if err != nil {
		EFatalf("Error: could not write to outfile.\nReason: %s", err.Error())
	}
	fmt.Printf("postx: wrote outfile to %v.\n", path.Base(*outfile))
}
