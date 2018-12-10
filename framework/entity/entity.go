package entity

import (
	"github.com/egoholic/tribune/framework/props"
	"github.com/egoholic/tribune/framework/validation"
)

type Entity interface {
	validation.Validable
	props.PropIterable
}
