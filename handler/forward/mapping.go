/**
 * Created by YuYoung on 2023/4/12
 * Description: 短链映射到长链
 */

package forward

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"short_link_sys_core_be/conf"
	"short_link_sys_core_be/database/mysql"
	"short_link_sys_core_be/log"
)

var (
	linkModelInstance = mysql.Link{}
	queryURL          string
)

func mappingInit() {
	queryURL = conf.GlobalConfig.GetString("handler.forward.mappingIP.url")
}

// mappingLink 短链映射到长链
func mappingLink(shortLink string) (string, error) {
	// 布隆过滤器过滤

	// 从Redis中查询热点数据

	// 从数据库中查询
	var longLink string
	err := mysql.GetDBInstance().Model(&linkModelInstance).Select("long_link").Where("short_link = ?", shortLink).Take(&longLink).Error
	if err != nil {
		return "", err
	}
	return longLink, nil
}

type ipInfo struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Area    string `json:"area"`
}

func mappingRegion(ip string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?ip=%s", queryURL, ip))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.GetLogger().Error(err)
		}
	}(resp.Body)
	//body, err := io.ReadAll(resp.Body)

	var ipJSON map[string]ipInfo
	//fmt.Println(string(body))

	if err = json.NewDecoder(resp.Body).Decode(&ipJSON); err != nil {
		return "", err
	}

	return ipJSON[ip].Country, nil
}
