package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey;column:id;"`
	FirstName string    `json:"first_name" gorm:"not null;nvachar(300);column:first_name;"`
	LastName  string    `json:"last_name" gorm:"not null;nvachar(300);column:last_name;"`
	Email 	  string    `json:"email" gorm:"not null;nvachar(300);column:email;"`
	Phone     string    `json:"phone" gorm:"not null;nvachar(300);column:phone;"`
	UserName  string    `json:"user_name" gorm:"not null;nvachar(300);column:user_name;"`
	Password  string    `json:"password" gorm:"not null;nvachar(300);column:password;"`
	Avatar    string    `json:"avatar" gorm:"nvachar(300);column:avatar;"`
	Status    int       `json:"status" gorm:"default:1;column:status;"`
	IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt *time.Time  `json:"delete_at" gorm:"column:updated_at;"`
}

type UserRole struct {
	UserID int64 `json:"user_id" gorm:"column:user_id;"`
	User User `json:"user" gorm:"foreignKey: UserID"`
	RoleID int64 `json:"role_id"`
	Role Role `json:"role" gorm:"foreignKey:RoleID"`
	Status    int       `json:"status" gorm:"type: int(2);default:1"`
	IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}


type CustomerTime time.Time

func (t *CustomerTime) Scan(v interface{}) error {
	value, ok := v.([]uint8)
	if !ok {
		return fmt.Errorf("invalid type for customer time ")
	}
	if len(value) == 0 {
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(value))
	if err != nil {
		return err
	}
	*t = CustomerTime(parsedTime)
	return nil
}