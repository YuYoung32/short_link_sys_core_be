/**
 * Created by YuYoung on 2023/4/12
 * Description: 短链映射到长链
 */

package forward

import (
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"io"
	"net/http"
	"short_link_sys_core_be/conf"
	"short_link_sys_core_be/database/mysql"
	"short_link_sys_core_be/log"
	"time"
)

var (
	linkModelInstance = mysql.Link{}
	queryURL          string
	bf                *bloom.BloomFilter
	bfDetectInterval  = time.Hour
)

func mappingInit() {
	bfDetectInterval = time.Duration(conf.GlobalConfig.GetInt("handler.forward.mappingIP.bloomFilter.bfDetectInterval")) * time.Second
	queryURL = conf.GlobalConfig.GetString("handler.forward.mappingIP.url")
	bf = bloom.NewWithEstimates(
		conf.GlobalConfig.GetUint("handler.forward.mappingIP.bloomFilter.expectedNumberOfElements"),
		conf.GlobalConfig.GetFloat64("handler.forward.mappingIP.bloomFilter.falsePositiveRate"))
	var lastAmount int64
	var shortLinks []string
	var lastRefreshTime time.Time
	mysql.GetDBInstance().Model(&linkModelInstance).Count(&lastAmount)
	// 首次加载必须初始化
	mysql.GetDBInstance().Model(&linkModelInstance).Pluck("shortLink", &shortLinks)
	lastRefreshTime = time.Now()
	for _, link := range shortLinks {
		bf.AddString(link)
	}
	go func() {
		for {
			var currentAmount int64
			mysql.GetDBInstance().Model(&linkModelInstance).Count(&currentAmount)
			// 一小时必须刷新一次
			if currentAmount == lastAmount && time.Now().Sub(lastRefreshTime) < time.Hour {
				// 数据库无变化
				time.Sleep(bfDetectInterval)
			} else {
				mysql.GetDBInstance().Model(&linkModelInstance).Pluck("shortLink", &shortLinks)
				lastRefreshTime = time.Now()
				lastAmount = currentAmount
				bf.ClearAll()
				for _, link := range shortLinks {
					bf.AddString(link)
				}
			}
		}
	}()
}

// mappingLink 短链映射到长链
func mappingLink(shortLink string) (string, error) {
	// 布隆过滤器过滤
	if !bf.TestString(shortLink) {
		// 一定不存在
		return "", fmt.Errorf("short link not exist")
	}

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
