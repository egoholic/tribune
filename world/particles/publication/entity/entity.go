package entity

import (
	"github.com/egoholic/tribune/world/particles/publication/persistence"
)

type Publication struct {
	source *persistence.Publication
}

func Build(slug, title, preview, content, rawContent, publishedAt, createdAt string) *Publication {
	source := persistence.Build(slug, title, preview, content, rawContent, publishedAt, createdAt)
	return &Publication{source}
}

func (p *Publication) Slug() string {
	return p.source.Slug
}

func (p *Publication) Title() string {
	return p.source.Title
}

func (p *Publication) Content() string {
	return p.source.Content
}

func (p *Publication) RawContent() string {
	return p.source.RawContent
}

func (p *Publication) PublishedAt() string {
	return p.source.PublishedAt
}

func (p *Publication) CreatedAt() string {
	return p.source.CreatedAt
}