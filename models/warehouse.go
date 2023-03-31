package models

import "time"

type Warehouse struct {
	Id         int64      `json:"id" gorm:"column:"primaryKey;id";"`
	ProductID  int64      `json:"product_id" gorm:"column:product_id;"`
	Product    Product    `json:"product" gorm:"foreignKey:ProductID;"`
	ImportDate int64      `json:"import_date" gorm:"column:import_date;"`
	ExportDate int64      `json:"export_date" gorm:"column:export_date;"`
	Quantity   int64      `json:"quantity" gorm:"column:quantity;"`
	Unit       string     `json:"unit" gorm:"column:unit;"`
	IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt   *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}