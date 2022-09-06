package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id        int64 `gorm:"primaryKey;autoIncrement:true"`
	Name      string
	Email     string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *Customer) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Customer) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (Customer) TableName() string {
	return "customer"
}
