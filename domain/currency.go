package domain

import (
	"gorm.io/gorm"
)

type ModelBase struct {
	gorm.Model
}

type Currency struct {
	ModelBase
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Quotation Quotation `json:"quotation" gorm:"embedded"`
}

type Quotation struct {
	Buy  float32 `json:"buy"`
	Sell float32 `json:"sell"`
}
