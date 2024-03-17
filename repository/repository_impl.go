package repository

import (
	"database/sql"
	"errors"

	"github.com/MochammadQemalFirza/assignment2/model/domain"
)

type RepositoryImpl struct {
	db *sql.DB
}

func(r *RepositoryImpl)CreateOrdersItems(orders domain.Orders, items []domain.Items)error{
	var orderID int
	tx, err := r.db.Begin()
	if err != nil {
		return errors.New("something went wrong")
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
		return errors.New("something went wrong")
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
		   return errors.New("something went wrong")
        }
    }

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errors.New("something went wrong")
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
		return nil,errors.New("something went wrong")
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
			return nil,errors.New("something went wrong")
		 }
		ListOrdersItems = append(ListOrdersItems, item)
	}

	return ListOrdersItems,nil

}

func(r *RepositoryImpl)GetOrderById(orderId int) ([]domain.Items, error){
	ListOrdersItemsByID := []domain.Items{}
	SelectOrdersById := `
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
   		 items ON orders.order_id = items.order_id
	WHERE 
		orders.order_id = $1;
	`

	result, err := r.db.Query(SelectOrdersById, orderId)

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
		ListOrdersItemsByID = append(ListOrdersItemsByID, item)
	}

	if err != nil {
		return nil, err
	}
	
	return ListOrdersItemsByID,nil
}

func(r *RepositoryImpl)	UpdateOrdersItems(orders domain.Orders, items []domain.Items) error{
	tx, err := r.db.Begin()
	if err != nil{
		return err
	}

	existingItems, err := r.GetOrderById(orders.OrderID)
	if err != nil {
		return err
	}
	if len(existingItems) == 0 {
		return errors.New("order with given ID not found")
	}

	UpdateOrdersItemsById := `
	UPDATE 
		orders 
	SET  
		Customer_name = $2, 
		updated_at = Now()
	WHERE 
		order_id = $1
	`
	_, err = r.db.Exec(UpdateOrdersItemsById, orders.OrderID, orders.CustomerName)

	if err != nil {
		tx.Rollback()
		return err
	}

	UpdateItemsOrdersByID := `
	UPDATE 
		items
	SET 
		item_code = $1,
		description = $2,
		quantity = $3
	WHERE 
		item_code = $1 AND order_id = $4
	`
	for _, item := range items {
		found := false
		for _, existingItem := range existingItems {
			if existingItem.ItemCode == item.ItemCode {
				found = true
				break
			}
		}
		if !found {
			return errors.New("item not found in order")
		}
		_, err = r.db.Exec(
			UpdateItemsOrdersByID, 
			item.ItemCode, 
			item.Description, 
			item.Quantity, 
			orders.OrderID)
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