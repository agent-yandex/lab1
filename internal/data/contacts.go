package data

import (
	"time"

	"lab1/internal/models"
)

var TestContacts = []models.Contact{
	{
		ID:         1,
		Username:   "john_doe",
		GivenName:  "John",
		FamilyName: "Doe",
		Phone: []struct {
			TypeID      int `json:"type_id"`
			CountryCode int `json:"country_code"`
			Operator    int `json:"operator"`
			Number      int `json:"number"`
		}{
			{TypeID: 1, CountryCode: 1, Operator: 555, Number: 1234567}, // Мобильный
		},
		Email:     []string{"john.doe@example.com"},
		Birthdate: time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:         2,
		Username:   "jane_smith",
		GivenName:  "Jane",
		FamilyName: "Smith",
		Phone: []struct {
			TypeID      int `json:"type_id"`
			CountryCode int `json:"country_code"`
			Operator    int `json:"operator"`
			Number      int `json:"number"`
		}{
			{TypeID: 1, CountryCode: 44, Operator: 777, Number: 9876543}, // Мобильный
			{TypeID: 2, CountryCode: 44, Operator: 20, Number: 12345678}, // Домашний
		},
		Email:     []string{"jane.smith@example.com", "jane.work@example.org"},
		Birthdate: time.Date(1985, time.July, 22, 0, 0, 0, 0, time.UTC),
	},
}
