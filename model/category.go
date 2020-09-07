package model

import "github.com/jinzhu/gorm"

// Category model
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;" json:"name"`
}
