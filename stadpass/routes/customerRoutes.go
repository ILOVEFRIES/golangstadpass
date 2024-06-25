package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"stadpass/controllers"
)

func InitializeCustomerRoutes(cc *controllers.CustomerController) *mux.Router {
	router := mux.NewRouter()

	// Create a new customer
	router.HandleFunc("/create_customer", cc.CreateCustomerHandler).Methods(http.MethodPost)

	// Retrieve all customer
	router.HandleFunc("/show_customer", cc.GetCustomerHandler).Methods(http.MethodGet)

	// Delete a customer by ID
	router.HandleFunc("/delete_customer/{id}", cc.DeleteCustomerHandler).Methods(http.MethodDelete)

	// Update a customer by ID
	router.HandleFunc("/update_customer/{id}", cc.UpdateCustomerHandler).Methods(http.MethodPut)

	return router
}
