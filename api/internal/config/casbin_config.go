package config

import (
	"github.com/casbin/casbin/persist"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	_ "github.com/go-sql-driver/mysql"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"log"
)

type CasbinConf struct {
	ModelText string `json:"modelText"`
}

func (conf CasbinConf) NewCasbin(dpType, dsn string) (*casbin.Enforcer, error) {
	adapter := gormadapter.NewAdapter(dpType, dsn, true)

	var text string
	if conf.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = conf.ModelText
	}

	m, err := model.NewModelFromString(text)
	logx.Must(err)

	enforcer, err := casbin.NewEnforcer(m, adapter)
	logx.Must(err)

	err = enforcer.LoadPolicy()
	logx.Must(err)

	return enforcer, nil
}

func (conf CasbinConf) MustNewCasbin(dbType, dsn string) *casbin.Enforcer {
	csb, err := conf.NewCasbin(dbType, dsn)
	if err != nil {
		logx.Errorw("initialize Casbin failed", logx.Field("detail:", err.Error()))
		log.Fatalf("initialize Casbin failed,err:%s", err.Error())
		return nil
	}
	return csb
}

func (conf CasbinConf) MustNewRedisWatcher(c redis.RedisConf, f func(str string)) persist.Watcher {
	w, err := rediswatcher.NewWatcher(c.Host, rediswatcher.WatcherOptions{
		Options: redis2.Options{
			Network:  "tcp",
			Password: c.Pass,
		},
		Channel:    "/casbin",
		IgnoreSelf: false,
	})
	logx.Must(err)

	err = w.SetUpdateCallback(f)
	logx.Must(err)

	return w
}

func (conf CasbinConf) MustNewCasbinWithNewWatcher(dbType, dsn string, c redis.RedisConf) *casbin.Enforcer {
	cbn := conf.MustNewCasbin(dbType, dsn)
	w := conf.MustNewRedisWatcher(c, func(data string) {
		rediswatcher.DefaultUpdateCallback(cbn)(data)
	})
	err := cbn.SetWatcher(w)
	logx.Must(err)
	err = cbn.SavePolicy()
	logx.Must(err)
	return cbn
}
