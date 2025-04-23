package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerFacturaRoutes(r *mux.Router, facturaHandler *handlers.FacturaHandler) {
	r.HandleFunc("/factura", facturaHandler.CreateFactura).Methods("POST")
	r.HandleFunc("/factura", facturaHandler.GetAllFactura).Methods("GET")
	r.HandleFunc("/factura/{id}", facturaHandler.GetByIdFactura).Methods("GET")
	r.HandleFunc("/factura/{id}", facturaHandler.UpdateFactura).Methods("PUT")
	r.HandleFunc("/factura/{id}", facturaHandler.DeleteFactura).Methods("DELETE")
}
