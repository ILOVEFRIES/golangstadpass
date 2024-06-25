package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/controllers"
)

func InitializeItemRoutes(cc *controllers.ItemController) *mux.Router {
	router := mux.NewRouter()

	// Create a new customer
	router.HandleFunc("/create_item", cc.CreateItemHandler).Methods(http.MethodPost)

	// Retrieve all customer
	router.HandleFunc("/show_item", cc.GetItemHandler).Methods(http.MethodGet)

	// Delete a customer by ID
	router.HandleFunc("/delete_item/{id}", cc.DeleteItemHandler).Methods(http.MethodDelete)

	// Update a customer by ID
	router.HandleFunc("/update_item/{id}", cc.UpdateItemHandler).Methods(http.MethodPut)

	return router
}
