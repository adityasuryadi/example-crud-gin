package models

import "time"

type GetCustomersResponse struct {
	Id        int64     `json:"id""`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCustomerRequest struct {
	Name    string `binding:"required" json:"name"`
	Email   string `binding:"required,email" json:"email"`
	Balance int64  `binding:"required,gte=10,lte=1000" json:"balance"`
}
