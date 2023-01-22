package svc

import (
	jwtx "mall/service/common/jwt"
	"mall/service/product/api/internal/config"
	"mall/service/product/rpc/productclient"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productclient.Product
	JwtHeader  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		JwtHeader:  jwtx.NewJwtheaderMiddleware(c.Auth).Handle,
	}
}
