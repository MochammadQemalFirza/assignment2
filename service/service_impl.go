package service

import (
	"github.com/MochammadQemalFirza/assignment2/model/domain"
	web "github.com/MochammadQemalFirza/assignment2/model/web"
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

func (s *ServiceImpl) GetAllOrdersItems() ([]web.CustItem, error) {
	listOrdersItemsWeb := []web.CustItem{}
	// listItemsDomain := []domain.Items{}
	// listItemsWeb := []web.Items{}
	listOrdersItemsDomain, err := s.Repository.GetAllOrdersItems()
	if err != nil {
		return nil, err
	}

// for _,order:= range listOrdersItemsDomain{
// listOrdersItemsWeb = append(listOrdersItemsWeb, web.CustItem{
// 	OrderedAt: order.Orders.OrderedAt,
// 	CustomerName:  order.Orders.CustomerName,
// 	Items : listItemsWeb,
// 	})
// 	for _,item:= range listItemsDomain{
// 		listItemsWeb = append(listItemsWeb,web.Items{
// 			ItemCode: item.ItemCode,
// 			Description: item.Description,
// 			Quantity: item.Quantity,
// 		} )
// 	}
// }

	for _, item := range listOrdersItemsDomain {
		var listItemsWeb []web.Items 
		listItemsWeb = append(listItemsWeb, web.Items{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
		listOrdersItemsWeb = append(listOrdersItemsWeb, web.CustItem{
			OrderedAt:    item.Orders.OrderedAt,
			CustomerName: item.Orders.CustomerName,
			Items:        listItemsWeb,
		})
	}
	
	return listOrdersItemsWeb, nil
}

func NewService(Repository repository.Repository) Service {
	return &ServiceImpl{Repository: Repository}
}