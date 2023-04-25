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
)

func Init() {
	mappingInit()
}

func ForwardHandler(ctx *gin.Context) {
	logger := log.GetLogger()
	shortLink := ctx.Param("shortLink")
	longLink, err := mappingLink(shortLink)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Debugf("record %v not found: %v", err.Error(), shortLink)
		} else {
			logger.Error("failed to query database: " + err.Error())
		}
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	var addr string
	if ctx.Request.Header.Get("X-Real-IP") != "" {
		addr = ctx.Request.Header.Get("X-Real-IP")
	} else if ctx.Request.Header.Get("X-Forwarded-For") != "" {
		addr = ctx.Request.Header.Get("X-Forwarded-For")
	} else {
		addr = ctx.ClientIP()
	}

	// 非阻塞记录访问日志
	go func(addr string) {
		logger := log.GetLogger()

		region, err := mappingRegion(addr)
		if err != nil {
			logger.Debug("invalid remote addr: " + addr)
			region = "未知"
		}
		mysql.GetDBInstance().Create(&mysql.Visit{
			ShortLink: shortLink,
			IP:        addr,
			Region:    region,
		})
	}(addr)
	ctx.Redirect(http.StatusTemporaryRedirect, longLink)
}
