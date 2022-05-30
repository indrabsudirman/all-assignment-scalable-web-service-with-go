package repositories

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/models"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	CreateOrder(*models.Order) error
	GetAllOrders() (*[]models.Order, error)
	UpdateOrder(uint, *models.Order) error
	DeleteOrder(uint) error
	FindByID(uint) (*models.Order, error)
	FindByIDItem(uint) (*models.Item, error)
}

type oderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &oderRepository{
		db: db,
	}
}

// CreateOrder implements OrderRepository
func (o *oderRepository) CreateOrder(order *models.Order) error {
	err := o.db.Create(order).Error
	return err
}

// GetAllOrders implements OrderRepository
func (o *oderRepository) GetAllOrders() (*[]models.Order, error) {
	orders := []models.Order{}
	err := o.db.Preload("Items").Find(&orders).Error
	return &orders, err
}

// DeleteOrder implements OrderRepository
func (o *oderRepository) DeleteOrder(id uint) error {
	order := models.Order{}
	e := o.db.First(&order, "id=?", id).Error
	if e != nil {
		return e
	}
	err := o.db.Delete(&order, "id=?", id).Error
	return err
}

func (o *oderRepository) FindByID(id uint) (*models.Order, error) {
	order := models.Order{}

	err := o.db.Find(&order, id).Error

	return &order, err
}

func (o *oderRepository) FindByIDItem(id uint) (*models.Item, error) {
	item := models.Item{}
	err := o.db.Find(&item, "order_id=?", id).Error
	return &item, err
}

// UpdateOrder implements OrderRepository
func (o *oderRepository) UpdateOrder(id uint, request *models.Order) error {
	var order models.Order

	err := o.db.First(&order, "id=?", id).Error
	if err != nil {
		return err
	}

	itemRepository := NewItemRepository(o.db)
	for _, item := range request.Items {
		_, err := itemRepository.GetItem(item.ID, id)

		if err != nil {
			return err
		}
		err = itemRepository.UpdateItem(item.ID, &models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
		if err != nil {
			return err
		}
	}
	err = o.db.Model(&order).Where("id=?", id).Updates(models.Order{
		CustomerName: request.CustomerName,
		UpdatedAt:    request.OrderedAt,
	}).Error
	return err

}
