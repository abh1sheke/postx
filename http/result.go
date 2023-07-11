package http

import (
	"fmt"
	"log"
	"os"
)

type Res struct {
	Data   string
	Status string
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
	output := fmt.Sprintf("Total requests made: %v\n", len(*r.List))
	for i, v := range *r.List {
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
