package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

type DatabaseConf struct {
	Host        string `json:",env=DATABASE_HOST"`
	Port        int    `json:",env=DATABASE_PORT"`
	Username    string `json:",default=root,env=DATABASE_USERNAME"`
	Password    string `json:",optional,env=DATABASE_PASSWORD"`
	DBName      string `json:",default=source_manager,env=DATABASE_DBNAME"`
	Type        string `json:",default=mysql,env=DATABASE_TYPE"`
	MaxOpenConn int    `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`
	MaxIdleConn int    `json:",optional,default=20,env=DATABASE_MAX_CONNECTIONS"`
}

func (DatabaseConf *DatabaseConf) DbToString() string {
	return fmt.Sprintf("%s"+":"+"%s"+"@tcp("+"%s"+":"+"%d"+")/"+"%s"+"?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseConf.Username,
		DatabaseConf.Password,
		DatabaseConf.Host,
		DatabaseConf.Port,
		DatabaseConf.DBName)
}

func (DatabaseConf *DatabaseConf) NewDbClient(connRead, connWrite string) *gorm.DB {
	var ormLogger logger.Interface
	ormLogger = logger.Default

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(DatabaseConf.MaxOpenConn)
	sqlDB.SetMaxIdleConns(DatabaseConf.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	Db := db

	//主从配置，读写分离
	_ = Db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connRead)},
			Replicas: []gorm.Dialector{mysql.Open(connWrite)},
			Policy:   dbresolver.RandomPolicy{},
		},
	))

	migration(Db)

	return Db
}

func migration(db *gorm.DB) {
	err := db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate()

	if err != nil {
		fmt.Errorf("数据库迁移出错:%s", err)
	}
	return
}
