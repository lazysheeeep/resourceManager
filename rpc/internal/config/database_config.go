package config

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type DatabaseConf struct {
	Host        string `json:",env=DATABASE_HOST"`
	Port        int    `json:",env=DATABASE_PORT"`
	Username    string `json:",default=root,env=DATABASE_USERNAME"`
	Password    string `json:",optional,env=DATABASE_PASSWORD"`
	DBName      string `json:",default=source_manager,env=DATABASE_DBNAME"`
	SSLMode     string `json:"optional,env=DATABASE_SSL_MODE"`
	Type        string `json:",default=mysql,env=DATABASE_TYPE"`
	MaxOpenConn int    `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`
	CacheTime   int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
	DBPath      string `json:".optional,env=DATABASE_DBPATH"`
	MysqlConfig string `json:",optional,env=DATABASE_MYSQL_CONFIG"`
}

func (c DatabaseConf) NewNoCacheDriver() *entsql.Driver {
	db, err := sql.Open(c.Type, c.MysqlDSN())
	logx.Must(err)

	db.SetMaxOpenConns(c.MaxOpenConn)
	driver := entsql.OpenDB(c.Type, db)

	return driver
}

// MysqlDSN returns mysql DSN
func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.MysqlConfig)
}
