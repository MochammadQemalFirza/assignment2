package domain

import "time"

type ItemsOrder struct{
	Orders Orders
	Items Items
}


type Items struct {
	ItemID int `json:"item_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Orders Orders `json:"orders"`
}

type Orders struct {
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}