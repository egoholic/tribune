package routing

import (
	"github.com/egoholic/tribune/world/domains/reading/handlers"
	"github.com/gorilla/mux"
)

func Build() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", handlers.RenderTheLatestOne).Methods("GET")
	rtr.HandleFunc("/publications", handlers.RenderAllWithPagination).Methods("GET")
	rtr.HandleFunc("/publications/{slug:[0-9a-z-]}", handlers.RenderOne).Methods("GET")

	return rtr
}
