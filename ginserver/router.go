package ginserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
	"github.com/liuhangkaixcode/websocket/global"
)

func InitRouter(g *gin.Engine)  {
	pprof.Register(g)

	if global.Global_Config_Manger.Jaeger.Active == 1 {
		fmt.Println("=active")
		g.Use(jaegerCheck)
	}
	 //连接
    g.GET("/connect",checkUserId,ConnctServer)
    //发送消息服务器
    g.POST("/sendmsg",SendMsg)

}
