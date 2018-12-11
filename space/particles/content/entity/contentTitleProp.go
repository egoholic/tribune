package entity

import (
	. "github.com/egoholic/tribune/framework/validation"
)

type ContentTitle struct{ value string }

func (t *ContentTitle) Name() string {
	return "Title"
}

func (t *ContentTitle) Read() interface{} {
	return t.value
}

func (t *ContentTitle) Write(v interface{}) ValidationResult {
	t.value = v.(string)
	node := N(t.Name())
	return &ValidationError{&node}
}
