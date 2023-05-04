/**
 * Created by YuYoung on 2023/4/12
 * Description: 数据库连接
 */

package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"short_link_sys_core_be/conf"
	_ "short_link_sys_core_be/conf"
	"short_link_sys_core_be/log"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	moduleLogger := log.GetLogger()

	var dsn = conf.GlobalConfig.GetString("mysql.username") + ":" +
		conf.GlobalConfig.GetString("mysql.password") + "@tcp(" +
		conf.GlobalConfig.GetString("mysql.host") + ":" +
		conf.GlobalConfig.GetString("mysql.port") + ")/" +
		conf.GlobalConfig.GetString("mysql.database")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.MainLogger, logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Warn,
		}),
		PrepareStmt: true,
	})
	if err != nil {
		moduleLogger.Error("failed to connect database: " + err.Error())
		panic(err)
	}
	moduleLogger.Info("connect database success")
	sqlDB, err := db.DB()
	if err != nil {
		moduleLogger.Error("failed to get sqlDB: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	autoMigrate()
}

// GetDBInstance 获取数据库实例, 其他包使用
func GetDBInstance() *gorm.DB {
	if db == nil {
		log.GetLoggerWithSkip(2).Error("db is nil")
		panic("db is nil")
	}
	return db
}

func autoMigrate() {
	db := GetDBInstance()
	autoMigrateVisitModel(db)
}
