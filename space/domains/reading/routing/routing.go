package routing

import (
	"github.com/egoholic/tribune/space/domains/reading/handlers"
	"github.com/gorilla/mux"
)

func Build() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", handlers.RenderTheLatestPublished).Methods("GET")
	rtr.HandleFunc("/publications", handlers.RenderAllWithPagination).Methods("GET")
	rtr.HandleFunc("/publications/{slug:[0-9a-z-]}", handlers.RenderOneBySlug).Methods("GET")

	return rtr
}
