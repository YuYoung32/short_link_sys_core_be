/**
 * Created by YuYoung on 2023/4/12
 * Description: 日志配置文件
 */

package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"short_link_sys_core_be/conf"
)

var (
	// MainLogger 全局Logrus实例
	MainLogger = logrus.New()
)

// Init 配置Logrus
func Init() {
	level := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	MainLogger.SetLevel(level[conf.GlobalConfig.GetString("log.level")])

	MainLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

	//logFilePath := logConf.DicPath + time.Now().Format("2006-01-02-15-04-05") + "_log.log"
	//file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//if err != nil {
	//	panic(err)
	//}

	//MainLogger.SetOutput(io.MultiWriter(file, os.Stdout))
	MainLogger.SetOutput(os.Stdout)

	MainLogger.Info("Logrus init success")
}
