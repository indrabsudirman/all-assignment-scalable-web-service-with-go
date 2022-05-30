package params

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/models"
	"time"
)

type CreateOrder struct {
	OrderedAt    time.Time     `json:"ordered_at"`
	CustomerName string        `json:"customer_name"`
	Item         []models.Item `json:"items"`
}
