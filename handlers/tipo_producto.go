package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/tipo_producto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TipoProductoHandler tipo_producto.TipoProductoDTO

func NewTipoProductoHandler() *TipoProductoHandler {
	return &TipoProductoHandler{}
}

func (h *TipoProductoHandler) CreateTipoProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	tipoProducto, err := tipo_producto.Create(h.Descripcion)

	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create tipo producto", "details": %s}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":     "Tipo producto created",
			"id":          fmt.Sprintf("%d", tipoProducto.ID),
			"descripcion": tipoProducto.Descripcion,
		},
	)
}

func (h *TipoProductoHandler) GetAllTipoProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tipoProductos, err := tipo_producto.GetAll()

	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get tipo productos", "details": %s}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tipoProductos)
}

func (h *TipoProductoHandler) GetByIdTipoProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	tipoProducto, err := tipo_producto.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get tipo producto", "details": %s}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":     "Tipo producto found",
			"id":          fmt.Sprintf("%d", tipoProducto.ID),
			"descripcion": tipoProducto.Descripcion,
		},
	)
}

func (h *TipoProductoHandler) UpdateTipoProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	if err := tipo_producto.Update(h.Descripcion, uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update tipo producto", "details": %s}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Tipo producto updated"})
}

func (h *TipoProductoHandler) DeleteTipoProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	if err := tipo_producto.Delete(uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete tipo producto", "details": %s}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Tipo producto deleted"})
}
