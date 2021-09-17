package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string
}

type Account struct {
	Id              uint     `gorm:"primaryKey"`
	Age         	int
	UserName        string
}

func (Account) TableName() string {
	return "account"
}