package user

import (
	"context"
	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(req *types.UsernameReq) (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.CoreRpc.GetUserByUsername(l.ctx, &core.UsernameReq{Username: req.Username})

	if err != nil {
		return nil, err
	}

	return &types.UserInfo{
		Username:    data.Username,
		Password:    data.Password,
		LoginStatus: data.LoginStatus,
		Avatar:      data.Avatar,
		Email:       data.Email,
		Phone:       data.Phone,
		RoleName:    data.RoleName,
	}, nil
}
