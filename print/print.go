package print

// A module containing wrapper functions over fmt.Fprint and its derivatives

import (
	"fmt"
	"os"
)

func Print(a ...any) {
	fmt.Fprint(os.Stdout, a...)
}

func Printf(s string, a ...any) {
	fmt.Fprintf(os.Stdout, s, a...)
}

func Println(a ...any) {
	fmt.Fprintln(os.Stdout, a...)
}

func Eprint(a ...any) {
	fmt.Fprint(os.Stderr, a...)
}

func Eprintf(s string, a ...any) {
	fmt.Fprintf(os.Stderr, s, a...)
}

func Eprintln(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

func Efatalf(s string, a ...any) {
	Eprintf(s, a...)
	os.Exit(1)
}
