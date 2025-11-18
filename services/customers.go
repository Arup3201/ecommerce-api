package services

import "github.com/Arup3201/ec-api/models"

func CustomerLogin() {
	models.CreateNewCustomer("first_name", "last_name", "email", "address", "city", "state", "pincode", "phone")
}
