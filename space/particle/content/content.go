package content

import (
	"github.com/egoholic/tribune/framework/prop"
	"github.com/egoholic/tribune/framework/validation"
)

type Content struct {
	props map[string]prop.Prop
}

const name = "Content"

func (c *Content) Name() string {
	return name
}

func New() Content {
	return Content{map[string]prop.Prop{}}
}

func Make(title, body string) Content {
	content := Content{map[string]prop.Prop{}}
	ct := ContentTitle{title}
	content.AssignProp(&ct)
	cb := ContentBody{body}
	content.AssignProp(&cb)

	return content
}

func (c *Content) AssignProp(p prop.Prop) error {
	c.props[p.Name()] = p
	return nil
}

func (c *Content) Props() map[string]prop.Prop {
	return c.props
}

func (c *Content) Validate() validation.ValidationResult {
	node := validation.N(c.Name())
	result := validation.Make(&node)
	return &result
}
