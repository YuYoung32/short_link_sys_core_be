/**
 * Created by YuYoung on 2023/4/12
 * Description: CPU监控测试
 */

package monitor

import (
	"testing"
	"time"
)

func TestCPUStaticInfoSet(t *testing.T) {
	cpuStaticInfoSet()
	t.Log(cpuModelName)
	t.Log(cpuCoreNum)
	t.Log(cpuThreadNum)
	t.Log(cpuCacheSize)
	t.Log(cpuSpeed)
}

func TestCPUUsage(t *testing.T) {
	for {
		time.Sleep(time.Second)
		t.Log(cpuDynamicInfo())
	}
}
