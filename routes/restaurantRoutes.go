package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/controllers"
)

func InitializeRestaurantRoutes(cc *controllers.RestaurantController) *mux.Router {
	router := mux.NewRouter()

	// Create a new customer
	router.HandleFunc("/create_restaurant", cc.CreateRestaurantHandler).Methods(http.MethodPost)

	// Retrieve all customer
	router.HandleFunc("/show_restaurant", cc.GetRestaurantHandler).Methods(http.MethodGet)

	// Delete a customer by ID
	router.HandleFunc("/delete_restaurant/{id}", cc.DeleteRestaurantHandler).Methods(http.MethodDelete)

	// Update a customer by ID
	router.HandleFunc("/update_restaurant/{id}", cc.UpdateRestaurantHandler).Methods(http.MethodPut)

	return router
}
