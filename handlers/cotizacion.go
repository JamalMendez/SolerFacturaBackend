package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/cotizacion"
	"ggstudios/solerfacturabackend/controllers/cotizacion_descripcion"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CotizacionHandler struct {
	cotizacion.CotizacionDTOSend
	CotizacionDescripcion []cotizacion_descripcion.ProductoDTO
}

func NewCotizacionHandler() *CotizacionHandler {
	return &CotizacionHandler{}
}

func (h *CotizacionHandler) CreateCotizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	cotizacionCreated, err := cotizacion.Create(
		h.ClienteID, h.TipoPagoID, h.CostoSubtotal, h.CostoTotal,
		h.Descuento, h.Envio, h.Secuencia, h.Cliente, h.Descripcion, h.EnDolares,
	)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create cotizacion", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	descripcionCreated, err := cotizacion_descripcion.Create(cotizacionCreated.ID, h.CotizacionDescripcion)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create cotizacion description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"cotizacion":  cotizacionCreated,
			"descripcion": descripcionCreated,
		},
	)
}

func (h *CotizacionHandler) GetAllCotizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cotizaciones, err := cotizacion.GetAll()
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get cotizaciones", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cotizaciones)
}

func (h *CotizacionHandler) GetByIdCotizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid cotizacion ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	cotizacion, err := cotizacion.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get cotizacion", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	descripcion, err := cotizacion_descripcion.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get cotizacion description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"cotizacion":  cotizacion,
			"descripcion": descripcion,
		},
	)
}

func (h *CotizacionHandler) UpdateCotizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid cotizacion ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	err = cotizacion.Update(
		h.ClienteID, h.TipoPagoID, h.CostoSubtotal, h.CostoTotal,
		h.Descuento, h.Envio, uint(id), h.Secuencia, h.Cliente, h.Descripcion, h.EnDolares,
	)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update cotizacion", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := cotizacion_descripcion.Update(uint(id), h.CotizacionDescripcion); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update cotizacion description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Cotizacion updated successfully"})
}

func (h *CotizacionHandler) DeleteCotizacion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid cotizacion ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	err = cotizacion.Delete(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete cotizacion", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := cotizacion_descripcion.Delete(uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete cotizacion description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Cotizacion deleted successfully"})
}
