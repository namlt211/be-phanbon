package models

import "time"

type Permission struct {
	Id        int64      `json:"id" gorm:"primaryKey;column:id;"`
	PermissionName  string     `json:"permission_name" gorm:"not null;nvachar(300);column:permission_name;" `
	Description string `json:"description" gorm:"nvachar(300);column:description;"`
	Status    int        `json:"status" gorm:"type: int(2);default:1;column:status;"`
	IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}