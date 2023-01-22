package svc

import (
	jwtx "mall/service/common/jwt"
	"mall/service/pay/api/internal/config"
	"mall/service/pay/rpc/payclient"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	PayRpc    payclient.Pay
	JwtHeader rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PayRpc:    payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
		JwtHeader: jwtx.NewJwtheaderMiddleware(c.Auth).Handle,
	}
}
