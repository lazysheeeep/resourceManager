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

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *core.UsernameReq) (*core.UserInfo, error) {
	// todo: add your logic here and delete this line

	userDao := dao.NewUserDao(l.svcCtx.DbClient)
	user, exist, err := userDao.GetByUserName(in.Username)

	roleDao := dao.NewRoleDao(l.svcCtx.DbClient)
	role, _ := roleDao.GerRoleById(user.RoleId)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail:", err))
		return nil, err
	}

	if !exist {
		err = errors.New("用户不存在")
		l.Logger.Errorw(err.Error(), logx.Field("detail:", in))
		return nil, err
	}

	return assignment.AssignUser(user, role), nil
}
