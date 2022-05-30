package repositories

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/models"

	"github.com/jinzhu/gorm"
)

type ItemRepository interface {
	GetItem(uint, uint) (*models.Item, error)
	UpdateItem(uint, *models.Item) error
	DeleteItem(uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		db: db,
	}
}

// DeleteItem implements ItemRepository
func (i *itemRepository) DeleteItem(id uint) error {
	var item models.Item
	err := i.db.Where("order_id=?", id).Delete(&item).Error
	return err
}

// GetItem implements ItemRepository
func (i *itemRepository) GetItem(id uint, orderId uint) (*models.Item, error) {
	var item models.Item
	err := i.db.Where("id=?", id).Where("order_id=?", orderId).First(&item).Error
	return &item, err
}

// UpdateItem implements ItemRepository
func (i *itemRepository) UpdateItem(id uint, request *models.Item) error {
	var item models.Item
	err := i.db.Model(&item).Where("id=?", id).Updates(models.Item{
		ItemCode:    request.ItemCode,
		Description: request.Description,
		Quantity:    request.Quantity,
	}).Error
	return err
}
