package models

import "github.com/jinzhu/gorm"

type NewsLetterMember struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique_index"`
}
