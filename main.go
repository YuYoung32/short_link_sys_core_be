/**
 * Created by YuYoung on 2023/4/12
 * Description: 项目入口
 */

package main

import (
	"net/http"
	"short_link_sys_core_be/conf"
	"short_link_sys_core_be/database/mysql"
	"short_link_sys_core_be/handler/forward"
	"short_link_sys_core_be/handler/monitor"
	"short_link_sys_core_be/log"
)

func init() {
	conf.Init()
	log.Init()
	mysql.Init()
	monitor.Init()
	log.MainLogger.WithField("module", "main").Info("all module has init")
}

func main() {
	moduleLogger := log.MainLogger.WithField("module", "main")

	http.HandleFunc("/", monitor.MonitorHandler)
	http.HandleFunc("/:shortLink", forward.ForwardHandler)

	addr := conf.GlobalConfig.GetString("server.host") + ":" + conf.GlobalConfig.GetString("server.port")
	moduleLogger.Debug("listen: ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
		return
	}

	moduleLogger.Info("forward server is listening on port " + conf.GlobalConfig.GetString("server.port"))
}
