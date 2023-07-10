package http

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Res struct {
	Data   string
	Status string
}

type ResMutex struct {
	M      sync.Mutex
	Result *[]*Res
	head   int
}

func (r *ResMutex) resize() {
	capacity := 2 * len(*r.Result)
	newList := make([]*Res, capacity, capacity)
	for i, v := range *r.Result {
		newList[i] = v
	}
	r.Result = &newList
}

func (r *ResMutex) Add(item *Res) {
	if r.head+1 == len(*r.Result) {
		r.resize()
	}
	r.head++
	(*r.Result)[r.head] = item
}

func (r *ResMutex) SaveToFile(outfile *string, logger *log.Logger) {
	if len(*outfile) < 1 {
		return
	}
	fmt.Println("saving output...")
	output := fmt.Sprintf("Total requests made: %v\n", len(*r.Result))
	for i, v := range *r.Result {
		if v != nil {
			line := fmt.Sprintf("Request #%v:\n\t%v\n", i+1, *v)
			output += line
		}
	}
	err := os.WriteFile(
		*outfile, []byte(output), 0644,
	)
	if err != nil {
		fmt.Println("could not write to outfile.")
		fmt.Println(`check logs by running "cat $TMPDIR/postx.log".`)
		logger.Printf("outfile write error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote output to %v.\n", *outfile)
}

func InitResMutex(size uint) *ResMutex {
	resultList := make([]*Res, size, size)
	return &ResMutex{Result: &resultList, head: -1}
}
