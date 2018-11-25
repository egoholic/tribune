package main

import (
	"net/http"

	"github.com/egoholic/tribune/mongoSession"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

type ExtendedContext struct {
	echo.Context
	MongoDBSession *mgo.Session
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", dumbHandler)
	http.Handle("/", mongoSession.WithMongoSession(router))
}

func dumbHandler(w http.ResponseWriter, r *http.Request) {

}

func extendContextMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newContext := &ExtendedContext{c, initMongoDBSession()}
		return h(newContext)
	}
}

func initMongoDBSession() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
