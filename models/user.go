package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64     `gorm: "primaryKey" json:"id"`
	FirstName string    `gorm: "not null" "nvachar(300)" json:"first_name"`
	LastName  string    `gorm: "not null" "nvachar(300)" json:"last_name"`
	Email 	  string    `gorm: "not null" "nvachar(300)" json:"email"`
	Phone     string    `gorm: "not null" "nvachar(300)" json:"phone"`
	UserName  string    `gorm: "not null" "nvachar(300)" json:"user_name"`
	Password  string    `gorm: "not null" "nvachar(300)" json:"password"`
	Avatar    string    `gorm: "nvachar(300) json:"avatar"`
	Status    int       `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt *time.Time  `json:"delete_at" gorm:"column:updated_at;"`
}

type UserRole struct {
	UserID int64
	User User `gorm:"foreignKey: UserID"`
	RoleID int64
	Role Role `gorm:"foreignKey: RoleID"`
	Status    int       `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time `gorm: "default:CURRENT_TIMESTAMP()" "type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm: "default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()" "type:datetime" json:"updated_at"`
	DeleteAt *time.Time `gorm: "type:datetime" json:"delete_at"`
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