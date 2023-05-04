/**
 * Created by YuYoung on 2023/4/12
 * Description: 数据库等配置文件解析
 */

package conf

import (
	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper

func init() {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigName("config")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath("./conf")
	if err := GlobalConfig.ReadInConfig(); err != nil {
		panic(err)
	}

	if GlobalConfig.GetString("mode") == "dev" {
		devConfig := GlobalConfig.Sub("dev")
		for _, v := range devConfig.AllKeys() {
			GlobalConfig.Set(v, devConfig.Get(v))
		}
	}

}
