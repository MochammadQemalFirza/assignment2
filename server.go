package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/MochammadQemalFirza/assignment2/config"
	"github.com/joho/godotenv"
)

func main() {

	envPath := ".env"

	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db.CreateCon()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port tidak ditemukan")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg:= "first commit"
		fmt.Fprint(w,msg)
	})

	http.ListenAndServe(fmt.Sprintf(":%v", port),nil)

}