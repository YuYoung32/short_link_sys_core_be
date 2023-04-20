/**
 * Created by YuYoung on 2023/4/12
 * Description: 访问 ORM Model
 */

package mysql

import (
	"gorm.io/gorm"
	"short_link_sys_core_be/log"
)

type Visit struct {
	ShortLink string `json:"shortLink" gorm:"type:varchar(255) COLLATE utf8_bin"`
	IP        string `json:"ip"`
	Region    string `json:"region"`
	VisitTime int64  `json:"visitTime" gorm:"autoCreateTime"`
}

func autoMigrateVisitModel(db *gorm.DB) {
	err := db.AutoMigrate(&Visit{})
	if err != nil {
		log.GetLogger().Error("auto migrate failed: " + err.Error())
	}
}
