package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("before insert to table users")
	if len(u.Email) < 10 {
		err = fmt.Errorf("your email is too short")
	}
	return
}

func (u *User) Print() {
	fmt.Println("ID :", u.ID)
	fmt.Println("Email :", u.Email)
	for _, product := range u.Products {
		fmt.Println("\tProduct :", product.Name)
	}
}
