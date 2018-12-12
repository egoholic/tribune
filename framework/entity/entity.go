package entity

import (
	"github.com/egoholic/tribune/framework/basics"
	"github.com/egoholic/tribune/framework/props"
	"github.com/egoholic/tribune/framework/validation"
)

type Entity interface {
	basics.Named
	validation.Validable
	props.PropOwner
}
