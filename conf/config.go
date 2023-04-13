/**
 * Created by YuYoung on 2023/4/12
 * Description: 数据库等配置文件解析
 */

package conf

import (
	"github.com/spf13/viper"
	"short_link_sys_core_be/log"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/root/go_projects/short_link_sys_core_be/conf")
	//viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.MainLogger.WithField("module", "conf_init").Error("failed to read config file: " + err.Error())
		panic(err)
	}
}
