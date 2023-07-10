package Models

import (
	"time"
)

type Product struct {
	ProductId   int       `json:"product_id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary_key"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"` // Updated foreign key reference
}
