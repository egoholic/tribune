package entity

import (
	"time"

	"github.com/egoholic/tribune/framework/validation"
)

type PublicationPublishedAt struct{ value time.Time }

func (p *PublicationPublishedAt) Name() string {
	return "PublishedAt"
}

func (p *PublicationPublishedAt) Read() interface{} {
	return p.value
}

func (p *PublicationPublishedAt) Write(v interface{}) validation.ValidationResult {
	p.value = v.(time.Time)
	node := validation.N(p.Name())
	result := validation.Make(&node)
	return &result
}
