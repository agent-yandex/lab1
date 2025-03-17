package data

import "lab1/internal/models"

var TestGroups = []models.Group{
	{
		ID:          1,
		Title:       "Friends",
		Description: "Close friends group",
		Contacts:    []int{1}, // Связь с John Doe
	},
	{
		ID:          2,
		Title:       "Work Team",
		Description: "Colleagues from project team",
		Contacts:    []int{1, 2}, // Связь с John Doe и Jane Smith
	},
}
