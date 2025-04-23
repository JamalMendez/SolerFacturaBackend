package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerCotizacionRoutes(r *mux.Router, cotizacionHandler *handlers.CotizacionHandler) {
	r.HandleFunc("/cotizacion", cotizacionHandler.CreateCotizacion).Methods("POST")
	r.HandleFunc("/cotizacion", cotizacionHandler.GetAllCotizacion).Methods("GET")
	r.HandleFunc("/cotizacion/{id}", cotizacionHandler.GetByIdCotizacion).Methods("GET")
	r.HandleFunc("/cotizacion/{id}", cotizacionHandler.UpdateCotizacion).Methods("PUT")
	r.HandleFunc("/cotizacion/{id}", cotizacionHandler.DeleteCotizacion).Methods("DELETE")
}
