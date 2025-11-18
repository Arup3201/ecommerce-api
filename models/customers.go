package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Address   string
	City      string
	State     string
	PinCode   int
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateNewCustomer(firstName, lastName, email string,
	address, city, state, pincode, phone string) {
	ctx := context.Background()
	customer := &Customer{
		ID:        "UUID",
		FirstName: "Test",
		LastName:  "User",
	}
	err := gorm.G[Customer](db).Create(ctx, customer)
	if err != nil {
		fmt.Printf("[ERROR] gorm create error: %s", err)
	}
}
