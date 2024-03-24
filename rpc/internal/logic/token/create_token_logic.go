package token

import (
	"context"
	"errors"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/model"
	"time"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// token management
func (l *CreateTokenLogic) CreateToken(in *core.TokenInfo) (*core.BaseUUIDResp, error) {
	// todo: add your logic here and delete this line

	token := model.Token{
		Status:    1,
		Uuid:      *in.Uuid,
		Token:     *in.Token,
		ExpiredAt: time.Now().Unix() + *in.ExpiredAt,
	}

	tokenDao := dao.NewTokenDao(l.svcCtx.DbClient)
	err := tokenDao.CreateToken(token)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("detail:", err.Error()))
		return nil, errors.New("创建token失败")
	}

	return &core.BaseUUIDResp{Uuid: *in.Uuid, Msg: "创建token成功"}, nil
}
