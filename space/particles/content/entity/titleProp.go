package entity

import (
	. "github.com/egoholic/tribune/framework/validation"
)

type cTitle struct{ value string }

func (t *cTitle) Name() string {
	return "Title"
}

func (t *cTitle) Read() interface{} {
	return t.value
}

func (t *cTitle) Write(v interface{}) ValidationResult {
	t.value = v.(string)
	node := N(t.Name())
	return &ValidationError{&node}
}
