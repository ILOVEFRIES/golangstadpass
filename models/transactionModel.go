package models

import (
	"database/sql"
	"log"
)

type Transaction struct {
	ID              int
	totalPrice      int
	tax             int
	discount        int
	totalPriceAfter int
	customerID      int
	courierID       int
	orderStatus     int
}

func (c *Transaction) CreateTransactionData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO transaction (transaction_id, transaction_time, total_price, tax, discount, total_price_after,customer_id, courier_id, order_status) VALUES ($8, NOW(), $1, $2, $3, $4, $5, $6, $7)",
		c.totalPrice, c.tax, c.discount, c.totalPriceAfter, c.customerID, c.courierID, c.orderStatus, c.ID)
	return err
}

func (c *Transaction) DeleteTransactionData(db *sql.DB, transactionID int) error {
	_, err := db.Exec("DELETE FROM transaction WHERE transaction_id = $1", transactionID)
	return err
}

func (c *Transaction) GetTransactionData(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query("SELECT * FROM restaurants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.ID, &transaction.totalPrice, &transaction.tax, &transaction.discount, &transaction.totalPriceAfter, &transaction.customerID, &transaction.courierID, &transaction.orderStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		Transactions = append(Transactions, transaction)
	}

	return Transactions, nil
}

func (c *Transaction) UpdateTransactionData(db *sql.DB) error {
	_, err := db.Exec("UPDATE transaction SET total_price = $2, tax = $3, discount = $4, total_price_after = $5, customer_id = $6, courier_id = $7, order_status = $8 WHERE transaction_id = $1",
		c.ID, c.totalPrice, c.tax, c.discount, c.totalPriceAfter, c.customerID, c.courierID, c.orderStatus)
	return err
}
