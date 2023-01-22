package logic

import (
	"context"
	"database/sql"
	"fmt"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"
	"mall/service/user/rpc/types/user"

	"github.com/dtm-labs/client/dtmgrpc"

	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRevertLogic) CreateRevert(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 获取RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Uid,
		})
		if err != nil {
			return fmt.Errorf("用户不存在")
		}

		// 查询用户最新创建的订单
		resOrder, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.Uid)
		if err != nil {
			return fmt.Errorf("订单不存在")
		}
		// 修改订单状态9，标识订单已失效，并更新订单
		resOrder.Status = 9
		err = l.svcCtx.OrderModel.TxUpdate(l.ctx, tx, resOrder)
		if err != nil {
			return fmt.Errorf("订单更新失败")
		}
		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{}, nil
}
