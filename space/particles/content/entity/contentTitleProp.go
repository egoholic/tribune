package entity

import "github.com/egoholic/tribune/framework/validation"

type ContentTitle struct{ value string }

func (t *ContentTitle) Name() string {
	return "Title"
}

func (t *ContentTitle) Read() interface{} {
	return t.value
}

func (t *ContentTitle) Write(v interface{}) validation.ValidationResult {
	t.value = v.(string)
	node := validation.N(t.Name())
	result := validation.Make(&node)
	return &result
}
