package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int64  `gorm:"primaryKey;autoIncrement:true"`
	UserName  string `gorm:"column:user_name;type:varchar(255);unique;not null" json:"user_name"`
	Password  string `gorm:"column:password" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (User) TableName() string {
	return "user"
}
