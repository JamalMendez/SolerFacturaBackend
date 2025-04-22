package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerClienteRoutes(r *mux.Router, clienteHandler *handlers.ClienteHandler) {
	r.HandleFunc("/cliente", clienteHandler.CreateCliente).Methods("POST")
	r.HandleFunc("/cliente", clienteHandler.GetAllCliente).Methods("GET")
	r.HandleFunc("/cliente/{id}", clienteHandler.GetByIdCliente).Methods("GET")
	r.HandleFunc("/cliente/{id}", clienteHandler.UpdateCliente).Methods("PUT")
	r.HandleFunc("/cliente/{id}", clienteHandler.DeleteCliente).Methods("DELETE")
}
