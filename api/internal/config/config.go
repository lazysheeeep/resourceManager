package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         AuthConf
	RedisConf    redis.RedisConf
	CoreRpc      zrpc.RpcClientConf
	McmsRpc      zrpc.RpcClientConf
	DataBaseConf DatabaseConf
	CasbinConf   CasbinConf
	ProjectConf  ProjectConf
	CROSConf     CROSConf
}

type AuthConf struct {
	AccessSecret string `json:",optional,env=AUTH_SECRET"`
	AccessExpire int64  `json:",optional,env=AUTH_EXPIRE"`
}

type ProjectConf struct {
	DefaultRoleId           uint64 `json:",default=1"`
	EmailCaptchaExpiredTime int    `json:",default=300"`
	SmsTemplateId           string `json:",optional"`
	SmsAppId                string `json:",optional"`
	SmsSignName             string `json:",optional"`
	RegisterVerify          string `json:",default=captcha,options=[captcha,email,sms,sms_or_email]"`
	LoginVerify             string `json:",default=captcha,options=[captcha,email,sms,sms_or_email,all]"`
	ResetVerify             string `json:",default=email,options=[email,sms,sms_or_email]"`
}

type CROSConf struct {
	Address string `json:",env=CROS_ADDRESS"`
}
