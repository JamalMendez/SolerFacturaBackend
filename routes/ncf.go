package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerNCFRoutes(r *mux.Router, ncfHandler *handlers.NCFHandler) {
	r.HandleFunc("/ncf", ncfHandler.CreateNCF).Methods("POST")
	r.HandleFunc("/ncf", ncfHandler.GetAllNCF).Methods("GET")
	r.HandleFunc("/ncf/{id}", ncfHandler.GetByIdNCF).Methods("GET")
	r.HandleFunc("/ncf/{id}", ncfHandler.UpdateNCF).Methods("PUT")
	r.HandleFunc("/ncf/{id}", ncfHandler.DeleteNCF).Methods("DELETE")
}
