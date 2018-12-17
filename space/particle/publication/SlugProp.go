package publication

import "github.com/egoholic/tribune/framework/validation"

type PublicationSlug struct{ value string }

func (p *PublicationSlug) Name() string {
	return "Slug"
}

func (p *PublicationSlug) Read() interface{} {
	return p.value
}

func (p *PublicationSlug) Write(v interface{}) validation.ValidationResult {
	p.value = v.(string)
	node := validation.N(p.Name())
	result := validation.Make(&node)
	return &result
}
