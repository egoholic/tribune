package association

type Pairwise interface {
	Left() interface{}
	Right() interface{}
}

type Association interface {
	Add(Pairwise)
	Name() string
	LMap(interface{}) interface{}
	RMap(interface{}) interface{}
}
