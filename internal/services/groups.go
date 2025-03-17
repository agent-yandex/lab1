package services

import (
	"net/http"
	"strconv"

	"lab1/internal/data"
	"lab1/internal/models"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.TestGroups)
}

func CreateGroup(c *gin.Context) {
	group := &models.Group{}
	err := c.BindJSON(group)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, group)
}

func UpdateGroup(c *gin.Context) {
	groupUpdate := &models.Group{}
	err := c.BindJSON(groupUpdate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, group := range data.TestGroups {
		if group.ID == groupUpdate.ID {
			data.TestGroups[i] = *groupUpdate
			c.IndentedJSON(http.StatusOK, group)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Group not found"})
}

func DeleteGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, group := range data.TestGroups {
		if group.ID == id {
			data.TestGroups = append(data.TestGroups[:i], data.TestGroups[i+1:]...)
			c.IndentedJSON(http.StatusOK, group)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Group not found"})
}
