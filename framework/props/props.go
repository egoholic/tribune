package props

import (
	"github.com/egoholic/tribune/framework/validation"
)

type Prop interface {
	Name() string
	Read() interface{}
	Write(interface{}) validation.ValidationResult
}

type PropOwner interface {
	Props() map[string]Prop
	AssignProp(Prop) error
}
