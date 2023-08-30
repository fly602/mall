package svc

import (
	"github.com/fly602/mall/service/order/rpc/orderclient"
	"github.com/fly602/mall/service/pay/model"
	"github.com/fly602/mall/service/pay/rpc/internal/config"
	"github.com/fly602/mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
