package handlers

import (
	"encoding/json"
	"fmt"
	"ggstudios/solerfacturabackend/controllers/factura"
	"ggstudios/solerfacturabackend/controllers/factura_descripcion"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FacturaHandler struct {
	factura.FacturaDTOSend
	FacturaDescripcion []factura_descripcion.ProductoDTO
}

func NewFacturaHandler() *FacturaHandler {
	return &FacturaHandler{}
}

func (h *FacturaHandler) CreateFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	facturaCreated, err := factura.Create(
		h.NCFID, h.ClienteID, h.TipoPagoID,
		h.CostoSubtotal, h.CostoTotal, h.Descuento,
		h.Envio, h.Cliente, h.Descripcion, h.EnDolares, h.FechaVencimiento,
	)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create factura", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	descripcionCreated, err := factura_descripcion.Create(facturaCreated.ID, h.FacturaDescripcion)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to create factura description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"factura":     facturaCreated,
			"descripcion": descripcionCreated,
		},
	)
}

func (h *FacturaHandler) GetAllFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	facturas, err := factura.GetAll()
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get facturas", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(facturas)
}

func (h *FacturaHandler) GetByIdFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid factura ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	factura, err := factura.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get factura", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	descripcion, err := factura_descripcion.GetById(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to get factura description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"factura":     factura,
			"descripcion": descripcion,
		},
	)
}

func (h *FacturaHandler) UpdateFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid factura ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		response := fmt.Sprintf(`{"error": "Invalid request payload", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	err = factura.Update(
		h.NCFID, h.ClienteID, h.TipoPagoID,
		h.CostoSubtotal, h.CostoTotal, h.Descuento,
		h.Envio, uint(id), h.Cliente, h.Descripcion, h.EnDolares, h.FechaVencimiento,
	)
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update factura", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := factura_descripcion.Update(uint(id), h.FacturaDescripcion); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to update factura description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Factura updated successfully"})
}

func (h *FacturaHandler) DeleteFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := fmt.Sprintf(`{"error": "Invalid factura ID", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	err = factura.Delete(uint(id))
	if err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete factura", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	if err := factura_descripcion.Delete(uint(id)); err != nil {
		response := fmt.Sprintf(`{"error": "Failed to delete factura description", "details": "%s"}`, err.Error())
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Factura deleted successfully"})
}
