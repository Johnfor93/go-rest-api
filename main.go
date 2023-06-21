package main

import (
	"github.com/Johnfor93/go-rest-api/controllers/contactcontroller"
	"github.com/Johnfor93/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/contacts", contactcontroller.ShowAll)
	r.GET("/api/contact/:id", contactcontroller.Find)
	r.POST("/api/contact", contactcontroller.Insert)
	r.PUT("/api/contact/:id", contactcontroller.Update)
	r.DELETE("/api/contact", contactcontroller.Delete)

	r.Run()
}