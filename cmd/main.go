package main

import (
	"log"

	"github.com/Rizal-Nurochman/handler"
	"github.com/Rizal-Nurochman/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connect database successfully!")

	r := gin.Default()

	r.GET("/", handler.HomeHandler)
	r.GET("/books", handler.GetAll)

	err = r.Run(":8000")
	if err != nil {
		panic(err.Error())
	}
	log.Fatal("Serve has ready. Run in port :8000")
}
