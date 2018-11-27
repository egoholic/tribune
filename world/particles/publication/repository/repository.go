package repository

import (
	"errors"
	"fmt"

	"github.com/egoholic/tribune/world/particles/publication/persistence"
	mgo "gopkg.in/mgo.v2"
)

type Repository struct {
	collection *mgo.Collection
}

func Make(c *mgo.Collection) *Repository {
	return &Repository{c}
}

func (r *Repository) Latest10() [10]*persistence.Publication {
	latest := [10]*persistence.Publication{}
	data, ok := r.collection.([]*persistence.Publication)
	if ok {
		for i, p := range data {
			latest[9-i] = p
		}
	}

	return latest
}

func (r *Repository) BySlug(slug string) (*persistence.Publication, error) {
	data, ok := ds.DataStore.([]*persistence.Publication)
	if ok {
		for i := 0; i < len(data); i++ {
			p := dataStoreMock[i]
			if p.Slug == slug {
				return p, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("Publication with slug = `%s` is not found!", slug))
}

func initMongoDBSession() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
