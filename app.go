package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
	"stadpass/controllers"
	"stadpass/routes"
)

func main() {
	// Read environment variables for database configuration
	dbHost := os.Getenv("database-stadpass1.cbuoussck2e3.ap-southeast-1.rds.amazonaws.com")
	dbPort := os.Getenv("5432")
	dbUser := os.Getenv("postgres")
	dbPassword := os.Getenv("SGU2024Tangerang")
	dbName := os.Getenv("stadpass")

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a database connection
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Connected to the database!")

	// Initialize your controller
	courierController := &controllers.CourierController{DB: db}
	customerController := &controllers.CustomerController{DB: db}
	itemController := &controllers.ItemController{DB: db}
	itemTransactionController := &controllers.ItemTransactionController{DB: db}
	restaurantController := &controllers.RestaurantController{DB: db}
	transactionController := &controllers.TransactionController{DB: db}

	// Initialize the routes
	courierRouter := routes.InitializeCourierRoutes(courierController)
	customerRouter := routes.InitializeCustomerRoutes(customerController)
	itemRouter := routes.InitializeItemRoutes(itemController)
	itemTransactionRouter := routes.InitializeItemTransactionRoutes(itemTransactionController)
	restaurantRouter := routes.InitializeRestaurantRoutes(restaurantController)
	transactionRouter := routes.InitializeTransactionRoutes(transactionController)

	// Combine all routes
	router := http.NewServeMux()
	router.Handle("/couriers/", courierRouter)
	router.Handle("/customers/", customerRouter)
	router.Handle("/item/", itemRouter)
	router.Handle("/item_transaction/", itemTransactionRouter)
	router.Handle("/router/", restaurantRouter)
	router.Handle("/transaction", transactionRouter)

	// Start your server
	http.ListenAndServe(":8080", router)
}
