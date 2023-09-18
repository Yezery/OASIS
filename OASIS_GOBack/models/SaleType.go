package models

type SaleType struct {
	Id   int    `gorm:"primaryKey;column:id"`
	Icon string `gorm:"column:icon"`
	Type string `gorm:"column:type"`
	Name string `gorm:"column:name"`
}
