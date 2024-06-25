package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/models"
	"strconv"
)

var item models.Item

type ItemController struct {
	DB *sql.DB
}

func (cc *ItemController) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := item.CreateItemData(cc.DB); err != nil {
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *ItemController) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	Items, err := item.GetItemData(cc.DB)
	if err != nil {
		http.Error(w, "Error retrieving items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Items)
}

func (cc *ItemController) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ItemID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if err := item.DeleteItemData(cc.DB, ItemID); err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cc *ItemController) UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	item.ID = itemID

	if err := item.UpdateItemData(cc.DB); err != nil {
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
