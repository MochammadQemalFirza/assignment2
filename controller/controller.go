package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateOrdersItemsHandler(c *gin.Context)
}