package token

import (
	"context"
	"errors"
	"resourceManager/rpc/assignment"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/model"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *core.TokenListReq) (*core.TokenListResp, error) {
	// todo: add your logic here and delete this line

	page := model.Page{
		PageNum:  uint64(in.Page),
		PageSize: uint64(in.PageSize),
	}

	tokenDao := dao.NewTokenDao(l.svcCtx.DbClient)
	tokens, total, err := tokenDao.GetTokens(page)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("detail:", err.Error()))
		return nil, errors.New("获取tokens失败")
	}

	return assignment.AssignTokens(tokens, total), nil
}
