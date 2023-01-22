package config

import (
	jwtx "mall/service/common/jwt"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth jwtx.JwtAuth

	UserRpc zrpc.RpcClientConf
}
