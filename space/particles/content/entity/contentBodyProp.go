package entity

import (
	"github.com/egoholic/tribune/framework/validation"
)

type ContentBody struct{ value string }

func (b *ContentBody) Name() string {
	return "Body"
}

func (b *ContentBody) Read() interface{} {
	return b.value
}

func (b *ContentBody) Write(v interface{}) validation.ValidationResult {
	b.value = v.(string)
	node := validation.N(b.Name())
	result := validation.Make(&node)
	return &result
}
