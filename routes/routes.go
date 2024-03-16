package routes

import (
	"database/sql"

	"github.com/MochammadQemalFirza/assignment2/controller"
	"github.com/MochammadQemalFirza/assignment2/repository"
	"github.com/MochammadQemalFirza/assignment2/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine, db *sql.DB) {
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	g.POST("/orders", controller.CreateOrdersItemsHandler)
	g.GET("/orders", controller.GetAllOrdersItems)
}

