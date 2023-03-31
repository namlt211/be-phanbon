package models

import "time"

type Order struct {
	Id int64              `json:"id" gorm:"primaryKey;column:id;"`
	CustomerID int64      `json:"customer_id" gorm:"column:customer_id;"`
	Customer Customer     `json:"customer" gorm:"foreignKey:CustomerID;"`
    TotalAmount float64 `json:"total_amount" gorm:"column:total_amount;"`
	Status    int        `json:"status" gorm:"type:int(2);default:1;"`
    IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}
type OrderDetails struct {
    ID        int64      `json:"id" gorm:"primaryKey;column:id;"`
    OrderID   int64      `json:"order_id" gorm:"column:order_id;"`
    Order     Order      `json:"order" gorm:"foreignKey:OrderID;"`
    ProductID int64      `json:"product_id" gorm:"column:product_id;"`
    Product   Product    `json:"product" gorm:"foreignKey:ProductID;"`
    Quantity  int64      `json:"quantity" gorm:"type:int(9);column:quantity;"`
    Status    int        `json:"status" gorm:"type:int(2);default:1;column:status;"`
    IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}