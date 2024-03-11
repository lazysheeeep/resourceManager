package user

import (
	"context"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.CoreRpc.CreateUser(l.ctx,
		&core.UserInfo{
			Username: req.Username,
			Password: req.Password,
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
