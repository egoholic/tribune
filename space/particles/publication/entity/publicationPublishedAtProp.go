package entity

import "time"

type PublicationPublishedAt struct{ value time.Time }

func (p *PublicationPublishedAt) Name() string {
	return "PublishedAt"
}

func (p *PublicationPublishedAt) Read() interface{} {
	return p.value
}

func (p *PublicationPublishedAt) Write(v interface{}) ValidationResult {
	p.value = v.(string)
	node := N(p.Name())
	return &ValidationError{&node}
}
