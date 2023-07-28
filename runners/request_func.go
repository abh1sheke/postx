package runners

import (
	"log"
	"net/http"
	"sync"

	"github.com/abh1sheke/postx/parser"
)

type RequestFunc = func(
	id int,
	client *http.Client,
	args *parser.Args,
	wg *sync.WaitGroup,
	logger *log.Logger,
)
