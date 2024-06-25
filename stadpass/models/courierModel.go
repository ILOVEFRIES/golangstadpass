package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Courier represents a courier entity.
type Courier struct {
	ID             int
	CourierName    string
	PhoneNumber    string
	CourierPicture string
}

// CreateCourierData inserts a new courier record into the database.
func (c *Courier) CreateCourierData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO courier (courier_name, phone_number, courier_picture) VALUES ($1, $2, $3)",
		c.CourierName, c.PhoneNumber, c.CourierPicture)
	return err
}

// DeleteCourierData deletes a courier record from the database by ID.
func (c *Courier) DeleteCourierData(db *sql.DB, courierID int) error {
	_, err := db.Exec("DELETE FROM courier WHERE courier_id = $1", courierID)
	return err
}

// GetCourierData retrieves all courier records from the database.
func (c *Courier) GetCourierData(db *sql.DB) ([]Courier, error) {
	rows, err := db.Query("SELECT * FROM courier")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var couriers []Courier
	for rows.Next() {
		var courier Courier
		if err := rows.Scan(&courier.ID, &courier.CourierName, &courier.PhoneNumber, &courier.CourierPicture); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		couriers = append(couriers, courier)
	}

	return couriers, nil
}

// UpdateCourierData updates a courier record in the database.
func (c *Courier) UpdateCourierData(db *sql.DB) error {
	_, err := db.Exec("UPDATE courier SET courier_name = $1, phone_number = $2, courier_picture = $3 WHERE courier_id = $4",
		c.CourierName, c.PhoneNumber, c.CourierPicture, c.ID)
	return err
}
