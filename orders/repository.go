package orders

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Order, error)
	FindByID(ID int) (Order, error)
	Save(order Order) (Order, error)
	Update(order Order) (Order, error)
	Delete(order Order) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Order, error) {
	var orders []Order

	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *repository) FindByID(ID int) (Order, error) {
	var order Order

	err := r.db.Preload("Items").Where("id = ?", ID).Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) Save(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) Update(order Order) (Order, error) {
	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) Delete(order Order) (Order, error) {
	err := r.db.Delete(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
