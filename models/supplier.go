package models

import "time"

type Supplier struct {
	Id          int64      `json:"id" gorm:"primaryKey;column:id;"`
	Name        string     ` json:"name" gorm:"not null;column:name;"`
	Status    int       ` json:"status" gorm:"type:int(2);column:status;default:1;"`
	Description string     `json:"description" gorm:"column:description;"`
	CreatedAt   time.Time  `json:"created_at" gorm:"TIMESTAMP DEFAULT CURRENT_TIMESTAMP();type:datetime;column:created_at;" `
	UpdatedAt   time.Time  `json:"updated_at" gorm:"TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP();type:datetime;column:updated_at;"`
	DeleteAt    *time.Time `json:"delete_at" gorm:"type:datetime;column:deleted_at"`
}