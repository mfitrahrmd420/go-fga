package v1

import (
	engine "github.com/Calmantara/go-fga/config/gin"
	"github.com/Calmantara/go-fga/pkg/domain/order"
	"github.com/Calmantara/go-fga/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type OrderRouterImpl struct {
	ginEngine    engine.HttpServer
	routerGroup  *gin.RouterGroup
	orderHandler order.OrderHandler
}

func NewOrderRouter(ginEngine engine.HttpServer, orderHandler order.OrderHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/orders")
	return &OrderRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, orderHandler: orderHandler}
}

func (o *OrderRouterImpl) get() {
	// all path for get method are here
	o.routerGroup.GET("user/:userId", o.orderHandler.GetOrdersByUserHdl)
}

func (o *OrderRouterImpl) post() {
	// all path for post method are here
	o.routerGroup.POST("", o.orderHandler.InsertOrderHdl)
}

func (o *OrderRouterImpl) Routers() {
	o.post()
	o.get()
}
