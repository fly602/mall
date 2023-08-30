package logic

import (
	"context"

	"github.com/fly602/mall/service/order/rpc/types/order"
	"github.com/fly602/mall/service/pay/model"
	"github.com/fly602/mall/service/pay/rpc/internal/svc"
	"github.com/fly602/mall/service/pay/rpc/types/pay"
	"github.com/fly602/mall/service/user/rpc/types/user"

	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	resOrder, err := l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 校验订单是否处于待支付状态
	if resOrder.Status != model.ORDER_STATU_PAYING {
		return nil, status.Error(100, "订单状态异常，取消支付")
	}

	// 校验支付金额是否正确
	if resOrder.Amount != in.Amount {
		return nil, status.Error(100, "金额不正确")
	}

	// 查询订单是否已经创建支付
	_, err = l.svcCtx.PayModel.FindOneByOid(l.ctx, in.Oid)
	if err == nil {
		return nil, status.Error(100, "订单已创建支付")
	}
	newPay := model.Pay{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: in.Amount,
		Source: 0,
		Status: 0,
	}

	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newPay.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CreateResponse{
		Id: newPay.Id,
	}, nil
}
