package models

import "time"

type Order struct {
	Id int64              `json:"id" gorm:"primaryKey;column:id;"`
	CustomerID int64      `json:"customer_id" gorm:"column:customer_id;"`
	Customer Customer     `json:"customer" gorm:"foreignKey:CustomerID;"`
    OrderDate *time.Time `json:"order_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP();column:order_date;"`
    OrderDetails []OrderDetails `json:"order_detail"`
	Status    int        `json:"status" gorm:"type:int(2);default:1;"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP();type:datetime;column:create_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP();type:datetime;column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"type:datetime;column:delete_at;"`
}
type OrderDetails struct {
    ID        int64      `json:"id" gorm:"primaryKey;column:id;"`
    OrderID   int64      `json:"order_id" gorm:"column:order_id;"`
    Order     Order      `json:"order" gorm:"foreignKey:OrderID;"`
    ProductID int64      `json:"product_id" gorm:"column:product_id;"`
    Product   Product    `json:"product" gorm:"foreignKey:ProductID;"`
    Quantity  int64      `json:"quantity" gorm:"type:int(9);column:quantity;"`
    Price     float64    `json:"price" gorm:"type:decimal(10,2);column:price;"`
    Status    int        `json:"status" gorm:"type:int(2);default:1;column:status;"`
    CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP();type:datetime;column:created_at;"`
    UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP();type:datetime;column:updated_at;"`
    DeletedAt *time.Time `json:"deleted_at" gorm:"type:datetime;column:deleted_at;index"`
}
