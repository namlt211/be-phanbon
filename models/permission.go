package models

import "time"

type Permission struct {
	Id        int64      `gorm: "primaryKey" json:"id"`
	PermissionName  string     `gorm: "not null" "nvachar(300)" json:"permission_name"`
	Description string `gorm: "nvachar(300)" json:"description"`
	Status    int        `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time  `gorm: "not null"  "default: CURRENT_TIMESTAMP()" "type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm: "not null"  "default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()" "type:datetime" json:"updated_at"`
	DeleteAt  *time.Time `gorm:  "type:datetime" json:"delete_at"`
}