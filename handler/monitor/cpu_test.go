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
	t.Log(cpuModel)
	t.Log(coreNum)
	t.Log(threadNum)
	t.Log(cacheSize)
	t.Log(cpuSpeed)
}

func TestCPUUsage(t *testing.T) {
	for {
		time.Sleep(time.Second)
		t.Log(cpuUsage())
	}
}
