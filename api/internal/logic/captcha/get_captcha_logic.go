package captcha

import (
	"context"
	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaResp, err error) {

	if id, b64s, _, err := l.svcCtx.Captcha.Generate(); err != nil {
		logx.Errorw("生成验证码失败", logx.Field("detail:", err))
		return &types.CaptchaResp{Code: 13, Msg: "生成验证码失败", Data: types.CaptchaInfo{}}, nil
	} else {
		resp = &types.CaptchaResp{
			Msg: "生成验证码成功",
			Data: types.CaptchaInfo{
				CaptchaId: id,
				ImgPath:   b64s,
			},
		}
		return resp, nil
	}

	return
}
