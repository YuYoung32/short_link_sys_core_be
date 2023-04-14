/**
 * Created by YuYoung on 2023/4/13
 * Description: 磁盘监控测试
 */

package monitor

import (
	"testing"
	"time"
)

func TestDiskStaticInfoSet(t *testing.T) {
	diskStaticInfoSet()
	t.Log(diskTotal)
}

func TestDiskDynamicInfo(t *testing.T) {
	/*
		dd if=/dev/zero of=tempfile bs=1M count=1024 conv=fdatasync
		dd if=tempfile of=/dev/null bs=1M count=1024
	*/
	readBytes1, writeBytes1, _, _ := diskDynamicInfo()
	mbDiv := uint64(1024 * 1024)
	//gbDiv := mbDiv * 1024
	for {
		time.Sleep(time.Second)
		readBytes2, writeBytes2, used, free := diskDynamicInfo()
		t.Log((readBytes2-readBytes1)/mbDiv, (writeBytes2-writeBytes1)/mbDiv, used/mbDiv, free/mbDiv)
		readBytes1, writeBytes1 = readBytes2, writeBytes2
	}
}
