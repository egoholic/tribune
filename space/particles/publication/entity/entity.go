package entity

import (
	"time"

	. "github.com/egoholic/tribune/framework/props"
	. "github.com/egoholic/tribune/framework/validation"
	ce "github.com/egoholic/tribune/space/particles/content/entity"
	"github.com/egoholic/tribune/space/particles/publication/entity/slugs"
)

type Publication struct {
	content ce.Content
	props   map[string]Prop
}

const name = "Publication"

func (p *Publication) Name() string {
	return name
}

func New() Publication {
	content := Content{map[string]Prop{}}
	return Publication{content, map[string]Prop{}}
}

func Make(content *Content) Publication {
	publication := Publication{*content, map[string]Prop{}}
	slug := slugs.Make(content.Props()["Title"].Read())
	ps := PublicationSlug{slug}
	publication.AssignProp(&ps)
	publishedAt := time.Now()
	ppa := PublicationPublishedAt{publishedAt}
	publication.AssignProp(&ppa)

	return publication
}

func (p *Publication) AssignProp(prop Prop) error {
	p.props[prop.Name()] = prop
	return nil
}

func (p *Publication) Props() map[string]Prop {
	return p.props
}

func (p *Publication) Validate() ValidationResult {
	node := N(p.Name())
	return SimpleValidationResult{&node}
}
