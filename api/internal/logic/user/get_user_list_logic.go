package user

import (
	"context"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.CoreRpc.GetUserList(l.ctx, &core.UserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	var users []types.UserInfo

	for _, v := range data.Data {
		user := types.UserInfo{
			Username:    v.Username,
			LoginStatus: v.LoginStatus,
			Avatar:      v.Avatar,
			Email:       v.Email,
			Phone:       v.Phone,
		}
		users = append(users, user)
	}

	return &types.UserListResp{
		Total: data.Total,
		Data:  users,
	}, nil
}
