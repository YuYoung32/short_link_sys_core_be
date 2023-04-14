/**
 * Created by YuYoung on 2023/4/12
 * Description: 内存监控
 */

package monitor

import (
	"github.com/shirou/gopsutil/v3/mem"
	"short_link_sys_core_be/log"
)

var (
	memTotal  uint64
	swapTotal uint64
)

func memStaticInfoSet() {
	if memoryStat, err := mem.VirtualMemory(); err == nil {
		memTotal = memoryStat.Total
		swapTotal = memoryStat.SwapTotal
	} else {
		log.GetLogger().Error(err)
	}
}

// memDynamicInfo 获取内存使用情况 每隔一秒调用
func memDynamicInfo() (memUsed, memFree, swapUsed uint64) {
	if memoryStat, err := mem.VirtualMemory(); err == nil {
		memUsed = memoryStat.Used
		memFree = memoryStat.Available
		swapUsed = memoryStat.SwapTotal - memoryStat.SwapFree
	} else {
		log.GetLogger().Error(err)
	}
	return
}
