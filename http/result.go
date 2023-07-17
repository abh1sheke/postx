package http

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
)

type Res struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}

var Data chan *Res = make(chan *Res)

type Result struct {
	List *[]*Res
	head int
}

func (r *Result) resize() {
	capacity := 2 * len(*r.List)
	newList := make([]*Res, capacity, capacity)
	for i, v := range *r.List {
		newList[i] = v
	}
	r.List = &newList
}

func (r *Result) Add(item *Res) {
	if r.head+1 == len(*r.List) {
		r.resize()
	}
	r.head++
	(*r.List)[r.head] = item
}

func (r *Result) SaveToFile(outfile *string, logger *log.Logger) {
	if len(*outfile) < 1 {
		return
	}
	fmt.Println("saving output...")
	output := make([]Res, len(*r.List))
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

func InitResultList(size uint) *Result {
	list := make([]*Res, size, size)
	return &Result{List: &list, head: -1}
}

func Consumer(r *Result) {
	for {
		select {
		case data := <-Data:
			if data == nil {
				break
			} else {
				r.Add(data)
			}
		}
	}
}
