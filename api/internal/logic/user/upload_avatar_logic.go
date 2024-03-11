package user

import (
	"context"
	"resourceManager/rpc/types/core"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAvatarLogic) UploadAvatar(req *types.AvatarInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.CoreRpc.UploadAvatar(l.ctx, &core.AvatarInfo{
		UserId: req.UserId,
		Path:   req.Path,
	})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Message: data.Msg}, nil
}
