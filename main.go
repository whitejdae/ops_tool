package main

import (
	"context"
	"fmt"
	"net/http"
	"ops_tool/conf"
	"ops_tool/dao/mysql"
	"ops_tool/public"
	"ops_tool/routes"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf.InitConfig()

	public.InitLogger()

	public.InitTrans("zh")

	mysql.InitDB()

	mysql.InsertUser()

	route := routes.SetupRoute()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.Conf.System.App.Host, conf.Conf.System.App.Port),
		Handler: route,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			public.Log.Errorf("listen: %v\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	public.Log.Info("Shutdown Server ...")

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		public.Log.Errorf("Server Shutdown: %s\n", err)
	}

	public.Log.Info("Server exiting!")
}
