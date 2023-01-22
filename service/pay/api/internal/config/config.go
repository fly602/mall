package config

import (
	jwtx "mall/service/common/jwt"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth   jwtx.JwtAuth
	PayRpc zrpc.RpcClientConf
}
