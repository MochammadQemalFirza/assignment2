package repository

import (
	"database/sql"

	"github.com/MochammadQemalFirza/assignment2/model/domain"
)

type RepositoryImpl struct {
	db *sql.DB
}

func(r *RepositoryImpl)CreateOrdersItems(orders domain.Orders, items []domain.Items)error{
	var orderID int
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	CreateOrders :=`
	INSERT INTO "orders"
	(customer_name)
	VALUES($1) 
	RETURNING order_id`
	err = tx.QueryRow(
		CreateOrders, 
		orders.CustomerName,
		).Scan(&orderID)
    if err != nil {
		tx.Rollback()
        return err
    }

	CreateItems :=`
	INSERT INTO items 
	(order_id, item_code, description, quantity) 
	VALUES ($1, $2, $3, $4)
	`
	for _, item := range items {
        _, err := tx.Exec(
			CreateItems, 
			orderID, 
			item.ItemCode, 
			item.Description, 
			item.Quantity)
        if err != nil {
           tx.Rollback()
		   return err
        }
    }

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}
	
	return nil
}

func NewRepository(db *sql.DB) Repository{
	return &RepositoryImpl{db: db}
}