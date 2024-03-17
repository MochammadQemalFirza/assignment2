package routes

import (
	"database/sql"

	"github.com/MochammadQemalFirza/assignment2/controller"
	"github.com/MochammadQemalFirza/assignment2/repository"
	"github.com/MochammadQemalFirza/assignment2/service"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func InitRouter(g *gin.Engine, db *sql.DB) {
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	docs.SwaggerInfo.Title = "H8 Assignment 2"
	docs.SwaggerInfo.Description = "Ini adalah tugas ke 2 dari kelas Kominfo"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8083"
	docs.SwaggerInfo.Schemes = []string{"http"}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	g.POST("/orders", controller.CreateOrdersItems)
	g.GET("/orders", controller.GetAllOrdersItems)
	g.PUT("/orders/:order_id", controller.UpdateOrdersItems)
	g.DELETE("/orders/:order_id", controller.DeleteOrdersItems)
}

