package mysql

import (
	"ops_tool/conf"
	"ops_tool/public"
)

var DingMysql = dingMysqlInfo{}

type dingMysqlInfo struct {
	Id       int    `orm:"id"`
	UserName string `orm:"username"`
	Name     string `orm:"name"`
	Number   string `orm:"number"`
}

func GetUserPhone(user string) string {
	row := db.QueryRow("select * from dingding where username = ?;", user)
	err := row.Scan(&DingMysql.Id, &DingMysql.UserName, &DingMysql.Name, &DingMysql.Number)
	if err != nil {
		public.Log.Errorf("dao SelectUser select failed,err:%v", err)
		return conf.Conf.JenkinsDingTalk.DingDing.DefaultNotifier
	}
	return DingMysql.Number
}
