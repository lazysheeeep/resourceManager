package user

import (
	"context"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.IdReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.CoreRpc.DeleteUser(l.ctx, &core.IdReq{Id: req.Id})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Message: data.Msg}, nil
}
