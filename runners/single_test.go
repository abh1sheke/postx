package runners

import (
	"fmt"
	"os"
	"testing"

	"github.com/abh1sheke/postx/logging"
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
	logFile, logger, err := logging.InitLogging()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logFile.Close()
	args.Parallel = &t.N

	Single(&args, logger)
}

