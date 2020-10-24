package ginserver

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
)

func InitRouter(g *gin.Engine)  {
	pprof.Register(g)
	g.Use(checkUserId)
	 //连接
    g.GET("/connect",ConnctServer)
    //发送服务器
    g.POST("/sendmsg",SendMsg)

}
