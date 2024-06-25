package models

import (
	"database/sql"
	"log"
)

type Restaurant struct {
	ID      int
	name    string
	picture string
}

func (c *Restaurant) CreateRestaurantData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO restaurants (restaurant_name, restaurant_picture) VALUES ($1, $2)",
		c.name, c.picture)
	return err
}

func (c *Restaurant) DeleteRestaurantData(db *sql.DB, restaurantID int) error {
	_, err := db.Exec("DELETE FROM restaurants WHERE restaurant_id = $1", restaurantID)
	return err
}

func (c *Restaurant) GetRestaurantData(db *sql.DB) ([]Restaurant, error) {
	rows, err := db.Query("SELECT * FROM restaurants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Restaurants []Restaurant
	for rows.Next() {
		var restaurant Restaurant
		if err := rows.Scan(&restaurant.ID, &restaurant.name, &restaurant.picture); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		Restaurants = append(Restaurants, restaurant)
	}

	return Restaurants, nil
}

func (c *Restaurant) UpdateRestaurantData(db *sql.DB) error {
	_, err := db.Exec("UPDATE restaurants SET restaurant_name = $2, restaurant_picture = $3 WHERE restaurant_id = $1",
		c.ID, c.name, c.picture)
	return err
}
