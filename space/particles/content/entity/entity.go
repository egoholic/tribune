package entity

import (
	. "github.com/egoholic/tribune/framework/props"
	. "github.com/egoholic/tribune/framework/validation"
	"github.com/pkg/errors"
)

type Content struct {
	props map[string]Prop
}

type ValidationError struct {
	root *ValidationNode
}

type cTitle struct{ value string }

func (t *cTitle) Name() string {
	return "Title"
}

func (t *cTitle) Read() interface{} {
	return t.value
}

func (t *cTitle) Write(v interface{}) ValidationError {
	t.value = v.(string)
	node := N(t.Name())
	return ValidationError{&node}
}

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

func Make(title, body string) Content {
	content := Content{}
	ct := cTitle{title}
	content.AddProp(&ct)
	cb := cBody{body}
	content.AddProp(&cb)

	return content
}

func (c *Content) AddProp(p Prop) error {
	if c.props[p.Name()] != nil {
		c.props[p.Name()] = p
		return nil
	}

	return errors.Errorf("Property `%s` already exists!", p.Name())
}

func (c *Content) Props() map[string]Prop {
	return c.props
}

func (c *Content) Validate() ValidationError {
	node := N("content")
	return ValidationError{&node}
}

func (ve *ValidationError) IsValid() bool {
	return ve.root.IsValid()
}
