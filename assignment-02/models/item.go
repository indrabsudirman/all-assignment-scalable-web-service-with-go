package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID          uint       `gorm:"primaryKey" json:"item_id,omitempty"`
	ItemCode    string     `gorm:"not null;type:varchar(191)" json:"item_code"`
	Description string     `gorm:"not null;type:varchar(191)" json:"description"`
	Quantity    int        `gorm:"not null" json:"quantity"`
	OrderID     uint       `json:"order_id"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (i *Item) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("before insert item to database")
	if len(i.ItemCode) < 2 || len(i.Description) < 5 || i.Quantity <= 0 {
		err = fmt.Errorf("item code must bigger than 2 chars or description must bigger than 5 chars or quantity must bigger than 0")
	}
	return err
}
