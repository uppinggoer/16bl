package sql

import (
	"fmt"

	"util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var strConnect string

func init() {
	// 默认监听 127.0.0.1：8080
	var serverConf = util.Conf("db")
	host := serverConf.MustValue("db_host", "127.0.0.1").(string)
	port := serverConf.MustValue("db_port", "3306").(string)
	dbName, err := serverConf.Value("db_name")
	if nil != err {
		// log err
		return
	}
	dbUser, err := serverConf.Value("db_user")
	if nil != err {
		// log err
		return
	}
	passwd, err := serverConf.Value("db_passwd")
	if nil != err {
		// log err
		return
	}

	strConnect = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser.(string), passwd.(string), host, port, dbName.(string))

	// 启动时就打开数据库连接
	open()
}

func open() error {
	var err error

	DB, err = gorm.Open("mysql", strConnect)
	if err != nil {
		return err
	}

	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(20)

	// Disable table name's pluralization
	DB.SingularTable(true)

	// TODO:暂时全部开启日志
	DB.LogMode(true)

	return nil
}
