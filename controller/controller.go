package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateOrdersItems(c *gin.Context)
	GetAllOrdersItems(c *gin.Context)
	UpdateOrdersItems(c *gin.Context)
}