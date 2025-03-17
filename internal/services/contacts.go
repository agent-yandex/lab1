package services

import (
	"net/http"
	"strconv"

	"lab1/internal/data"
	"lab1/internal/models"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.TestContacts)
}

func CreateContact(c *gin.Context) {
	var newContact models.Contact

	if err := c.BindJSON(&newContact); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, newContact)
}

func UpdateContact(c *gin.Context) {
	var updatedContact models.Contact

	if err := c.BindJSON(&updatedContact); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	for i, contact := range data.TestContacts {
		if contact.ID == updatedContact.ID {
			data.TestContacts[i] = updatedContact
			c.IndentedJSON(http.StatusOK, updatedContact)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Contact not found",
	})
}

func DeleteContact(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format: " + err.Error(),
		})
		return
	}

	for i, contact := range data.TestContacts {
		if contact.ID == id {
			data.TestContacts = append(data.TestContacts[:i], data.TestContacts[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Contact deleted",
				"id":      id,
			})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Contact not found",
	})
}
