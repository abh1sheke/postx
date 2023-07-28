package logging

import (
	"fmt"
	"log"
	"os"
	"path"
)

func SaveToFile(r *string, outfile *string, logger *log.Logger) {
	if len(*outfile) < 1 {
		return
	}
	err := os.WriteFile(
		*outfile, []byte(*r), 0644,
	)
	if err != nil {
		fmt.Println("could not write to outfile.")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("outfile write error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("postx: wrote outfile to %v.\n", path.Base(*outfile))
}
