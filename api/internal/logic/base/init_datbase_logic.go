package base

import (
	"context"
	"errors"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatbaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatbaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatbaseLogic {
	return &InitDatbaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatbaseLogic) InitDatbase() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.InitDatabase(l.ctx, &core.Empty{})
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin policy", logx.Field("detail:", err.Error()))
		return nil, errors.New("初始化数据库失败")
	}
	return &types.BaseMsgResp{Message: result.Msg}, nil
}
