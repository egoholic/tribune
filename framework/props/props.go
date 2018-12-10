package props

import (
	"github.com/egoholic/tribune/framework/validation"
)

type Name string
type Value interface{}

type Prop interface {
	Name() string
	Read() interface{}
	Write(interface{}) validation.ValidationResult
}

type PropOwner interface {
	Props() map[string]Prop
	Prop(string) Prop
}

type PropIterable interface {
	PropIterator() PropIterator
}

type PropIterator interface {
	Next() Prop
}