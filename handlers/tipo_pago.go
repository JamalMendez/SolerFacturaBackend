package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/tipo_pago"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TipoPagoHandler tipo_pago.TipoPagoDTO

func NewTipoPagoHandler() *TipoPagoHandler {
	return &TipoPagoHandler{}
}

func (h *TipoPagoHandler) CreateTipoPago(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	tipoPago, err := tipo_pago.Create(h.Descripcion)

	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create tipo pago", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":     "Tipo pago created",
			"id":          fmt.Sprintf("%d", tipoPago.ID),
			"descripcion": tipoPago.Descripcion,
		},
	)
}

func (h *TipoPagoHandler) GetAllTipoPago(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tipoPagos, err := tipo_pago.GetAll()

	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get tipo pagos", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tipoPagos)
}

func (h *TipoPagoHandler) GetByIdTipoPago(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid user ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	tipoPago, err := tipo_pago.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get tipo pago", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":     "Tipo pago found",
			"id":          fmt.Sprintf("%d", tipoPago.ID),
			"descripcion": tipoPago.Descripcion,
		},
	)
}

func (h *TipoPagoHandler) UpdateTipoPago(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid user ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := tipo_pago.Update(h.Descripcion, uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update tipo pago", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Tipo pago updated"})
}

func (h *TipoPagoHandler) DeleteTipoPago(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid user ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := tipo_pago.Delete(uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete tipo pago", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Tipo pago deleted"})
}
