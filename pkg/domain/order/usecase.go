package order

import "context"

type OrderUsecase interface {
	GetOrdersByUserSvc(ctx context.Context, userId string) (result []Order, err error)
	InsertOrderSvc(ctx context.Context, input Order) (result Order, err error)
}
