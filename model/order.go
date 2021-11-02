package model

import (
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Model
	CustomerID  uuid.UUID `gorm:"ForeignKey:CustomerID;type:varchar(36)" json:"customerId"`
	ItemName    string    `gorm:" type:varchar(30)" json:"itemName"`
	ItemDesc    string    `gorm:" type:varchar(100)" json:"itemDesc"`
	Quantity    int       `gorm:" type:int" json:"quantity"`
	CostPerUnit float64   `gorm:" type:double" json:"costPerUnit"`
	IsPaid      *bool     `gorm:" type:tinyint" json:"isPaid"`
}
