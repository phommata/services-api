package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string `gorm:"type:varchar(100)"`
	Versions 	[]Version `gorm:"polymorphic:Owner;"`
}