package models

import (
	"database/sql"
	"log"
)

type Item struct {
	ID           int
	item         string
	price        float64
	itemPicture  string
	restaurantID int
}

func (c *Item) CreateItemData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO item (item, price, item_picture, restaurant_id) VALUES ($1, $2, $3, $4)",
		c.item, c.price, c.itemPicture, c.restaurantID)
	return err
}

func (c *Item) DeleteItemData(db *sql.DB, itemID int) error {
	_, err := db.Exec("DELETE FROM item WHERE item_id = $1", itemID)
	return err
}

func (c *Item) GetItemData(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT * FROM item")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.item, &item.price, &item.itemPicture, &item.restaurantID); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		items = append(items, item)
	}

	return items, nil
}

func (c *Item) UpdateItemData(db *sql.DB) error {
	_, err := db.Exec("UPDATE item SET item = $2, price = $3, item_picture= $4, restaurant_id = $5 WHERE item_id = $1",
		c.ID, c.item, c.price, c.itemPicture, c.restaurantID)
	return err
}
