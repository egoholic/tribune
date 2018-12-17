package publication

import (
	"time"

	"github.com/egoholic/tribune/framework/prop"
	"github.com/egoholic/tribune/framework/validation"
	"github.com/egoholic/tribune/space/particle/content"
	"github.com/egoholic/tribune/space/particle/publication/slug"
)

type Publication struct {
	content content.Content
	props   map[string]prop.Prop
}

const name = "Publication"

func (p *Publication) Name() string {
	return name
}

func New() Publication {
	cnt := content.New()
	return Publication{cnt, map[string]prop.Prop{}}
}

func Make(cnt *content.Content) Publication {
	pb := Publication{*cnt, map[string]prop.Prop{}}
	slug := slug.Make(cnt.Props()["Title"].Read().(string))
	ps := PublicationSlug{slug}
	pb.AssignProp(&ps)
	publishedAt := time.Now()
	ppa := PublicationPublishedAt{publishedAt}
	pb.AssignProp(&ppa)

	return pb
}

func (p *Publication) AssignProp(prop prop.Prop) error {
	p.props[prop.Name()] = prop
	return nil
}

func (p *Publication) Props() map[string]prop.Prop {
	return p.props
}

func (p *Publication) Validate() validation.ValidationResult {
	node := validation.N(p.Name())
	result := validation.Make(&node)
	return &result
}
