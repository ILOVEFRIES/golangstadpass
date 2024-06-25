package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/models"
	"strconv"
)

var customer models.Customer

type CustomerController struct {
	DB *sql.DB
}

func (cc *CustomerController) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := customer.CreateCustomerData(cc.DB); err != nil {
		http.Error(w, "Error creating customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *CustomerController) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := customer.GetCustomerData(cc.DB)
	if err != nil {
		http.Error(w, "Error retrieving customers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (cc *CustomerController) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	if err := customer.DeleteCustomerData(cc.DB, customerID); err != nil {
		http.Error(w, "Error deleting customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cc *CustomerController) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	customer.ID = customerID

	if err := customer.UpdateCustomerData(cc.DB); err != nil {
		http.Error(w, "Error updating customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
