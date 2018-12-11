package entity

import (
	. "github.com/egoholic/tribune/framework/props"
	. "github.com/egoholic/tribune/framework/validation"
	"github.com/pkg/errors"
)

type Content struct {
	props map[string]Prop
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
