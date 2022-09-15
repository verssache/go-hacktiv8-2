package orders

type Service interface {
	FindAll() ([]Order, error)
	FindByID(ID FindOrderInput) (Order, error)
	Save(orderInput SaveOrderInput) (Order, error)
	Update(ID FindOrderInput, orderInput UpdateOrderInput) (Order, error)
	Delete(ID FindOrderInput) (Order, error)
	FindOrderPerson(ID FindOrderInput) (OrderPerson, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) FindByID(ID FindOrderInput) (Order, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *service) Save(orderInput SaveOrderInput) (Order, error) {
	order := Order{}
	order.CustomerName = orderInput.CustomerName

	var orderItems []Item
	for _, item := range orderInput.Items {
		orderItem := Item{}
		orderItem.Code = item.Code
		orderItem.Description = item.Description
		orderItem.Qty = item.Quantity
		orderItems = append(orderItems, orderItem)
	}
	order.Items = orderItems

	newOrder, err := s.repository.Save(order)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (s *service) Update(ID FindOrderInput, orderInput UpdateOrderInput) (Order, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return order, err
	}

	order.CustomerName = orderInput.CustomerName

	var orderItems []Item
	for _, item := range orderInput.Items {
		orderItem := Item{}
		orderItem.ID = item.ID
		orderItem.Code = item.Code
		orderItem.Description = item.Description
		orderItem.Qty = item.Quantity
		orderItems = append(orderItems, orderItem)
	}

	order.Items = orderItems

	updatedOrder, err := s.repository.Update(order)
	if err != nil {
		return updatedOrder, err
	}

	return updatedOrder, nil
}

func (s *service) Delete(ID FindOrderInput) (Order, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return order, err
	}

	deletedOrder, err := s.repository.Delete(order)
	if err != nil {
		return deletedOrder, err
	}

	return deletedOrder, nil
}

func (s *service) FindOrderPerson(ID FindOrderInput) (OrderPerson, error) {
	order, err := s.repository.FindByID(ID.ID)
	if err != nil {
		return OrderPerson{}, err
	}

	person, err := s.repository.Person()
	if err != nil {
		return OrderPerson{}, err
	}

	orderPerson := OrderPerson{}
	orderPerson.Order = order
	orderPerson.Person = person

	return orderPerson, nil
}
