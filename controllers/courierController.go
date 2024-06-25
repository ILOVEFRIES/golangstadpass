package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"stadpass/models"
	"strconv"

	"github.com/gorilla/mux"
)

var courier models.Courier

// CourierController handles HTTP requests related to couriers.
type CourierController struct {
	DB *sql.DB
}

// CreateCourierHandler creates a new courier record.
func (cc *CourierController) CreateCourierHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&courier); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := courier.CreateCourierData(cc.DB); err != nil {
		http.Error(w, "Error creating courier", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetCouriersHandler retrieves all courier records.
func (cc *CourierController) GetCouriersHandler(w http.ResponseWriter, r *http.Request) {
	couriers, err := courier.GetCourierData(cc.DB)
	if err != nil {
		http.Error(w, "Error retrieving couriers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(couriers)
}

// DeleteCourierHandler deletes a courier record by ID.
func (cc *CourierController) DeleteCourierHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courierID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid courier ID", http.StatusBadRequest)
		return
	}

	if err := courier.DeleteCourierData(cc.DB, courierID); err != nil {
		http.Error(w, "Error deleting courier", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateCourierHandler updates a courier record by ID.
func (cc *CourierController) UpdateCourierHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courierID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid courier ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&courier); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	courier.ID = courierID

	if err := courier.UpdateCourierData(cc.DB); err != nil {
		http.Error(w, "Error updating courier", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
