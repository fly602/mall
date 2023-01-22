package svc

import (
	jwtx "mall/service/common/jwt"
	"mall/service/user/api/internal/config"
	user "mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc   user.User
	JwtHeader rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRpc:   user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		JwtHeader: jwtx.NewJwtheaderMiddleware(c.Auth).Handle,
	}
}
