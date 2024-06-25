package models

import (
	"database/sql"
	"log"
)

type ItemTransaction struct {
	ID           int
	restaurantID int
	itemID       int
	quantity     int
	oldID        int
}

func (c *ItemTransaction) CreateItemTransactionData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO item_transactions (transaction_id, restaurant_id, item_id, quantity) VALUES ($1, $2, $3, $4)",
		c.ID, c.restaurantID, c.itemID, c.quantity)
	return err
}

func (c *ItemTransaction) DeleteItemTransactionData(db *sql.DB, itemTransactionID int) error {
	_, err := db.Exec("DELETE FROM item_transactions WHERE transaction_id = $1", itemTransactionID)
	return err
}

func (c *ItemTransaction) GetItemTransactionData(db *sql.DB) ([]ItemTransaction, error) {
	rows, err := db.Query("SELECT * FROM item_transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemTransactions []ItemTransaction
	for rows.Next() {
		var itemTransaction ItemTransaction
		if err := rows.Scan(&itemTransaction.itemID, &itemTransaction.restaurantID, &itemTransaction.itemID, &itemTransaction.quantity); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		itemTransactions = append(itemTransactions, itemTransaction)
	}

	return itemTransactions, nil
}

func (c *ItemTransaction) UpdateItemTransactionData(db *sql.DB) error {
	_, err := db.Exec("UPDATE item_transactions SET transaction_id = $1, restaurant_id = $2, item_id = $3, quantity = $4 WHERE transaction_id = $5",
		c.ID, c.restaurantID, c.itemID, c.quantity, c.oldID)
	return err
}
