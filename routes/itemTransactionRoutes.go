package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/controllers"
)

func InitializeItemTransactionRoutes(cc *controllers.ItemTransactionController) *mux.Router {
	router := mux.NewRouter()

	// Create a new customer
	router.HandleFunc("/create_item_transaction", cc.CreateItemTransactionHandler).Methods(http.MethodPost)

	// Retrieve all customer
	router.HandleFunc("/show_item_transaction", cc.GetItemTransactionHandler).Methods(http.MethodGet)

	// Delete a customer by ID
	router.HandleFunc("/delete_item_transaction/{id}", cc.DeleteItemTransactionHandler).Methods(http.MethodDelete)

	// Update a customer by ID
	router.HandleFunc("/update_item_transaction/{id}", cc.UpdateItemTransactionHandler).Methods(http.MethodPut)

	return router
}
