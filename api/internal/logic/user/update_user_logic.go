package user

import (
	"context"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.CoreRpc.UpdateUser(l.ctx, &core.UserInfo{
		Username: req.Username,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Phone:    req.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{
		Message: data.Msg,
	}, nil
}
