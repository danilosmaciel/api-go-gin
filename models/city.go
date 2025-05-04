package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	IbgeCode int32  `json:"id" gorm:"unique"`
	Name     string `json:"nome"`
	StateID  int32
}
