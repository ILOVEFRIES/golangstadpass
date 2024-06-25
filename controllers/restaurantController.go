package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/models"
	"strconv"
)

var restaurant models.Restaurant

type RestaurantController struct {
	DB *sql.DB
}

func (cc *RestaurantController) CreateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := restaurant.CreateRestaurantData(cc.DB); err != nil {
		http.Error(w, "Error creating restaurant", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *RestaurantController) GetRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	Restaurants, err := restaurant.GetRestaurantData(cc.DB)
	if err != nil {
		http.Error(w, "Error retrieving restaurants", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Restaurants)
}

func (cc *RestaurantController) DeleteRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RestaurantID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}

	if err := restaurant.DeleteRestaurantData(cc.DB, RestaurantID); err != nil {
		http.Error(w, "Error deleting restaurant", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cc *RestaurantController) UpdateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	restaurant.ID = restaurantID

	if err := restaurant.UpdateRestaurantData(cc.DB); err != nil {
		http.Error(w, "Error updating restaurant", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
