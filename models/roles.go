package models

import "time"

type Role struct {
	Id        int64      `json:"id" gorm: "primaryKey;column:id;"`
	RoleName  string     `json:"role_name" gorm:"not null;nvachar(300);column:role_name"`
	Description string  `json:"description" gorm:"nvachar(300);column:description;"`
	Status    int       `json:"status" gorm:"type:int(2);default:1;column:status;"`
	CreatedAt time.Time `json:"created_at" gorm:"TIMESTAMP DEFAULT CURRENT_TIMESTAMP();type:datetime;column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP();type:datetime;column:updated_at;"`
	DeleteAt *time.Time `json:"delete_at" gorm:"type:datetime;column:delete_at;"`
}
type RolePermission struct {
	RoleID       int64       `json:"role_id" gorm:"column:role_id;"`
	Role         Role        `gorm:"foreignKey:RoleID;"`
	PermissionID int64       `json:"permission_id" gorm:"column:permission_id;"`
	Permission   Permission  `gorm:"foreignKey:PermissionID;"`
	Status       int         `json:"status" gorm:"type:int(2);default:1;column:status;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}
