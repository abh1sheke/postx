package runners

import (
	"log"
	"net/http"
	"sync"

	"github.com/abh1sheke/postx/parser"
	"github.com/abh1sheke/postx/result"
)

type RequestFunc = func(
	id int,
  c chan *result.Data,
	client *http.Client,
	args *parser.Args,
	wg *sync.WaitGroup,
	logger *log.Logger,
)
