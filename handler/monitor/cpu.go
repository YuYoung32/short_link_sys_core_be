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
	cpuModel  string
	coreNum   int
	threadNum int
	cacheSize int // B
	cpuSpeed  int // MHz
	startTime time.Time
)

func cpuStaticInfoSet() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.MainLogger.WithField("module", "monitor").Error("cpuStaticInfoSet: ", err)
		return
	}
	if coreNum, err = cpu.Counts(false); err != nil {
		log.MainLogger.WithField("module", "monitor").Error("cpuStaticInfoSet: ", err)
		return
	}
	if threadNum, err = cpu.Counts(true); err != nil {
		log.MainLogger.WithField("module", "monitor").Error("cpuStaticInfoSet: ", err)
		return
	}
	cpuModel = cpuInfo[0].ModelName
	cacheSize = int(cpuInfo[0].CacheSize)
	cpuSpeed = int(cpuInfo[0].Mhz)
	startTime = time.Now()
}

// cpuUsage CPU使用率 每隔一秒调用
func cpuUsage() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		log.MainLogger.WithField("module", "monitor").Error("cpuUsage: ", err)
		return 0
	}
	return percent[0]
}
