package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/cliente"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ClienteHandler cliente.ClienteDTO

func NewClienteHandler() *ClienteHandler {
	return &ClienteHandler{}
}

func (h *ClienteHandler) CreateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	clienteCreated, err := cliente.Create(h.RNC_Cedula, h.Nombre, h.Apellido, h.Email, h.Direccion, h.Ciudad, h.Telefono, h.Celular)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create cliente", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clienteCreated)
}

func (h *ClienteHandler) GetAllCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	clientes, err := cliente.GetAll()

	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get clientes", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(clientes)
}

func (h *ClienteHandler) GetByIdCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid cliente ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	cliente, err := cliente.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get cliente", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cliente)
}

func (h *ClienteHandler) UpdateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid cliente ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	err = cliente.Update(h.RNC_Cedula, h.Nombre, h.Apellido, h.Email, h.Direccion, h.Ciudad, h.Telefono, h.Celular, uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update cliente", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente updated successfully"})
}

func (h *ClienteHandler) DeleteCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid cliente ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	err = cliente.Delete(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete cliente", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente deleted successfully"})
}
