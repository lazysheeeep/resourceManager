package user

import (
	"context"
	"resourceManager/rpc/assignment"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/model"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *core.UserListReq) (*core.UserListResp, error) {
	// todo: add your logic here and delete this line

	userDao := dao.NewUserDao(l.svcCtx.DbClient)

	page := model.Page{
		PageNum:  in.Page,
		PageSize: in.PageSize,
	}

	users, total, err := userDao.GetUserList(page)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail", "获取用户失败"))
		return nil, err
	}

	return assignment.AssignListUser(users, uint64(total)), nil
}
