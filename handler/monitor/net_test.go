/**
 * Created by YuYoung on 2023/4/13
 * Description: 网络监控测试
 */

package monitor

import (
	"testing"
	"time"
)

func TestNetStaticInfoSet(t *testing.T) {
	netStaticInfoSet()
	t.Log("ipv4:", ipv4)
	t.Log("mac:", mac)
}

func TestNetDynamicInfo(t *testing.T) {
	netStaticInfoSet()

	sendBytes1, recvBytes1 := netDynamicInfo()
	t.Log(sendBytes1, recvBytes1)
	mbDiv := float64(1024 * 1024)
	for {
		time.Sleep(time.Second)
		sendBytes2, recvBytes2 := netDynamicInfo()
		t.Log(float64(sendBytes2-sendBytes1)/mbDiv, float64(recvBytes2-recvBytes1)/mbDiv)
		sendBytes1, recvBytes1 = sendBytes2, recvBytes2
	}
}
