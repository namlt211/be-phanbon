package models

import "time"

type Customer struct {
	Id        int64      `json:"id" gorm:"primaryKey;column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Phone     string     `json:"phone" gorm:"column:phone"`
	Address   string     `json:"address" gorm:"column:address`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}