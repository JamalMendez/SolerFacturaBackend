package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/ncf"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type NCFHandler ncf.NCFDTO

func NewNCFHandler() *NCFHandler {
	return &NCFHandler{}
}

func (h *NCFHandler) CreateNCF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	ncfCreated, err := ncf.Create(h.Serie, h.Tipo, h.Secuencia)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create NCF", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ncfCreated)
}

func (h *NCFHandler) GetAllNCF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	ncfs, err := ncf.GetAll()

	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get NCFs", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(ncfs)
}

func (h *NCFHandler) GetByIdNCF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid ncf ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	ncf, err := ncf.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get ncf", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":   "NCF found",
			"id":        fmt.Sprintf("%d", ncf.ID),
			"serie":     ncf.Serie,
			"tipo":      ncf.Tipo,
			"secuencia": ncf.Secuencia,
		},
	)
}

func (h *NCFHandler) UpdateNCF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid ncf ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := ncf.Update(h.Serie, h.Tipo, h.Secuencia, uint(id)); err != nil {
		response := fmt.Sprintf(`"error": "Falied to update ncf", "details": "%s"`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "NCF updated",
		},
	)
}

func (h *NCFHandler) DeleteNCF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid ncf ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := ncf.Delete(uint(id)); err != nil {
		response := fmt.Sprintf(`"error": "Failed to delete ncf", "details": "%s"`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "NCF deleted",
		},
	)
}
