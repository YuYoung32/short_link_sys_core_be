/**
 * Created by YuYoung on 2023/4/12
 * Description: 项目入口
 */

package main

import (
	"github.com/spf13/viper"
	"net/http"
	"short_link_sys_core_be/handler/forward"
	"short_link_sys_core_be/handler/monitor"
	"short_link_sys_core_be/log"
)

func main() {
	moduleLogger := log.MainLogger.WithField("module", "main")

	http.HandleFunc("/", monitor.MonitorHandler)
	http.HandleFunc("/:shortLink", forward.ForwardHandler)

	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
		return
	}

	moduleLogger.Info("forward server is listening on port " + viper.GetString("server.port"))
}
