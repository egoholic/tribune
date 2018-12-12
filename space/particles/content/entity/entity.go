package entity

import (
	. "github.com/egoholic/tribune/framework/props"
	. "github.com/egoholic/tribune/framework/validation"
)

type Content struct {
	props map[string]Prop
}

const name = "Content"

func (c *Content) Name() string {
	return name
}

func New() Content {
	return Content{map[string]Prop{}}
}

func Make(title, body string) Content {
	content := Content{map[string]Prop{}}
	ct := ContentTitle{title}
	content.AssignProp(&ct)
	cb := ContentBody{body}
	content.AssignProp(&cb)

	return content
}

func (c *Content) AssignProp(p Prop) error {
	c.props[p.Name()] = p
	return nil
}

func (c *Content) Props() map[string]Prop {
	return c.props
}

func (c *Content) Validate() ValidationResult {
	node := N(c.Name())
	return SimpleValidationResult{&node}
}
