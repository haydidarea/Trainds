package main

import (
	"github.com/Trainds/Goland/models"
	"github.com/Traninds/Goland/controllers/Productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.DB.ConnectDatabase()

	r.GET("/api/products", Productcontroller.Index)
	r.GET("/api/products/:id", Productcontroller.Show)
	r.POST("/api/products", Productcontroller.Create)
	r.PUT("/api/products", Productcontroller.Update)
	r.DELETE("/api/products", Productcontroller.Delete)

	r.Run()
}
