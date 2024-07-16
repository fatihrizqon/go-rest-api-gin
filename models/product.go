package models

type Product struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(64)" json:"name"`
	Detail string `gorm:"type:text" json:"detail"`
}
