package models

import (
	"github.com/jinzhu/gorm"
)

type Version struct {
	gorm.Model
	ServiceId int
	Version   string `gorm:"type:varchar(10)";"polymorphic:Owner;"`
}

