package main

import (
	"github.com/Rizal-Nurochman/handler"
	"github.com/Rizal-Nurochman/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDB()

	r := gin.Default()

	r.GET("/", handler.HomeHandler)
	r.GET("/books", handler.GetAll)

	r.Run(":8000")
}
