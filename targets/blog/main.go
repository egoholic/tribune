package main

import (
	"net/http"

	publishing "github.com/egoholic/tribune/outerspace_adapters/http/domains/publishing/handlers"
	reading "github.com/egoholic/tribune/outerspace_adapters/http/domains/reading/handlers"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/publications", publishing.CreatePublication).Methods("POST")
	rtr.HandleFunc("/publications/{slug:[0-9a-z-]}", publishing.UpdatePublication).Methods("PATCH")
	rtr.HandleFunc("/publications", reading.RenderAllWithPagination).Methods("GET")
	rtr.HandleFunc("/publications/{slug:[0-9a-z-]}", reading.RenderOneBySlug).Methods("GET")
	rtr.HandleFunc("/", reading.RenderTheLatestPublished).Methods("GET")

	http.Handle("/", rtr)
}
