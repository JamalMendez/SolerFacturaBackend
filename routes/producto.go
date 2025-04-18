package routes

import (
	"ggstudios/solerfacturabackend/handlers"

	"github.com/gorilla/mux"
)

func registerProductoRoutes(r *mux.Router, tipoPagoHandler *handlers.ProductoHandler) {
	r.HandleFunc("/producto", tipoPagoHandler.CreateProducto).Methods("POST")
	r.HandleFunc("/producto", tipoPagoHandler.GetAllProducto).Methods("GET")
	r.HandleFunc("/producto/{id}", tipoPagoHandler.GetByIdProducto).Methods("GET")
	r.HandleFunc("/producto/{id}", tipoPagoHandler.UpdateProducto).Methods("PUT")
	r.HandleFunc("/producto/{id}", tipoPagoHandler.DeleteProducto).Methods("DELETE")
}
