package user

import (
	"context"
	"errors"
	"resourceManager/rpc/assignment"
	"resourceManager/rpc/dao"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *core.IdReq) (*core.UserInfo, error) {
	// todo: add your logic here and delete this line

	userDao := dao.NewUserDao(l.svcCtx.DbClient)
	user, exist, err := userDao.GetByUserId(in.GetId())

	roleDao := dao.NewRoleDao(l.svcCtx.DbClient)
	role, err := roleDao.GerRoleById(user.RoleId)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail", user))
		return nil, err
	}

	if !exist {
		err = errors.New("用户不存在")
		l.Logger.Errorw(err.Error(), logx.Field("detail", err))
		return nil, err
	}

	return assignment.AssignUser(user, role), nil
}
