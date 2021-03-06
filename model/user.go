package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(250);not null"`
	Telephone string `grom:"varchar(110);not null;unique"`
	Password string `grom:"size:255;not null"`
}
