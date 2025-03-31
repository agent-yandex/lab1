package main

import (
	"lab1/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/groups", services.GetGroups)
	r.POST("/groups", services.CreateGroup)
	r.DELETE("/groups/:id", services.DeleteGroup)
	r.PUT("/groups", services.UpdateGroup)

	r.GET("/contacts", services.GetContacts)
	r.POST("/contacts", services.CreateContact)
	r.DELETE("/contacts/:id", services.DeleteContact)
	r.PUT("/contacts", services.UpdateContact)

	if err := r.Run(":6080"); err != nil {
		panic(err)
	}
}
