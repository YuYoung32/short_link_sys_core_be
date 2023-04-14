/**
 * Created by YuYoung on 2023/4/12
 * Description: 数据库连接
 */

package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"short_link_sys_core_be/conf"
	_ "short_link_sys_core_be/conf"
	"short_link_sys_core_be/log"
	"time"
)

var db *gorm.DB

func Init() {
	var err error
	logger := log.MainLogger.WithField("func", "database_init")

	var dsn = conf.GlobalConfig.GetString("mysql.username") + ":" +
		conf.GlobalConfig.GetString("mysql.password") + "@tcp(" +
		conf.GlobalConfig.GetString("mysql.host") + ":" +
		conf.GlobalConfig.GetString("mysql.port") + ")/" +
		conf.GlobalConfig.GetString("mysql.database")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect database: " + err.Error())
		panic(err)
	}
	logger.Info("connect database success")
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("failed to get sqlDB: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	autoMigrate()
}

// GetDBInstance 获取数据库实例, 其他包使用
func GetDBInstance() *gorm.DB {
	if db == nil {
		log.MainLogger.WithField("module", "database").Error("db is nil")
		panic("db is nil")
	}
	return db
}

func autoMigrate() {
	db := GetDBInstance()
	autoMigrateVisitModel(db)
}
