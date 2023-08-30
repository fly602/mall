package logic

import (
	"context"

	"github.com/fly602/mall/service/order/model"
	"github.com/fly602/mall/service/order/rpc/internal/svc"
	"github.com/fly602/mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ExpireLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpireLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpireLogic {
	return &ExpireLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpireLogic) Expire(id int64) {
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, id)
	if err != nil {
		if err == model.ErrNotFound {
			l.Error(status.Error(100, "订单不存在"))
			return
		}
		l.Error(status.Error(500, "查询失败"))
	}
	if res.Status != model.ORDER_STATU_PAYED {
		res.Status = model.ORDER_STATU_TIMEOUT
		err = l.svcCtx.OrderModel.Update(l.ctx, res)
		if err != nil {
			l.Error(status.Error(500, "更新数据库失败"))
			return
		}
		l.Info("订单超时, oid=%v", id)
		_, err = l.svcCtx.ProductRpc.DecrStock(l.ctx, &product.DecrStockRequest{
			Id:  res.Pid,
			Num: 1,
		})
		if err != nil {
			l.Error(status.Error(500, "库存回退失败"))
		} else {
			l.Debug("库存回退成功")
		}
	}
}
