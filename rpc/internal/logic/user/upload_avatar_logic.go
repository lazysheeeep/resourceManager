package user

import (
	"context"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/qiniu"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadAvatarLogic) UploadAvatar(in *core.AvatarInfo) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	file, fileSize, err := qiniu.GetFile(in.Path)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail:", "获取头像失败"))
		return nil, err
	}

	url, err := qiniu.UploadToQiNi(l.svcCtx.Config.QiniuConf, file, fileSize)

	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail:", "上传头像到七牛云失败"))
		return nil, err
	}

	userDao := dao.NewUserDao(l.svcCtx.DbClient)
	user, _, _ := userDao.GetByUserId(in.UserId)
	user.Avatar = url
	err = userDao.UpdateUser(user.Username, user)
	if err != nil {
		l.Logger.Errorw(err.Error(), logx.Field("detail:", "数据库更新头像失败"))
		return nil, err

	}

	return &core.BaseResp{Msg: "上传头像成功"}, nil
}
