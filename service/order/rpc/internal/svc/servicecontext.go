package svc

import (
	"time"

	"github.com/fly602/mall/service/order/model"
	"github.com/fly602/mall/service/order/rpc/internal/config"
	"github.com/fly602/mall/service/product/rpc/productclient"
	"github.com/fly602/mall/service/user/rpc/userclient"

	"github.com/RussellLuo/timingwheel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel

	UserRpc     userclient.User
	ProductRpc  productclient.Product
	Timingwheel timingwheel.TimingWheel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		OrderModel:  model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		Timingwheel: *timingwheel.NewTimingWheel(time.Second, model.ORDER_EXPIRE_MAX),
	}
}
