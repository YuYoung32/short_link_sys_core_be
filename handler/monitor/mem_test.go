/**
 * Created by YuYoung on 2023/4/13
 * Description: 内存监控测试
 */

package monitor

import (
	"testing"
	"time"
)

func TestMemStaticInfoSet(t *testing.T) {
	memStaticInfoSet()
	mbDiv := uint64(1024 * 1024)
	t.Log(memTotal/mbDiv, swapTotal/mbDiv)
}

func TestMemDynamicInfo(t *testing.T) {
	mbDiv := uint64(1024 * 1024)
	for {
		time.Sleep(time.Second)
		memUsed, memFree, swapUsed := memDynamicInfo()
		t.Log(memUsed/mbDiv, memFree/mbDiv, swapUsed/mbDiv)
	}

}
