package token

import (
	"context"
	"errors"
	"resourceManager/rpc/dao"
	"strconv"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *core.IdReq) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	tokenDao := dao.NewTokenDao(l.svcCtx.DbClient)
	id, _ := strconv.Atoi(in.Id)
	token, err := tokenDao.GetTokenById(uint(id))
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("detail:", err))
		return nil, errors.New("获取token时失败")
	}

	err = tokenDao.DeleteToken(token)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("detail:", err.Error()))
		return nil, errors.New("删除tokens失败")
	}

	return &core.BaseResp{Msg: "删除token失败"}, nil
}
