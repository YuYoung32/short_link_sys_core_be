/**
 * Created by YuYoung on 2023/4/12
 * Description: 磁盘监控
 */

package monitor

import (
	"github.com/shirou/gopsutil/v3/disk"
	"short_link_sys_core_be/log"
)

const devicename = "vda"

var (
	diskTotal uint64
)

func diskStaticInfoSet() {
	if usageStat, err := disk.Usage(`/`); err == nil {
		diskTotal = usageStat.Total
	} else {
		log.MainLogger.WithField("module", "monitor").Error("diskStaticInfoSet: ", err)
	}
}

// diskDynamicInfo 磁盘动态信息 每隔一秒调用
func diskDynamicInfo() (readBytes, writeBytes, diskUsed, diskFree uint64) {
	ioCounters, err := disk.IOCounters(devicename)
	if err != nil {
		log.MainLogger.WithField("module", "monitor").Error("diskDynamicInfo: ", err)
		return
	}
	readBytes = ioCounters[devicename].ReadBytes
	writeBytes = ioCounters[devicename].WriteBytes
	if usageStat, err := disk.Usage(`/`); err == nil {
		diskUsed = usageStat.Used
		diskFree = usageStat.Free
	} else {
		log.MainLogger.WithField("module", "monitor").Error("diskDynamicInfo: ", err)
	}
	return
}
