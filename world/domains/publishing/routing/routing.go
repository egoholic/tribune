package routing

import (
	"github.com/egoholic/tribune/world/domains/publishing/handlers"
	"github.com/gorilla/mux"
)

func Build() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/publications", handlers.CreatePublication).Methods("POST")
	rtr.HandleFunc("/publications/{slug:[0-9a-z-]}", handlers.UpdatePublication).Methods("PATCH")

	return rtr
}
