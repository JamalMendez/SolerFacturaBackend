package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/producto"
	"ggstudios/solerfacturabackend/db_connection"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductoHandler db_connection.Producto

func NewProductoHandler() *ProductoHandler {
	return &ProductoHandler{}
}

func (h *ProductoHandler) CreateProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	producto, err := producto.Create(h.Descripcion, h.Costo, h.CostoEnDolares, h.TPR_id)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create producto", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":          "Producto created",
			"id":               fmt.Sprintf("%d", producto.ID),
			"descripcion":      producto.Descripcion,
			"costo":            fmt.Sprintf("%d", producto.Costo),
			"costo_en_dolares": fmt.Sprintf("%d", producto.CostoEnDolares),
			"tpr_id":           fmt.Sprintf("%d", producto.TPR_id),
		},
	)
}

func (h *ProductoHandler) GetAllProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productos, err := producto.GetAll()
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get productos", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(productos)
}

func (h *ProductoHandler) GetByIdProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	productos, err := producto.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get producto", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(productos)
}

func (h *ProductoHandler) UpdateProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := producto.Update(h.Descripcion, h.Costo, h.CostoEnDolares, h.TPR_id, uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update producto", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Producto updated"})
}
func (h *ProductoHandler) DeleteProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := producto.Delete(uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete producto", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Producto deleted"})
}
