package order

import "time"

type Order struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	OrderedAt time.Time `json:"ordered_At"`
	UserID    uint64    `json:"user_id"`
}
