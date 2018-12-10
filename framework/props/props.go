package props

// N is for "Name"
type N string

// V is for "Value"
type V interface{}

type PropReader interface {
	ReadProp(N) V
}

type PropWriter interface {
	WriteProp(N, V) error
}

type PropIterable interface {
	PropIterator() PropIterator
}

type PropIterator interface {
	NextNV() (N, V)
}
