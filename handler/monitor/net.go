/**
 * Created by YuYoung on 2023/4/12
 * Description: 网络监控
 */

package monitor

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"io"
	"net/http"
	"short_link_sys_core_be/conf"
	_ "short_link_sys_core_be/conf"
	"short_link_sys_core_be/log"
)

var (
	ipv4 string
	mac  string

	getDefaultNIC func() netlink.Link
)

func setDefaultNIC() {
	logger := log.GetLogger()
	routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
	if err != nil {
		logger.Error(err)
	}

	for _, route := range routes {
		if route.Dst == nil {
			linkIndex := route.LinkIndex
			getDefaultNIC = func() netlink.Link {
				nic, err := netlink.LinkByIndex(linkIndex)
				if err != nil {
					logger.Error(err)
				}
				return nic
			}
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func getPublicIPv4() (ipv4 string) {
	logger := log.GetLogger()
	url := conf.GlobalConfig.GetString("monitor.publicIPQueryURL")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		logger.Error(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(resp.Body)

	// 读取响应数据
	if body, err := io.ReadAll(resp.Body); err == nil {
		ipv4 = string(body)
	} else {
		logger.Error(err)
		return
	}
	return
}

func netStaticInfoSet() {
	setDefaultNIC()
	mac = getDefaultNIC().Attrs().HardwareAddr.String()
	ipv4 = getPublicIPv4()
}

func netDynamicInfo() (sendBytes, recvBytes uint64) {
	sendBytes = getDefaultNIC().Attrs().Statistics.TxBytes
	recvBytes = getDefaultNIC().Attrs().Statistics.RxBytes
	return
}
