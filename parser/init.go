package parser

type Args struct {
	Method    string
	URL       string
	Data      []string
	Numerical []string
	Headers   []string
	Parallel  int
	Loop      bool
	Files     []string
	Include   bool
	Time      bool
	Output    string
}

func NewArgs() {}
