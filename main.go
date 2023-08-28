package main

import (
	"balance_bay/src/controllers"
	"balance_bay/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":8080")
}