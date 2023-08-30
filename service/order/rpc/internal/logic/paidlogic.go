package logic

import (
	"context"

	"github.com/fly602/mall/service/order/model"
	"github.com/fly602/mall/service/order/rpc/internal/svc"
	"github.com/fly602/mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *order.PaidRequest) (*order.PaidResponse, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
	}
	res.Status = model.ORDER_STATU_PAYED
	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	l.Debug("订单支付成功, 订单信息: %+v", res)
	return &order.PaidResponse{}, nil
}
