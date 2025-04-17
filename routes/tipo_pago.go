package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerTipoPagoRoutes(r *mux.Router, tipoPagoHandler *handlers.TipoPagoHandler) {
	r.HandleFunc("/tipo_pago", tipoPagoHandler.CreateTipoPago).Methods("POST")
	r.HandleFunc("/tipo_pago", tipoPagoHandler.GetAllTipoPago).Methods("GET")
	r.HandleFunc("/tipo_pago/{id}", tipoPagoHandler.GetByIdTipoPago).Methods("GET")
	r.HandleFunc("/tipo_pago/{id}", tipoPagoHandler.UpdateTipoPago).Methods("PUT")
	r.HandleFunc("/tipo_pago/{id}", tipoPagoHandler.DeleteTipoPago).Methods("DELETE")
}
