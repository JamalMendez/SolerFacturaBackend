package routes

import (
	"net/http"

	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	r := router.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/health", healthCheck).Methods("GET")

	tipoProductoHandler := handlers.NewTipoProductoHandler()
	tipoPagoHandler := handlers.NewTipoPagoHandler()
	ncfHandler := handlers.NewNCFHandler()

	registerTipoProductoRoutes(r, tipoProductoHandler)
	registerTipoPagoRoutes(r, tipoPagoHandler)
	registerNCFRoutes(r, ncfHandler)

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
