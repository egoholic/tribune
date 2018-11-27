package main

import (
	"net/http"

	pr "github.com/egoholic/tribune/world/domains/publishing/routing"
)

func main() {
	publishingRouter := pr.Build()
	http.Handle("/", publishingRouter)
}
