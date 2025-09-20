package models

type Item struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Quantity   int    `db:"quantity"`
	LocationID int    `db:"location_id"`
}
