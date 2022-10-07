package order

import "context"

type OrderRepo interface {
	CreateOrder(ctx context.Context, newOrder *Order) (err error)
	GetOrdersByUser(ctx context.Context, userId string) (orders []Order, err error)
}
