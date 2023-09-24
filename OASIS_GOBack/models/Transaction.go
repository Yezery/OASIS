package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Turnover float64 `gorm:"column:turnover" json:"turnover"`
}
