package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/MochammadQemalFirza/assignment2/config"
	"github.com/MochammadQemalFirza/assignment2/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	envPath := ".env"

	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db:= db.CreateCon()
g:= gin.New()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port tidak ditemukan")
	}
routes.InitRouter(g,db)
g.Run(fmt.Sprintf(":%v", port))

}