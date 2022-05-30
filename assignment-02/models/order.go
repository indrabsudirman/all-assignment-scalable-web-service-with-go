package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"not null;type:varchar(191)" json:"customer_name"`
	Items        []Item
	OrderedAt    time.Time  `json:"ordered_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

func (o *Order) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("before insert order to database")
	if len(o.CustomerName) < 5 {
		err = fmt.Errorf("customer name too short")
	}
	return err
}
