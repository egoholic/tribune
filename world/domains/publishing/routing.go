package routing

import (
	"github.com/gorilla/mux"
	"github.com/egoholic/tribune/world/domains/publishing/handlers"
)

var Router mux.Router = mux.NewRouter()
Router.HandleFunc("/", handlers.RenderTheLatestOne).Methods("GET")
Router.HandleFunc("/publications", handlers.RenderAllWithPagination).Methods("GET")
Router.HandleFunc("/publications/{slug:[0-9a-z-]}", handlers.RenderOne).Methods("GET")

