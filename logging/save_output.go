package logging

import (
	"fmt"
	"log"
	"os"

	"github.com/abh1sheke/postx/http"
	"github.com/goccy/go-json"
)

func SaveToFile(r *http.Result, outfile *string, logger *log.Logger) {
	if len(*outfile) < 1 {
		return
	}
	fmt.Println("saving output...")
	output := make([]http.Res, len(*r.List))
	for _, v := range *r.List {
		if v != nil {
			output = append(output, *v)
		}
	}

	bytes, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("could not write to outfile.")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("outfile write error: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(
		*outfile, bytes, 0644,
	)
	if err != nil {
		fmt.Println("could not write to outfile.")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("outfile write error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote output to %v.\n", *outfile)
}
