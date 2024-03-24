package token

import (
	"context"
	"errors"
	"resourceManager/rpc/assignment"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenByIdLogic) GetTokenById(in *core.UUIDReq) (*core.TokenInfo, error) {
	// todo: add your logic here and delete this line

	tokenDao := dao.NewTokenDao(l.svcCtx.DbClient)
	token, err := tokenDao.GetLatestToken(in.Uuid)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("detail:", err.Error()))
		return nil, errors.New("获取token失败")
	}

	return assignment.AssignToken(token), nil
}
