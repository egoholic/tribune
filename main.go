package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RenderHome).Name("home")
	router.HandleFunc("/p/{publicationSlug:[a-z0-9]+}", RenderPublication).Name("publication")
	router.HandleFunc("/r/{rubricSlug:[a-z0-9]+}", RenderRubric).Name("rubric")
	router.HandleFunc("/s/", RenderSearchResults).Queries("qs", "{qs}").Name("searchResults")
	http.ListenAndServe("0.0.0.0:8000", router)
}

func RenderHome(rw http.ResponseWriter, r *http.Request) {
	session := assignMongoDBSessionTo(r)
	var buf bytes.Buffer
	//	session.DB("tribuneDevelop").C("publications").Find(nil)

	buf.WriteString(fmt.Sprintf("SESSION:%+v  |  REQUEST: %+v", *session, *r))

	rw.Write(buf.Bytes())
}

func RenderPublication(rw http.ResponseWriter, r *http.Request) {
	assignMongoDBSessionTo(r)
}

func RenderRubric(rw http.ResponseWriter, r *http.Request) {
	assignMongoDBSessionTo(r)
}

func RenderSearchResults(rw http.ResponseWriter, r *http.Request) {
	assignMongoDBSessionTo(r)

}

func assignMongoDBSessionTo(r *http.Request) *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	context.Set(r, "MongoDBSession", session)

	return session
}
