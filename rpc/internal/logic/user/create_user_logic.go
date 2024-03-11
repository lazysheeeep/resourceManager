package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/model"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User management
func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
	// todo: add your logic here and delete this line
	userDao := dao.NewUserDao(l.svcCtx.DbClient)

	user := model.User{
		Username: *in.Username,
		Email:    *in.Email,
		Phone:    *in.Phone,
		UUID:     uuid.New(),
	}
	user.SetPassword(*in.Password)

	err := userDao.Create(user)
	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail", user))
		return nil, errors.New("创建用户失败")

	}

	return &core.BaseUUIDResp{
		Uuid: user.UUID.String(),
		Msg:  "用户创建成功",
	}, nil
}
