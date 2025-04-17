package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerTipoProductoRoutes(r *mux.Router, tipoProductoHandler *handlers.TipoProductoHandler) {
	r.HandleFunc("/tipo_producto", tipoProductoHandler.CreateTipoProducto).Methods("POST")
	r.HandleFunc("/tipo_producto", tipoProductoHandler.GetAllTipoProducto).Methods("GET")
	r.HandleFunc("/tipo_producto/{id}", tipoProductoHandler.GetByIdTipoProducto).Methods("GET")
	r.HandleFunc("/tipo_producto/{id}", tipoProductoHandler.UpdateTipoProducto).Methods("PUT")
	r.HandleFunc("/tipo_producto/{id}", tipoProductoHandler.DeleteTipoProducto).Methods("DELETE")
}
