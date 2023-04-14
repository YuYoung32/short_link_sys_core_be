/**
 * Created by YuYoung on 2023/4/12
 * Description: CPU监控
 */

package monitor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"short_link_sys_core_be/log"
	"time"
)

var (
	cpuModelName string
	cpuCoreNum   int
	cpuThreadNum int
	cpuCacheSize int // B
	cpuSpeed     int // MHz
	startTime    int64
)

func cpuStaticInfoSet() {
	logger := log.GetLogger()
	cpuInfo, err := cpu.Info()
	if err != nil {
		logger.Error(err)
		return
	}
	if cpuCoreNum, err = cpu.Counts(false); err != nil {
		logger.Error(err)
		return
	}
	if cpuThreadNum, err = cpu.Counts(true); err != nil {
		logger.Error(err)
		return
	}
	cpuModelName = cpuInfo[0].ModelName
	cpuCacheSize = int(cpuInfo[0].CacheSize)
	cpuSpeed = int(cpuInfo[0].Mhz)
	startTime = time.Now().Unix()
}

// cpuDynamicInfo CPU使用率 每隔一秒调用
func cpuDynamicInfo() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		log.GetLogger().Error(err)
		return 0
	}
	return percent[0]
}
