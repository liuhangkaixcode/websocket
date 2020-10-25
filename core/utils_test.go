package core

import (
	"fmt"
	"testing"
)
//测试获取普通信息
func TestGetMsgInfo(t *testing.T) {
	message1 := getinfoMessageString("xxxx","debug")
	message2 := getinfoMessageString("xxxx","info",WithStaus("success"),WithReqName("订单接口"),WithServiceName("微服务1"))
	fmt.Println(message1,message2)
}

