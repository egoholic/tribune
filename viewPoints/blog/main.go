package main

import (
	"net/http"

	pr "github.com/egoholic/tribune/world/domains/publishing/routing"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	publishingRouter := pr.Build()
	http.Handle("/", publishingRouter)
}

func dumbHandler(w http.ResponseWriter, r *http.Request) {

}

func initMongoDBSession() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
