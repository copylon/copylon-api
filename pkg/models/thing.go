package models

import "gorm.io/gorm"

type Thing struct {
	gorm.Model
	Name          string `json:"name"`
	AlternateName string `json:"alternateName"`
	Description   string `json:"description"`
	Identifier    string `json:"identifier"`
	Url           string `json:"url"`
}
