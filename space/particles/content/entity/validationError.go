package entity

import (
	. "github.com/egoholic/tribune/framework/validation"
)

type ValidationError struct {
	root *ValidationNode
}

func (ve *ValidationError) IsValid() bool {
	return ve.root.IsValid()
}

func (ve *ValidationError) Nodes() ValidationNode {
	return *ve.root
}
