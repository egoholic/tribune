package handlers

import (
	"net/http"

	pr "github.com/egoholic/tribune/space/particles/publication/repository"
	mgo "gopkg.in/mgo.v2"
)

func RenderTheLatestPublished(w http.ResponseWriter, r *http.Request) {
	repo := pr.Make(collectionFrom(mongoSession()))
	repo.LatestPublished()
}

func RenderAllWithPagination(w http.ResponseWriter, r *http.Request) {
	repo := pr.Make(collectionFrom(mongoSession()))
	repo.LatestPublished10()
}

func RenderOneBySlug(w http.ResponseWriter, r *http.Request) {
	repo := pr.Make(collectionFrom(mongoSession()))
	repo.BySlug("slug") // TODO: add context and url/query string params to handler signatures.
}

func mongoSession() *mgo.Session {
	// TODO: move to config
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func collectionFrom(s *mgo.Session) *mgo.Collection {
	// TODO: move to config
	return s.DB("tribune").C("publications")
}
