package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/model"
	"time"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *core.IdReq) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	userDao := dao.NewUserDao(l.svcCtx.DbClient)

	user, exist, err := userDao.GetByUserId(in.GetId())

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail", user))
		return nil, errors.New("获取用户失败")
	}

	if !exist {
		err = errors.New("用户不存在")
		l.Logger.Errorw(err.Error(), logx.Field("detail", user))
		return nil, err
	}

	deletedUser := model.DeletedUser{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: gorm.DeletedAt{time.Now(), true},
		},
		Username:       user.Username,
		PasswordDigest: user.PasswordDigest,
		LoginStatus:    0,
		Avatar:         user.Avatar,
		Email:          user.Email,
		Phone:          user.Phone,
		RoleId:         1,
		UUID:           user.UUID,
	}

	deletedUserDao := dao.NewDeletedUserDao(l.svcCtx.DbClient)
	err = deletedUserDao.Create(deletedUser)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail", user))
		return nil, errors.New("删除用户失败")
	}

	err = userDao.DeleteUser(user)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail", user))
		return nil, errors.New("删除用户失败")
	}

	return &core.BaseResp{Msg: "用户删除成功"}, nil
}
