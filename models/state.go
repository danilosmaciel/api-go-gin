package models

import "gorm.io/gorm"

type State struct {
	gorm.Model
	Code   int32  `json:"id" gorm:"unique"`
	Name   string `json:"nome"`
	Sigla  string `json:"sigla"`
	Cities []City `gorm:"foreignKey:StateID;references:Code"`
}
