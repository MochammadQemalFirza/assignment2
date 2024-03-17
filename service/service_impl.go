package service

import (
	"errors"
	"fmt"

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

	groupedOrders := make(map[string]web.CustItem)


	listOrdersItemsDomain, err := s.Repository.GetAllOrdersItems()
	if err != nil {
		return nil, err
	}


	for _, item := range listOrdersItemsDomain {

		custItem, ok := groupedOrders[item.Orders.CustomerName]
		if !ok {
		
			custItem = web.CustItem{
				OrderedAt:    item.Orders.OrderedAt,
				CustomerName: item.Orders.CustomerName,
				Items:        []web.Items{},
			}
		}

	
		custItem.Items = append(custItem.Items, web.Items{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})

		
		groupedOrders[item.Orders.CustomerName] = custItem
	}


	var result []web.CustItem
	for _, custItem := range groupedOrders {
		result = append(result, custItem)
	}

	return result, nil
}

func(s *ServiceImpl)UpdateOrdersItems(orderID int, payload web.CustItem)(*web.CustItem, error){
	var updatedItems []domain.Items
	var items []domain.Items

	items, err := s.Repository.GetOrderById(orderID)
	
	if err != nil {
		return nil, err
	}

	order := domain.Orders{
		OrderID:      orderID,
		CustomerName: payload.CustomerName,
		OrderedAt:    payload.OrderedAt,
	}



	for _, item := range payload.Items {

		for _, dbItem := range items {
			if item.ItemCode == dbItem.ItemCode {
				updatedItem := domain.Items{
					ItemCode:    item.ItemCode,
					Description: item.Description,
					Quantity:    item.Quantity,
					Orders: domain.Orders{
						OrderID: orderID,
					},
				}
				updatedItems = append(updatedItems, updatedItem)
				break
			}
		}
	}

	for _, item := range payload.Items {
		found := false
		for _, dbItem := range items {
			if item.ItemCode == dbItem.ItemCode {
				found = true
				break
			}
		}

		if !found {
			newItem := domain.Items{
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				Orders: domain.Orders{
					OrderID: orderID,
				},
			}
			updatedItems = append(updatedItems, newItem)
		}
	}

	err = s.Repository.UpdateOrdersItems(order, updatedItems)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func(s *ServiceImpl)DeleteOrdersItemsByID(orderID int)error{

	existingItems, err := s.Repository.GetOrderById(orderID)
	fmt.Println(existingItems)
	if err != nil {
		return err
	}
	if len(existingItems) == 0 {
		return errors.New("order with given ID not found")
	}

	err = s.Repository.DeleteOrderById(orderID)

	if err != nil {
		return err
	}

	return nil
}

func NewService(Repository repository.Repository) Service {
	return &ServiceImpl{Repository: Repository}
}