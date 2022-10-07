package order

import (
	"context"
	"github.com/Calmantara/go-fga/config/postgres"
	"github.com/Calmantara/go-fga/pkg/domain/order"
)

type OrderRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewOrderRepo(pgCln postgres.PostgresClient) order.OrderRepo {
	return &OrderRepoImpl{pgCln: pgCln}
}

func (o OrderRepoImpl) CreateOrder(ctx context.Context, newOrder *order.Order) (err error) {
	err = o.pgCln.GetClient().
		Model(&order.Order{}).
		Create(newOrder).Error

	return
}

func (o OrderRepoImpl) GetOrdersByUser(ctx context.Context, userId string) (orders []order.Order, err error) {
	err = o.pgCln.GetClient().Raw("SELECT o.id, o.ordered_at, o.user_id FROM orders o INNER JOIN users u ON o.user_id = u.id").Scan(&orders).Error

	return
}
