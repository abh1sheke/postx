package http

import "sync"

type Res struct {
	Data   string
	Status string
}

type ResMutex struct {
	M      sync.Mutex
	Result *[]*Res
	length uint
	head   int
}

func (r *ResMutex) resize() {
    capacity := 2 * r.length
    newList := make([]*Res, capacity, capacity)
    for i, v := range *r.Result {
        newList[i] = v
    }
    r.Result = &newList
}

func (r *ResMutex) Add(item *Res) {
    if r.head + 1 == int(r.length) {

    }
   (*r.Result)[r.head+1] = item 
}

func InitResMutex(size uint) *ResMutex {
	resultList := make([]*Res, size, size)
	return &ResMutex{Result: &resultList, length: size, head: -1}
}
