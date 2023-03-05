package models

import "time"

type Role struct {
	Id        int64      `gorm: "primaryKey" json:"id"`
	RoleName  string     `gorm: "not null" "nvachar(300)" json:"role_name"`
	Description string `gorm: "nvachar(300)" json:"description"`
	Status    int       `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time `gorm: "TIMESTAMP DEFAULT CURRENT_TIMESTAMP" "type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm: "TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" "type:datetime" json:"updated_at"`
	DeleteAt *time.Time `gorm: "type:datetime" json:"delete_at"`
}

type RolePermission struct {
	RoleID int64
	Role Role `gorm:"foreignKey: RoleID"`
	PermissionID int64
	Permission Permission `gorm:"foreignKey: PermissionID"`
	Status    int       `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time `gorm: "TIMESTAMP DEFAULT CURRENT_TIMESTAMP" "type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm: "TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" "type:datetime" json:"updated_at"`
	DeleteAt *time.Time `gorm: "type:datetime" json:"delete_at"`
}