package entity

import (
	"github.com/egoholic/tribune/framework"
	"github.com/egoholic/tribune/space/particles/publication/persistence"
)

type Publication struct {
	persisted_attributes     AttributesSet
	not_persisted_attributes AttributesSet
}

type AttributesSet map[string]interface{}

func Wrap(source *framework.Persistence) *Publication {
	return &Publication{source}
}

func Build(slug, title, preview, content, rawContent, publishedAt, createdAt *string) *Publication {
	source := persistence.Build(slug, title, preview, content, rawContent, publishedAt, createdAt)
	return &Publication{source}
}
