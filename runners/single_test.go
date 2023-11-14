package runners

import (
	"testing"

	"github.com/abh1sheke/postx/parser"
)

func BenchmarkSingle(t *testing.B) {
	var method string = "GET"
	var url string = "https://jsonplaceholder.typicode.com/posts"
	var output string = ""
	var headers []string = []string{}
	var args parser.Args = parser.Args{
		Method:  &method,
		URL:     &url,
		Output:  &output,
		Headers: &headers,
	}
	args.Parallel = &t.N

	Single(&args)
}
