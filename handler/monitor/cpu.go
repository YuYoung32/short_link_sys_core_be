/**
 * Created by YuYoung on 2023/4/12
 * Description: CPU监控
 */

package monitor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"short_link_sys_core_be/log"
)

var (
	cpuModel  string
	coreNum   int
	threadNum int
	cacheSize int // B
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
}
