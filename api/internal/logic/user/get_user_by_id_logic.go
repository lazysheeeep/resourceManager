package user

import (
	"context"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.IdReq) (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.IdReq{Id: req.Id})

	if err != nil {
		return nil, err
	}

	return &types.UserInfo{
		Username:    data.Username,
		LoginStatus: data.LoginStatus,
		Avatar:      data.Avatar,
		Email:       data.Email,
		Phone:       data.Phone,
		RoleName:    data.RoleName,
	}, nil
}
