package svc

import (
	jwtx "github.com/fly602/mall/service/common/jwt"
	"github.com/fly602/mall/service/pay/api/internal/config"
	"github.com/fly602/mall/service/pay/rpc/payclient"

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
