package main

import (
	"APIGIN/Controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Controller.Connect()

	// Endpoints CRUD
	r.POST("/create", Controller.CreateHandler)
	r.GET("/retrieve/:id", Controller.RetrieveHandler)
	r.PUT("/update/:id", Controller.UpdateHandler)
	r.DELETE("/delete/:id", Controller.DeleteHandler)

	r.Run(":8080")
}
