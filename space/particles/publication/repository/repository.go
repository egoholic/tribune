package repository

import (
	"errors"
	"fmt"

	"github.com/egoholic/tribune/space/particles/publication/persistence"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func Make(c *mongo.Collection) *Repository {
	return &Repository{c}
}

func (r *Repository) Persist(p *persistence.Publication) error {
}

func (r *Repository) LatestPublished() *persistence.Publication {
	latest := &persistence.Publication{}

	return latest
}

func (r *Repository) LatestPublished10() *[]persistence.Publication {
	latestPublished10 := make([]persistence.Publication, 10)
	return &latestPublished10
}

func (r *Repository) BySlug(slug string) (*persistence.Publication, error) {
	publication := &persistence.Publication{}

	return nil, errors.New(fmt.Sprintf("Publication with slug = `%s` is not found!", slug))
}
