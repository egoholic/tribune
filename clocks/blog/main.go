package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	mgo "gopkg.in/mgo.v2"
)

type ExtendedContext struct {
	echo.Context
	MongoDBSession *mgo.Session
}

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(extendContextMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
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
