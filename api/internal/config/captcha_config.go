package config

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"image/color"
	"time"
)

type CaptchaConf struct {
	KeyLong   int    `json:",optional,default=6,env=CAPTCHA_KEY_LONG"`                                       // captcha length
	ImgWidth  int    `json:",optional,default=240,env=CAPTCHA_IMG_WIDTH"`                                    // captcha width
	ImgHeight int    `json:",optional,default=80,env=CAPTCHA_IMG_HEIGHT"`                                    // captcha height
	Driver    string `json:",optional,default=digit,options=[digit,string,math,chinese],env=CAPTCHA_DRIVER"` // captcha type
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
	Redis      *redis.Redis
}

func NewRedisStore(r *redis.Redis) *RedisStore {
	return &RedisStore{
		Expiration: time.Minute * 5,
		PreKey:     "CAPTCHA_",
		Redis:      r,
	}
}

func (r *RedisStore) Set(id string, value string) error {
	err := r.Redis.Setex(r.PreKey+id, value, int(r.Expiration.Seconds()))
	if err != nil {
		logx.Errorw("error occurs when captcha key sets to redis")
		return err
	}
	return nil
}

func (r *RedisStore) Get(key string, clear bool) string {
	value, err := r.Redis.Get(key)
	if err != nil {
		logx.Errorw("error occurs when captcha key got from redis")
		return ""
	}
	if clear {
		_, err := r.Redis.Del(key)
		if err != nil {
			logx.Errorw("error occurs when captcha key deleted from redis")
			return ""
		}
	}
	return value
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	key := r.PreKey + id
	v := r.Get(key, clear)
	return v == answer
}

func (r *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	r.Context = ctx
	return r
}

func MustNewRedisCaptcha(c CaptchaConf, r *redis.Redis) *base64Captcha.Captcha {
	driver := NewDriver(c)

	store := NewRedisStore(r)

	return base64Captcha.NewCaptcha(driver, store)
}

func NewDriver(c CaptchaConf) base64Captcha.Driver {
	var driver base64Captcha.Driver

	bgcolor := &color.RGBA{
		R: 254,
		G: 254,
		B: 254,
		A: 254,
	}

	// 字体名
	fonts := []string{
		"ApothecaryFont.ttf",
		"DENNEthree-dee.ttf",
		"Flim-Flam.ttf",
		"RitaSmith.ttf",
		"actionj.ttf",
		"chromohv.ttf",
	}

	switch c.Driver {
	case "digit":
		driver = base64Captcha.NewDriverDigit(c.ImgHeight, c.ImgWidth, c.KeyLong, 0.7, 80)
	case "string":
		driver = base64Captcha.NewDriverString(c.ImgHeight, c.ImgWidth, 12, 3, c.KeyLong, "qwertyupasdfghjkzxcvbnm23456789", bgcolor, nil, fonts)
	case "math":
		driver = base64Captcha.NewDriverMath(c.ImgHeight, c.ImgWidth, 12, 3, bgcolor, nil, fonts)
	default:
		driver = base64Captcha.NewDriverDigit(c.ImgHeight, c.ImgWidth, c.KeyLong, 0.7, 80)
	}

	return driver
}
