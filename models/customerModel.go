package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Customer struct {
	ID    int
	Seat  string
	Money float64
}

func (c *Customer) CreateCustomerData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO customers (seat, money) VALUES ($1, $2)",
		c.Seat, c.Money)
	return err
}

func (c *Customer) DeleteCustomerData(db *sql.DB, customerID int) error {
	_, err := db.Exec("DELETE FROM customers WHERE customer_id = $1", customerID)
	return err
}

func (c *Customer) GetCustomerData(db *sql.DB) ([]Customer, error) {
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Seat, &customer.Money); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *Customer) UpdateCustomerData(db *sql.DB) error {
	_, err := db.Exec("UPDATE customers SET seat = $1, money = $2 WHERE customer_id = $3",
		c.Seat, c.Money, c.ID)
	return err
}
