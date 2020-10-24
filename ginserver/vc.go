package ginserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhangkaixcode/websocket/core"
	"github.com/liuhangkaixcode/websocket/global"
	"strings"
)

//开启连接服务器
func ConnctServer(c *gin.Context)  {
	userid:=c.Query("userid")

	b, _ := checkRequestIsFast(userid)
	if b == false {
		fmt.Println("请求的时间间隔应该小于"+fmt.Sprintf("%d",global.Global_Config_Manger.WebSocket.InterValReq))
		return
	}

	port, e := core.NewPort(userid, c.Writer, c.Request, nil)
	if e!=nil {
		fmt.Println("初始化连接失败",e)
		return
	}
	port.ConnectHubToWork()
}

//发送消息
func SendMsg(c *gin.Context)  {

	var m core.WSMessage
	query := c.ShouldBindJSON(&m)
	fmt.Print(m,"======")
	if query !=nil {
		c.JSON(200,"shoubindquery异常")
		return
	}
	if len(m.FromId) == 0 {
		c.JSON(200,"formId等于0")
		return
	}
	if _,ok:=core.HubHandle().GetPort(m.FromId);!ok{
		c.JSON(200,"您视乎断开连接了")
		return
	}
	m.ToidArray= strings.Split(m.ToId, ",")
	if len(m.ToidArray)==0 {
		c.JSON(200,"接收人为空")
	}
	global.Global_DealWsMsg_Pool.Invoke(m)
	c.JSON(200,"发送成功")
}