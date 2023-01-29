package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	ORDER_STATU_NONE    = iota // 订单初始化状态
	ORDER_STATU_PAYED          // 订单已支付
	ORDER_STATU_PAYING         // 订单待支付
	ORDER_STATU_TIMEOUT        // 订单支付超时
	ORDER_STATU_FAILED         // 订单支付失败
)

var (
	ORDER_EXPIRE_MAX int64 = 60 // 订单最大超时失效时间
)
