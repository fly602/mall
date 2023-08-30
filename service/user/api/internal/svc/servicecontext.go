package svc

import (
	jwtx "github.com/fly602/mall/service/common/jwt"
	"github.com/fly602/mall/service/user/api/internal/config"
	user "github.com/fly602/mall/service/user/rpc/userclient"

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
