/**
 * Created by YuYoung on 2023/4/19
 * Description: 测试mapping相关函数
 */

package forward

import (
	"short_link_sys_core_be/conf"
	"testing"
)

func TestMappingRegion(t *testing.T) {
	conf.Init()
	Init()
	region, err := mappingRegion("4.4.4.4")
	if err != nil {
		t.Error(err)
	}
	t.Log(region)
	region, err = mappingRegion("114.34.21.6")
	if err != nil {
		t.Error(err)
	}
	t.Log(region)

}
