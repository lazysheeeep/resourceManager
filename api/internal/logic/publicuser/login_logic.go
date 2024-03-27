package publicuser

import (
	"context"
	"resourceManager/api/internal/errors"
	"resourceManager/api/internal/middleware"
	"resourceManager/rpc/model"
	"resourceManager/rpc/types/core"
	"time"

	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" && l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errors.NewCodeAbortedError("该登录方式已被禁止")
	}

	if ok := l.svcCtx.Captcha.Verify("CAPTCHA_"+req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.GetUserByUsername(l.ctx, &core.UsernameReq{Username: req.Username})
		if err != nil {
			return nil, err
		}

		if !model.CheckPassword(req.Password, *user.Password) {
			return nil, errors.NewCodeWrongPassError("密码错误")
		}

		token, err := middleware.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire,
			middleware.WithOption("userId", user.Id), middleware.WithOption("roleId", user.RoleCodes))
		if err != nil {
			return nil, err
		}

		// 将token添加进mysql
		_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
			Uuid:      user.Id,
			Token:     &token,
			ExpiredAt: &l.svcCtx.Config.Auth.AccessExpire,
		})
		if err != nil {
			return nil, err
		}

		// 同时删除已经校验过的验证码
		_, err = l.svcCtx.Redis.Del("CAPTCHA" + req.CaptchaId)
		if err != nil {
			logx.Errorw("删除验证码失败", logx.Field("detail:", err.Error()))
		}

		return &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{
				Message: "登陆成功！",
			},
			Data: types.LoginInfo{
				Uuid:   *user.Id,
				Token:  token,
				Expire: uint64(l.svcCtx.Config.Auth.AccessExpire),
			},
		}, nil
	} else {
		return nil, errors.NewCodeWrongPassError("验证码错误")
	}
}
