package order

import (
	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/Calmantara/go-fga/pkg/domain/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHdlImpl struct {
	orderUsecase order.OrderUsecase
}

func NewOrderHandler(orderUsecase order.OrderUsecase) order.OrderHandler {
	return &OrderHdlImpl{orderUsecase: orderUsecase}
}

// GetOrdersByUserHdl godoc
// @Summary get orders from user
// @Description this api will get orders from specific user
// @Tags orders
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} order.Order
// @Router /v1/orders/user [get]
func (o OrderHdlImpl) GetOrdersByUserHdl(ctx *gin.Context) {
	userId := ctx.Param("userId")

	ords, err := o.orderUsecase.GetOrdersByUserSvc(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
			Status:  "fail",
			Message: "something went wrong",
		})

		return
	}

	ctx.JSON(http.StatusOK, message.Response{
		Status: "success",
		Data:   ords,
	})
}

// InsertOrderHdl godoc
// @Summary insert new order
// @Description this api will insert new order with existing user data
// @Tags orders
// @Accept json
// @Produce json
// @Param order body order.Order true "order object body"
// @Success 201 {object} order.Order
// @Router /v1/orders [post]
func (o OrderHdlImpl) InsertOrderHdl(ctx *gin.Context) {
	var newOrder order.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Status:  "fail",
			Message: "failed to bind payload",
		})

		return
	}

	ord, err := o.orderUsecase.InsertOrderSvc(ctx, newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
			Status:  "fail",
			Message: "something went wrong",
		})

		return

	}

	ctx.JSON(http.StatusCreated, message.Response{
		Status: "success",
		Data:   ord,
	})
}
