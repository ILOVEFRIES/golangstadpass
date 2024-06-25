// routes/courierRoutes.go

package routes

import (
	"net/http"
	"stadpass/controllers"

	"github.com/gorilla/mux"
)

// InitializeCourierRoutes sets up the routes for courier-related endpoints.
func InitializeCourierRoutes(cc *controllers.CourierController) *mux.Router {
	router := mux.NewRouter()

	// Create a new courier
	router.HandleFunc("/create_courier", cc.CreateCourierHandler).Methods(http.MethodPost)

	// Retrieve all couriers
	router.HandleFunc("/show_couriers", cc.GetCouriersHandler).Methods(http.MethodGet)

	// Delete a courier by ID
	router.HandleFunc("/delete_courier/{id}", cc.DeleteCourierHandler).Methods(http.MethodDelete)

	// Update a courier by ID
	router.HandleFunc("/update_courier/{id}", cc.UpdateCourierHandler).Methods(http.MethodPut)

	return router
}
