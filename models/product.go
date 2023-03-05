package models

import "time"

type Product struct {
	Id        int64      `gorm: "primaryKey" json:"id"`
	ProductName string `gorm: "not null" "nvachar(300)" json:"product_name"`
	ProductDescription string `gorm: "nvachar(300)" json:"product_description"`
	ProductPrice float64 `gorm:"not null" "type: float(10,2)" json:"product_price"`
	ProductImage string `gorm: "nvachar(300)" json: "product_image" `
	Status    int        `gorm: "type: int(2)" json:"status"`
	CreatedAt time.Time  `gorm: "default: CURRENT_TIMESTAMP()" "type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm: "default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()" "type:datetime" json:"updated_at"`
	DeleteAt  *time.Time `gorm: "type:datetime" json:"delete_at"`
}