package models

import (
	"time"
)

type Product struct {
	Id        int64      `json:"id" gorm:"primaryKey;column:id;"`
	SupplierID int64 `json:"supplier_id" gorm:"column:supplier_id;"`
	Suppliers Supplier `json:"suppliers" gorm:"foreignKey:SupplierID;"`
	ProductName string `json:"product_name" gorm: "not null;nvachar(300);column:product_name;"`
	ProductDescription string `json:"product_description" gorm:"nvachar(300);column:product_description;"`
	ProductPrice float64 `json:"product_price" gorm:"not null;type:float(10,2);column:product_price;"`
	ProductImage string `json:"product_image "gorm: "nvachar(300);column:product_image;"`
	Status    int       `json:"status" gorm:"type:int(2);default:1;column:status;"`
	CreatedAt time.Time  `json:"created_at" gorm: "default:CURRENT_TIMESTAMP();type:datetime;column:created_at;"`
	UpdatedAt time.Time  ` json:"updated_at" gorm: "default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()" "type:datetime;column:updated_at"`
	DeleteAt  *time.Time `json:"delete_at" gorm: "type:datetime;column:delete_at;"`
}