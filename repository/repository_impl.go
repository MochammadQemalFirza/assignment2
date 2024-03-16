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

func(r *RepositoryImpl)GetAllOrdersItems()([]domain.Items, error){
	ListOrdersItems :=[]domain.Items{}
	SelectOrders := `
	SELECT
    	orders.order_id,
    	orders.ordered_at,
    	orders.customer_name,
    	items.item_id,
    	items.item_code,
    	items.description,
    	items.quantity
	FROM
    	orders
	JOIN
   		 items ON orders.order_id = items.order_id;
	`

	result,err := r.db.Query(SelectOrders)
	if err!=nil{
		return nil, err
	}
	for result.Next() {
		item := domain.Items{}
		err := result.Scan(
			&item.Orders.OrderID,
			&item.Orders.OrderedAt,
			&item.Orders.CustomerName,
			&item.ItemID,
			&item.ItemCode,
			&item.Description,
			&item.Quantity,
		)
		if err != nil {
			return nil,err
		 }
		ListOrdersItems = append(ListOrdersItems, item)
	}

	return ListOrdersItems,nil

}

func NewRepository(db *sql.DB) Repository{
	return &RepositoryImpl{db: db}
}