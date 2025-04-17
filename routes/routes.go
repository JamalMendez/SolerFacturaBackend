package routes

import (
	"net/http"

	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func InitRouter(tipoProductoHandler *handlers.TipoProductoHandler) *mux.Router {
	router := mux.NewRouter()

	r := router.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/health", healthCheck).Methods("GET")

	registerTipoProductoRoutes(r, tipoProductoHandler)

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
