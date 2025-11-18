package services

import "github.com/Arup3201/ec-api/models"

var Customer = customerService{}

type customerService struct{}

func (c customerService) Login() {
	models.CreateNewCustomer("first_name", "last_name", "email", "address", "city", "state", "pincode", "phone")
}
