package repository

import (
	"errors"
	"fmt"

	"github.com/egoholic/tribune/space/particles/publication/persistence"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	collection *mgo.Collection
}

func Make(c *mgo.Collection) *Repository {
	return &Repository{c}
}

func (r *Repository) Persist(p *persistence.Publication) error {
	return r.collection.Insert(&p)
}

func (r *Repository) LatestPublished() *persistence.Publication {
	latest := &persistence.Publication{}
	r.collection.Find(nil).Sort("publishedAt").Limit(1).One(latest)

	return latest
}

func (r *Repository) LatestPublished10() *[]persistence.Publication {
	latestPublished10 := make([]persistence.Publication, 10)
	r.collection.Find(nil).Sort("publishedAt").Limit(10).All(&latestPublished10)
	return &latestPublished10
}

func (r *Repository) BySlug(slug string) (*persistence.Publication, error) {
	publication := &persistence.Publication{}
	r.collection.Find(bson.M{"slug": slug}).Limit(1).One(publication)

	return nil, errors.New(fmt.Sprintf("Publication with slug = `%s` is not found!", slug))
}

// TODO: replace with config and DI
func MongoDBSession() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

// TODO: replace with config and DI
func CollectionFrom(s *mgo.Session) *mgo.Collection {
	return s.DB("blog").C("publications")
}
