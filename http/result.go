package http

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

func (r *Result) Consumer() {
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

func InitResultList(size uint) *Result {
	list := make([]*Res, size, size)
	return &Result{List: &list, head: -1}
}
