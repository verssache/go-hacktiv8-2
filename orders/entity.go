package orders

import "time"

type Order struct {
	ID           int
	CustomerName string
	CreatedAt    time.Time
	Items        []Item `gorm:"Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Item struct {
	ID          int
	Code        string
	Description string
	Qty         int
	OrderID     int
}
