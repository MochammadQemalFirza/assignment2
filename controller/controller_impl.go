package controller

import (
	"net/http"

	"github.com/MochammadQemalFirza/assignment2/model/web"
	"github.com/MochammadQemalFirza/assignment2/service"
	"github.com/gin-gonic/gin"
)

type ControllerImpl struct {
	Service service.Service
}

func (controller *ControllerImpl) CreateOrdersItemsHandler(c *gin.Context) {
	
	payload := web.CustItem{}
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := controller.Service.CreateOrdersItems(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order items"})
		return
	}

	
	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created order",
		"result": res,
	})
	
}

func NewController(Service service.Service) Controller {
	return &ControllerImpl{Service: Service}
}