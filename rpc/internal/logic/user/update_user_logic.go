package user

import (
	"context"
	"errors"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *core.UserInfo) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	userDao := dao.NewUserDao(l.svcCtx.DbClient)

	user, exist, err := userDao.GetByUserName(*in.Username)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail:", in))
		return nil, err
	}

	if !exist {
		err = errors.New("用户不存在")
		l.Logger.Errorw(err.Error(), logx.Field("detail", err))
		return nil, err
	}

	user.Email = *in.Email
	user.Phone = *in.Phone

	err = userDao.UpdateUser(*in.Username, user)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail:", user))
		return nil, err
	}

	return &core.BaseResp{Msg: "修改信息成功"}, nil
}
