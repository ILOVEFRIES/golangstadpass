package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/controllers"
)

func InitializeTransactionRoutes(cc *controllers.TransactionController) *mux.Router {
	router := mux.NewRouter()

	// Create a new customer
	router.HandleFunc("/create_transaction", cc.CreateTransactionHandler).Methods(http.MethodPost)

	// Retrieve all customer
	router.HandleFunc("/show_transaction", cc.GetTransactionHandler).Methods(http.MethodGet)

	// Delete a customer by ID
	router.HandleFunc("/delete_transaction/{id}", cc.DeleteTransactionHandler).Methods(http.MethodDelete)

	// Update a customer by ID
	router.HandleFunc("/update_transaction/{id}", cc.UpdateTransactionHandler).Methods(http.MethodPut)

	return router
}
