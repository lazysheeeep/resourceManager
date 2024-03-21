package base

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"resourceManager/rpc/dao"
	"resourceManager/rpc/model"
	"time"

	"resourceManager/rpc/internal/svc"
	"resourceManager/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *core.Empty) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line
	l.ctx = context.Background()

	// 初始化期间对redis上锁
	lock := redis.NewRedisLock(l.svcCtx.Redis, "init_database_lock")
	lock.SetExpire(60)
	if ok, err := lock.Acquire(); !ok || err != nil {
		if !ok {
			logx.Error("last initialization is running")
			return nil, errors.New("数据库正在初始化")
		} else {
			logx.Errorw(err.Error(), logx.Field("detail:", err.Error()))
			return nil, errors.New("redis错误")
		}
	}
	defer func() {
		recover()
		lock.Release()
	}()

	//判断数据库是否已经初始化
	if l.svcCtx.DbClient.Migrator().HasTable(&model.Flag{}) {
		return &core.BaseResp{Msg: "数据库已初始化"}, nil
	}

	//redis设计默认状态值
	_ = l.svcCtx.Redis.Set("database_error_msg", "")
	_ = l.svcCtx.Redis.Set("database_init_state", "0")

	//给一个管理员账号
	err := l.insertAdminUser()
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("detail:", err.Error()))
		return nil, errors.New("管理员帐号初始化出错")
	}

	return &core.BaseResp{
		Msg: "数据库初始化成功！",
	}, nil
}

func (l *InitDatabaseLogic) insertAdminUser() error {
	admin := model.User{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Username:    "admin",
		LoginStatus: 1,
		Email:       "",
		Phone:       "",
		RoleId:      0,
	}
	admin.SetPassword("123456")
	userDao := dao.NewUserDao(l.svcCtx.DbClient)
	err := userDao.Create(admin)
	if err != nil {
		return err
	}
	return nil
}
