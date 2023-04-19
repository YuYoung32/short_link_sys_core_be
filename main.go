/**
 * Created by YuYoung on 2023/4/12
 * Description: 项目入口
 */

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"short_link_sys_core_be/conf"
	"short_link_sys_core_be/database/mysql"
	"short_link_sys_core_be/handler/forward"
	"short_link_sys_core_be/handler/monitor"
	"short_link_sys_core_be/log"
	"syscall"
	"time"
)

func init() {
	conf.Init()
	log.Init()
	mysql.Init()
	monitor.Init()
	forward.Init()
	log.GetLogger().Info("all module has init")
}

func main() {
	moduleLogger := log.GetLogger()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	//engine.Use(gin.LoggerWithWriter(log.MainLogger.Writer()))
	engine.Use(log.Middleware)

	engine.GET("/", monitor.MonitorHandler)
	engine.Any("/:shortLink", forward.ForwardHandler)

	runAddr := conf.GlobalConfig.GetString("server.host") + ":" + conf.GlobalConfig.GetString("server.port")
	srv := &http.Server{
		Addr:    runAddr,
		Handler: engine,
	}
	go func() {
		moduleLogger.Info("forward server is listening on " + runAddr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			moduleLogger.Error(err)
			panic(err)
		}
	}()

	// 阻塞, 等待结束
	sig := <-sigCh
	moduleLogger.Info("receive signal: ", sig, ", start to exit...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		moduleLogger.Error(err)
	}
}
