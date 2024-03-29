package config

import "fmt"

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

func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.DBName)
}
