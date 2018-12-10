package entity

import (
	"github.com/egoholic/tribune/framework/props"
	"github.com/egoholic/tribune/framework/validation"
)

type Publication struct {
	props map[string]props.Prop
}

type ValidationError struct {
	root *validation.ValidationNode
}

func (p *Publication) Props() map[string]props.Prop {
	return p.props
}

func (p *Publication) Prop(name string) props.Prop {
	return p.props[name]
}

func (p *Publication) Validate() ValidationError {
	node := validation.N("content")
	return ValidationError{&node}
}

func (ve *ValidationError) IsValid() bool {
	return ve.root.IsValid()
}
