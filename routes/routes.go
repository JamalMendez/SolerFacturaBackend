package routes

import (
	"net/http"

	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRouter() http.Handler {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r := router.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/health", healthCheck).Methods("GET")

	tipoProductoHandler := handlers.NewTipoProductoHandler()
	tipoPagoHandler := handlers.NewTipoPagoHandler()
	ncfHandler := handlers.NewNCFHandler()
	productoHandler := handlers.NewProductoHandler()
	clienteHandler := handlers.NewClienteHandler()
	facturaHandler := handlers.NewFacturaHandler()
	cotizacionHandler := handlers.NewCotizacionHandler()

	registerTipoProductoRoutes(r, tipoProductoHandler)
	registerTipoPagoRoutes(r, tipoPagoHandler)
	registerNCFRoutes(r, ncfHandler)
	registerProductoRoutes(r, productoHandler)
	registerClienteRoutes(r, clienteHandler)
	registerFacturaRoutes(r, facturaHandler)
	registerCotizacionRoutes(r, cotizacionHandler)

	return c.Handler(router)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
