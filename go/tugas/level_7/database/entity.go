package database

import "fmt"

type ItemEntity struct {
	Id     int64          `json:"id,omitempty"`
	Name   string         `json:"name"`
	Status string         `json:"status"`
	Amount int            `json:"amount"`
	Detail []DetailEntity `json:"detail,omitempty"`
}

type DetailEntity struct {
	Id     int64  `db:"id"`
	ItemId int64  `db:"item_id"`
	Name   string `db:"name"`
}

func (item ItemEntity) ToString() string {
	return fmt.Sprintf(
		"Item {ID: %d, Name: %s, Status: %s, Amount: %d}",
		item.Id,
		item.Name,
		item.Status,
		item.Amount,
	)
}
