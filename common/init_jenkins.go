package common

import (
	"context"
	"github.com/bndr/gojenkins"
	"ops_tool/conf"
	"ops_tool/public"
)

var jk *gojenkins.Jenkins

func Init(url string, ctx context.Context) {
	var err error
	jk, err = gojenkins.CreateJenkins(nil, url, conf.Conf.JenkinsDingTalk.Jenkins.Name, conf.Conf.JenkinsDingTalk.Jenkins.Password).Init(ctx)
	if err != nil {
		public.Log.Error("connect jenkins failed,err:%v\n", err)
		return
	}
	public.Log.Info("connect jenkins Successfully")
}
