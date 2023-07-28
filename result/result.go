package result

type Result struct {
	List *[]*Data
	head int
}

func (r *Result) resize() {
	capacity := 2 * len(*r.List)
	newList := make([]*Data, capacity, capacity)
	for i, v := range *r.List {
		newList[i] = v
	}
	r.List = &newList
}

func (r *Result) Add(item *Data) {
	if r.head+1 == len(*r.List) {
		r.resize()
	}
	r.head++
	(*r.List)[r.head] = item
}

func (r *Result) Consumer() {
	for {
		select {
		case data := <-DataChan:
			if data == nil {
				break
			} else {
				r.Add(data)
			}
		}
	}
}

func InitResultList(size uint) *Result {
	list := make([]*Data, size, size)
	return &Result{List: &list, head: -1}
}
