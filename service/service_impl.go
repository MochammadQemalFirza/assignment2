package service

import (
	"github.com/MochammadQemalFirza/assignment2/model/domain"
	"github.com/MochammadQemalFirza/assignment2/model/web"
	"github.com/MochammadQemalFirza/assignment2/repository"
)

type ServiceImpl struct {
	Repository repository.Repository
}

func (s *ServiceImpl) CreateOrdersItems(payload web.CustItem) (*web.CustItem, error) {

	orders := domain.Orders{
		OrderedAt:    payload.OrderedAt,
		CustomerName: payload.CustomerName,
	}

	domainItems := []domain.Items{}

	for _, item := range payload.Items {
		domainItems = append(domainItems, domain.Items{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
		}


	err := s.Repository.CreateOrdersItems(orders, domainItems)
	if err != nil {
		return nil, err
	}


	return &web.CustItem{
		OrderedAt:    orders.OrderedAt,
		CustomerName: orders.CustomerName,
		Items:        payload.Items,
	}, nil
}

func NewService(Repository repository.Repository) Service {
	return &ServiceImpl{Repository: Repository}
}