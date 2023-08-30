package config

import (
	jwtx "github.com/fly602/mall/service/common/jwt"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth       jwtx.JwtAuth
	ProductRpc zrpc.RpcClientConf
}
