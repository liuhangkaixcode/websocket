package core

import (
	"fmt"
	"github.com/liuhangkaixcode/websocket/global"
	"github.com/panjf2000/ants/v2"
	"time"
)

func InitWebSocketBasic()  {

	go func() {
		for{
			time.Sleep(time.Second*10)
			//fmt.Println("交换机的容量是",GetHubInstance().GetAllportsMap())
		}
	}()

	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {

		if V,OK:=i.(WSMessage);OK {
			fmt.Println("i===========",V)
			dealWithMessage(V)
		}
	})
	global.Global_DealWsMsg_Pool = p
	fmt.Println("==========", err)
	if err != nil {
		panic("初始化antsPool失败")
	}else{
		fmt.Println("初始化dealwsMsgPool成功")
	}


}

