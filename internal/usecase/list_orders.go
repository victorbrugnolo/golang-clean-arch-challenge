package usecase

import "github.com/victorbrugnolo/golang-clean-arch-challenge/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var ordersDTO []OrderOutputDTO

	for _, order := range orders {
		ordersDTO = append(ordersDTO, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersDTO, nil
}
