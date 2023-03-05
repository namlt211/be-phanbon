package models

import "time"

type Order struct {
	Id int64 `gorm: "primaryKey" json:"id"`
	UserID int64
	User User `gorm:"foreignKey: UserID"`
	OrderTotal int `gorm: "type: int(9)" json:"order_total`
	Status    int        `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time  `gorm: "default: CURRENT_TIMESTAMP()" "type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm: "default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()" "type:datetime" json:"updated_at"`
	DeleteAt  *time.Time `gorm: "type:datetime" json:"delete_at"`
}

type OrderDetails struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	OrderID int64 
	Order Order `gorm:"foreignKey: OrderID"`
	ProductID int64
	Product Product `gorm:"foreignKey: ProductID"`
	Quantity int64 `gorm:"type:int(9)" json:"quantity"`
	Price float64 `gorm:"type:float(10,2)" json:"price"`
	Status    int        `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time  `gorm: "default: CURRENT_TIMESTAMP()" "type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm: "default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()" "type:datetime" json:"updated_at"`
	DeleteAt  *time.Time `gorm: "type:datetime" json:"delete_at"`
}