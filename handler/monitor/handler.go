/**
 * Created by YuYoung on 2023/4/12
 * Description: 性能监控
 */

package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"short_link_sys_core_be/conf"
	"short_link_sys_core_be/log"
	"sync"
	"time"
)

func init() {
	setStaticInfo()
	setDynamicInfo()
	log.GetLogger().Info("monitor static and dynamic info init")
}

var (
	staticInfo StaticInfo

	dynamicInfo     DynamicInfo
	dynamicInfoLock sync.Mutex

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func setStaticInfo() {
	cpuStaticInfoSet()
	memStaticInfoSet()
	diskStaticInfoSet()
	netStaticInfoSet()

	staticInfo.CPUStaticInfo.Name = cpuModelName
	staticInfo.CPUStaticInfo.CoreNum = cpuCoreNum
	staticInfo.CPUStaticInfo.ThreadNum = cpuThreadNum
	staticInfo.CPUStaticInfo.CacheSize = cpuCacheSize
	staticInfo.CPUStaticInfo.StartTime = startTime

	staticInfo.MemStaticInfo.PhysicalTotalSize = memTotal
	staticInfo.MemStaticInfo.SwapTotalSize = swapTotal

	staticInfo.DiskStaticInfo.DiskTotalSize = diskTotal

	staticInfo.NetStaticInfo.MAC = mac
	staticInfo.NetStaticInfo.IPv4 = ipv4
}

// 每隔一段时间更新一次动态信息
func setDynamicInfo() {
	readBytesOld, writeBytesOld, _, _ := diskDynamicInfo()
	sendBytesOld, recvBytesOld := netDynamicInfo()
	go func() {
		for {
			time.Sleep(time.Duration(conf.GlobalConfig.GetInt("handler.monitor.interval")) * time.Millisecond)

			dynamicInfoLock.Lock()
			dynamicInfo.CPUUsageRatioSec = int(cpuDynamicInfo())

			memUsed, memAvail, swapUsed := memDynamicInfo()
			dynamicInfo.MemUsageSec = memUsed
			dynamicInfo.MemAvailSec = memAvail
			dynamicInfo.SwapUsage = swapUsed

			readBytesNew, writeBytesNew, diskUsed, diskFree := diskDynamicInfo()
			dynamicInfo.DiskReadSec = readBytesNew - readBytesOld
			dynamicInfo.DiskWriteSec = writeBytesNew - writeBytesOld
			dynamicInfo.DiskUsageSec = diskUsed
			dynamicInfo.DiskAvailSec = diskFree
			readBytesOld, writeBytesOld = readBytesNew, writeBytesNew

			sendBytesNew, recvBytesNew := netDynamicInfo()
			dynamicInfo.NetSendSec = sendBytesNew - sendBytesOld
			dynamicInfo.NetRecvSec = recvBytesNew - recvBytesOld
			sendBytesOld, recvBytesOld = sendBytesNew, recvBytesNew
			dynamicInfoLock.Unlock()
		}
	}()
}

func MonitorHandler(ctx *gin.Context) {
	moduleLogger := log.GetLogger()
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		moduleLogger.Debug(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			moduleLogger.Debug("handler.monitor handler", err)
		}
	}(conn)

	authed := false
	staticSend := false
	for {
		// 首次auth
		if !authed {
			err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			if err != nil {
				moduleLogger.Error("set read deadline failed:", err)
				return
			}
			_, message, err := conn.ReadMessage()
			if err != nil {
				moduleLogger.Error("read connection failed:", err)
				err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, ""))
				return
			}
			moduleLogger.Debug("remote auth: ", string(message))
			moduleLogger.Debug("local auth: ", conf.GlobalConfig.GetString("handler.monitor.authToken"))
			if string(message) == conf.GlobalConfig.GetString("handler.monitor.authToken") {
				authed = true
			} else {
				err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "no auth"))
				if err != nil {
					moduleLogger.Error("write connection failed:", err)
					return
				}
			}
		}

		// 发送静态信息, 仅发送一次
		if !staticSend {
			err := conn.WriteJSON(staticInfo)
			if err != nil {
				moduleLogger.Error("write connection failed: ", err)
				return
			}
			staticSend = true
		}

		// 发送动态信息, 每秒一次
		dynamicInfoLock.Lock()
		err := conn.WriteJSON(dynamicInfo)
		dynamicInfoLock.Unlock()
		if err != nil {
			moduleLogger.Debug("write connection failed:", err)
			return
		}
		time.Sleep(time.Second)
	}
}
