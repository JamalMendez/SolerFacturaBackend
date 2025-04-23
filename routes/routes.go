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
	productoHandler := handlers.NewProductoHandler()
	clienteHandler := handlers.NewClienteHandler()
	facturaHandler := handlers.NewFacturaHandler()

	registerTipoProductoRoutes(r, tipoProductoHandler)
	registerTipoPagoRoutes(r, tipoPagoHandler)
	registerNCFRoutes(r, ncfHandler)
	registerProductoRoutes(r, productoHandler)
	registerClienteRoutes(r, clienteHandler)
	registerFacturaRoutes(r, facturaHandler)

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
