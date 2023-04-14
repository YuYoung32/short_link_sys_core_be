/**
 * Created by YuYoung on 2023/4/13
 * Description: 监控发送相关结构体
 */

package monitor

type StaticInfo struct {
	CPUStaticInfo struct {
		Name      string `json:"name"`
		CoreNum   int    `json:"coreNum"`
		ThreadNum int    `json:"threadNum"`
		CacheSize int    `json:"cacheSize"`
		StartTime int64  `json:"startTime"`
	} `json:"cpuStaticInfo"`
	MemStaticInfo struct {
		PhysicalTotalSize uint64 `json:"physicalTotalSize"`
		SwapTotalSize     uint64 `json:"swapTotalSize"`
	} `json:"memStaticInfo"`
	DiskStaticInfo struct {
		DiskTotalSize uint64 `json:"diskTotalSize"`
	} `json:"diskStaticInfo"`
	NetStaticInfo struct {
		IPv4 string `json:"ipv4"`
		MAC  string `json:"mac"`
	} `json:"netStaticInfo"`
}

type DynamicInfo struct {
	CPUUsageRatioSec int `json:"cpuUsageRatioLastSec"`

	MemUsageSec uint64 `json:"memUsageLastSec"`
	MemAvailSec uint64 `json:"memAvailLastSec"`
	SwapUsage   uint64 `json:"swapUsageLastSec"`

	DiskReadSec  uint64 `json:"diskReadLastSec"`
	DiskWriteSec uint64 `json:"diskWriteLastSec"`
	DiskUsageSec uint64 `json:"diskUsageLastSec"`
	DiskAvailSec uint64 `json:"diskAvailLastSec"`

	NetRecvSec uint64 `json:"netRecvLastSec"`
	NetSendSec uint64 `json:"netSendLastSec"`
}
