package repository

import "github.com/MochammadQemalFirza/assignment2/model/domain"

type Repository interface {
	CreateOrdersItems(orders domain.Orders, items []domain.Items) error
	GetAllOrdersItems()([]domain.Items, error)
	GetOrderById(orderId int) ([]domain.Items, error)
	UpdateOrdersItems(orders domain.Orders, items []domain.Items) error
	DeleteOrderById(orderId int) error
}