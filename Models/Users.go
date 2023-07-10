package Models

import "time"

type User struct {
	UserId    int       `json:"user_id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary_key"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Products  []Product `json:"products" gorm:"foreignKey:UserID"` // Updated foreign key definition
}
