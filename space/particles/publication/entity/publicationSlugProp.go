package entity

type PublicationSlug struct{ value string }

func (p *PublicationSlug) Name() string {
	return "Slug"
}

func (p *PublicationSlug) Read() interface{} {
	return p.value
}

func (p *PublicationSlug) Write(v interface{}) ValidationResult {
	p.value = v.(string)
	node := N(p.Name())
	return &ValidationError{&node}
}
