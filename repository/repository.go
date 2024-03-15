package repository

import "github.com/MochammadQemalFirza/assignment2/model/domain"

type Repository interface {
	CreateOrdersItems(orders domain.Orders, items []domain.Items) error
}