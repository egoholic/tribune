package main

import (
	"net/http"

	rr "github.com/egoholic/tribune/world/domains/reading/routing"
)

func main() {
	readingRouter := rr.Build()
	http.Handle("/", readingRouter)
}
