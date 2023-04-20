/**
 * Created by YuYoung on 2023/4/20
 * Description:
 */

package mysql

import (
	"math/rand"
	"short_link_sys_core_be/conf"
	"strconv"
	"testing"
)

func TestData(t *testing.T) {
	conf.Init()
	Init()
	db := GetDBInstance()
	var shortLinks []string
	db.Model(&Link{}).Pluck("short_link", &shortLinks)
	shortLinksLen := len(shortLinks)
	provinces := []string{
		"北京市",
		"天津市",
		"河北省",
		"山西省",
		"内蒙古自治区",
		"辽宁省",
		"吉林省",
		"黑龙江省",
		"上海市",
		"江苏省",
		"浙江省",
		"安徽省",
		"福建省",
		"江西省",
		"山东省",
		"河南省",
		"湖北省",
		"湖南省",
		"广东省",
		"广西壮族自治区",
		"海南省",
		"重庆市",
		"四川省",
		"贵州省",
		"云南省",
		"西藏自治区",
		"陕西省",
		"甘肃省",
		"青海省",
		"宁夏回族自治区",
		"新疆维吾尔自治区",
		"台湾省",
		"香港特别行政区",
		"澳门特别行政区",
	}
	var beginTS = 1679326571
	var cross = 1681983825 - beginTS
	for i := 0; i < 1000; i++ {
		rand.Seed(int64(i))
		db.Create(&Visit{
			ShortLink: shortLinks[rand.Intn(shortLinksLen)],
			IP:        strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)),
			Region:    provinces[rand.Intn(len(provinces))],
			VisitTime: int64(rand.Intn(cross) + beginTS),
		})
	}
}
