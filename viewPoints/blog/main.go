package main

import (
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", dumbHandler)
	http.Handle("/", router)
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
