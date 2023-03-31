package models

import (
	"time"
)

type Product struct {
    Id                 int64           `json:"id" gorm:"primaryKey;column:id;"`
    SupplierID         int64           `json:"supplier_id" gorm:"column:supplier_id;"`
    Suppliers          Supplier        `json:"suppliers" gorm:"foreignKey:SupplierID;"`
    ProductName        string          `json:"product_name" gorm:"not null;nvarchar(300);column:product_name;"`
    ProductDescription string          `json:"product_description" gorm:"type:nvarchar(300);column:product_description;"`
    ProductPrice       float64         `json:"product_price" gorm:"not null;type:float(10,2);column:product_price;"`
    ProductImage       string          `json:"product_image" gorm:"nvarchar(300);column:product_image;"`
    Unit               string          `json:"unit" gorm:"nvarchar(300);column:unit;"`
    Discount           float64         `json:"discount" gorm:"type:float(10,2);default:0;column:discount;"`
    Status             int             `json:"status" gorm:"type:int(2);default:1;column:status;"`
    IsDelete int `json:"is_delete" gorm:"type:int(2);default:1;column:is_delete;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:updated_at;"`
}