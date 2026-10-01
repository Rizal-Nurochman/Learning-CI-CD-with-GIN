package handler

import (
	"github.com/Rizal-Nurochman/service"
	"github.com/gin-gonic/gin"
)

func HomeHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hello World",
	})
}

func GetAll(ctx *gin.Context)  {
	books, err := service.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err,
			"data": nil,
		})
	}

	ctx.JSON(200, gin.H{
		"message":"Get All Books Success!",
		"data": books,
	})
}