package service

import (
	"github.com/MochammadQemalFirza/assignment2/model/web"
)

type Service interface {
	CreateOrdersItems(payload web.CustItem)(*web.CustItem, error)
	GetAllOrdersItems()([]web.CustItem, error)
	UpdateOrdersItems(orderID int, payload web.CustItem)(*web.CustItem, error)
}