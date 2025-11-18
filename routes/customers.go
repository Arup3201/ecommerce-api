package routes

import (
	"net/http"

	"github.com/Arup3201/ec-api/services"
)

var Customer = customerRoute{}

type customerRoute struct{}

func (c customerRoute) Login(w http.ResponseWriter, r *http.Request) {
	services.Customer.Login()
}
