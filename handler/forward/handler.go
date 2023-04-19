/**
 * Created by YuYoung on 2023/4/12
 * Description: 转发请求
 */

package forward

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"short_link_sys_core_be/database/mysql"
	"short_link_sys_core_be/log"
	"strings"
)

func Init() {
	mappingInit()
}

func ForwardHandler(ctx *gin.Context) {
	shortLink := ctx.Param("shortLink")
	log.GetLogger().Debug("shortLink: ", shortLink)
	longLink, err := mappingLink(shortLink)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.GetLogger().Debugf("record %v not found: %v", err.Error(), shortLink)
		} else {
			log.GetLogger().Error("failed to query database: " + err.Error())
		}
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// 非阻塞记录访问日志
	go func(addr string) {
		logger := log.GetLogger()
		parts := strings.Split(addr, ":")
		if len(parts) != 2 {
			logger.Error("invalid remote addr: " + addr)
			return
		}
		region, err := mappingRegion(parts[0])
		if err != nil {
			logger.Error("invalid remote addr: " + addr)
			region = "未知"
		}
		mysql.GetDBInstance().Create(&mysql.Visit{
			ShortLink: shortLink,
			IP:        parts[0],
			Region:    region,
		})
	}(ctx.Request.RemoteAddr)
	ctx.Redirect(http.StatusTemporaryRedirect, longLink)
}
