package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"ops_tool/conf"
	"ops_tool/public"
	"os"
)

var db *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.Conf.System.Mysql.User,
		conf.Conf.System.Mysql.Password,
		conf.Conf.System.Mysql.Host,
		conf.Conf.System.Mysql.Port,
		conf.Conf.System.Mysql.Dbname,
	)

	// 也可以使用MustConnect连接不成功就panic
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		public.Log.Panicf("connect Mysql-DB failed, err:%v", err)
	}
	db.SetMaxOpenConns(conf.Conf.System.Mysql.MaxOpenCons)
	db.SetMaxIdleConns(conf.Conf.System.Mysql.MaxIdleCons)

	public.Log.Info("初始化mysql数据库成功!")
	return
}

// InsertUser 首先初始化加载数据
func InsertUser() {
	if conf.Conf.System.App.InitData {
		sqlFile, err := os.Open("dao/mysql/dinding.sql")
		if err != nil {
			public.Log.Errorf("open mysql failed, err:%v ", err)
			return
		}
		defer sqlFile.Close()

		sqlBytes, err := ioutil.ReadAll(sqlFile)
		if err != nil {
			public.Log.Errorf("ioutil sqlFile failed, err:%v ", err)
			return
		}
		sqlQuery := string(sqlBytes)
		_, err = db.Exec(sqlQuery)
		if err != nil {
			public.Log.Errorf("exec sqlFile failed, err:%v ", err)
			return
		}
		public.Log.Info("插入数据完成！")
	} else {
		public.Log.Info("自定义数据库初始化数据！")
		return
	}
}
