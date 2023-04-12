/**
 * Created by YuYoung on 2023/4/12
 * Description:
 */

package monitor

import (
	"testing"
)

func TestCPUStaticInfoSet(t *testing.T) {
	cpuStaticInfoSet()
	t.Log(cpuModel)
	t.Log(coreNum)
	t.Log(threadNum)
	t.Log(cacheSize)
}
