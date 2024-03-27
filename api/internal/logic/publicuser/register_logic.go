package publicuser

import (
	"context"
	"resourceManager/api/internal/errors"
	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this lin
	if l.svcCtx.Config.ProjectConf.RegisterVerify != "captcha" {
		return nil, errors.NewCodeAbortedError("注册方式被禁止")
	}

	//只查询一次 如果错了就会生成新的验证码
	if ok := l.svcCtx.Captcha.Verify("CAPTCHA_"+req.CaptchaId, req.Captcha, true); ok {
		_, err := l.svcCtx.CoreRpc.CreateUser(l.ctx, &core.UserInfo{
			Username: &req.Username,
			Password: &req.Password,
			Email:    &req.Email,
		})
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.Redis.Del("CAPTCHA_" + req.CaptchaId)
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail:", err))
		}

		resp = &types.BaseMsgResp{
			Message: "注册成功",
		}
		return resp, nil
	} else {
		return nil, errors.NewCodeError(3, "验证码错误")
	}
}
