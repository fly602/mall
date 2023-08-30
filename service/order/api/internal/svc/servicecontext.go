package svc

import (
	jwtx "github.com/fly602/mall/service/common/jwt"
	"github.com/fly602/mall/service/order/api/internal/config"
	"github.com/fly602/mall/service/order/rpc/orderclient"
	"github.com/fly602/mall/service/product/rpc/productclient"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   orderclient.Order
	ProductRpc productclient.Product
	JwtHeader  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		JwtHeader:  jwtx.NewJwtheaderMiddleware(c.Auth).Handle,
	}
}
