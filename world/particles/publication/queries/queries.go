package queries

import (
	"errors"
	"fmt"

	"github.com/egoholic/tribune/world/particles/publication/entity"
)

// TODO: replace with MongoDB
var dataStoreMock = []*entity.Publication{
	entity.Build("slug-1", "Title-1", "Preview-1", "Content-1", "RawContent-1", "published-at-1", "created-at-1"),
	entity.Build("slug-2", "Title-2", "Preview-2", "Content-2", "RawContent-2", "published-at-2", "created-at-2"),
	entity.Build("slug-3", "Title-3", "Preview-3", "Content-3", "RawContent-3", "published-at-3", "created-at-3"),
	entity.Build("slug-4", "Title-4", "Preview-4", "Content-4", "RawContent-4", "published-at-4", "created-at-4"),
	entity.Build("slug-5", "Title-5", "Preview-5", "Content-5", "RawContent-5", "published-at-5", "created-at-5"),
	entity.Build("slug-6", "Title-6", "Preview-6", "Content-6", "RawContent-6", "published-at-6", "created-at-6"),
	entity.Build("slug-7", "Title-7", "Preview-7", "Content-7", "RawContent-7", "published-at-7", "created-at-7"),
	entity.Build("slug-8", "Title-8", "Preview-8", "Content-8", "RawContent-8", "published-at-8", "created-at-8"),
	entity.Build("slug-9", "Title-9", "Preview-9", "Content-9", "RawContent-9", "published-at-9", "created-at-9"),
	entity.Build("slug-10", "Title-10", "Preview-10", "Content-10", "RawContent-10", "published-at-10", "created-at-10"),
}

func Latest10() [10]*entity.Publication {
	latest := [10]*entity.Publication{}
	for i, p := range dataStoreMock {
		latest[10-i] = p
	}
	return latest
}

func BySlug(slug string) (*entity.Publication, error) {
	for i := 0; i < len(dataStoreMock); i++ {
		p := dataStoreMock[i]
		if p.Slug() == slug {
			return p, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Publication with slug = `%s` is not found!", slug))
}
