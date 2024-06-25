package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/models"
	"strconv"
)

var itemTransaction models.ItemTransaction

type ItemTransactionController struct {
	DB *sql.DB
}

func (cc *ItemTransactionController) CreateItemTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&itemTransaction); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := itemTransaction.CreateItemTransactionData(cc.DB); err != nil {
		http.Error(w, "Error creating item transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *ItemTransactionController) GetItemTransactionHandler(w http.ResponseWriter, r *http.Request) {
	ItemTransactions, err := itemTransaction.GetItemTransactionData(cc.DB)
	if err != nil {
		http.Error(w, "Error retrieving items transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ItemTransactions)
}

func (cc *ItemTransactionController) DeleteItemTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ItemTransactionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item transaction ID", http.StatusBadRequest)
		return
	}

	if err := itemTransaction.DeleteItemTransactionData(cc.DB, ItemTransactionID); err != nil {
		http.Error(w, "Error deleting item transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cc *ItemTransactionController) UpdateItemTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemTransactionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item transaction ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&itemTransaction); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	itemTransaction.ID = itemTransactionID

	if err := itemTransaction.UpdateItemTransactionData(cc.DB); err != nil {
		http.Error(w, "Error updating item transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
