package models

import "time"

type Payment struct {
	Id int64              `json:"id" gorm:"primaryKey;column:id;"`
	OrderID   int64      `json:"order_id" gorm:"column:order_id;"`
    Order     Order      `json:"order" gorm:"foreignKey:OrderID;"`
	Money 	float64 `json:"money" gorm:"column:money;"`
	Status    int        `json:"status" gorm:"type:int(2);default:1;"`
    IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}