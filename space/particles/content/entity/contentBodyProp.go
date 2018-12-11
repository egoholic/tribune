package entity

import (
	. "github.com/egoholic/tribune/framework/validation"
)

type ContentBody struct{ value string }

func (b *ContentBody) Name() string {
	return "Body"
}

func (b *ContentBody) Read() interface{} {
	return b.value
}

func (b *ContentBody) Write(v interface{}) ValidationResult {
	b.value = v.(string)
	node := N(b.Name())
	return &ValidationError{&node}
}
