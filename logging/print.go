package logging

import (
	"fmt"
	"os"
)

func EPrint(a ...any) {
	fmt.Fprint(os.Stderr, a...)
}

func EPrintf(s string, a ...any) {
	fmt.Fprintf(os.Stderr, s, a...)
}

func EPrintln(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

func EFatalf(s string, a ...any) {
	EPrintf(s, a...)
	os.Exit(1)
}
