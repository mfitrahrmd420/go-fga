package order

import (
	"context"
	"github.com/Calmantara/go-fga/pkg/domain/order"
)

type OrderUsecaseImpl struct {
	orderRepo order.OrderRepo
}

func NewOrderUsecase(orderRepo order.OrderRepo) order.OrderUsecase {
	return &OrderUsecaseImpl{orderRepo: orderRepo}
}

func (o OrderUsecaseImpl) GetOrdersByUserSvc(ctx context.Context, userId string) (result []order.Order, err error) {
	orders, err := o.orderRepo.GetOrdersByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o OrderUsecaseImpl) InsertOrderSvc(ctx context.Context, input order.Order) (result order.Order, err error) {
	err = o.orderRepo.CreateOrder(ctx, &input)
	if err != nil {
		return order.Order{}, err
	}

	return input, nil
}
