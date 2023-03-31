package models

import "time"

type Supplier struct {
	Id          int64      `json:"id" gorm:"primaryKey;column:id;"`
	Name        string     ` json:"name" gorm:"not null;column:name;"`
	Status    int       ` json:"status" gorm:"type:int(2);column:status;default:1;"`
	Phone string ` json:"phone" gorm:"column:phone;"`
	Description string     `json:"description" gorm:"column:description;"`
	IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}