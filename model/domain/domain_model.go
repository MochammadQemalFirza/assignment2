package domain

import "time"


type Items struct {
	ItemID int `json:"item_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Orders Orders `json:"orders"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Orders struct {
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}