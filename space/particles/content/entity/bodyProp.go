package entity

import (
	. "github.com/egoholic/tribune/framework/validation"
)

type cBody struct{ value string }

func (b *cBody) Name() string {
	return "Body"
}

func (b *cBody) Read() interface{} {
	return b.value
}

func (b *cBody) Write(v interface{}) ValidationResult {
	b.value = v.(string)
	node := N(b.Name())
	return &ValidationError{&node}
}
