package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/liuhangkaixcode/websocket/core"
    "github.com/liuhangkaixcode/websocket/ginserver"
    "github.com/liuhangkaixcode/websocket/global"
    "github.com/liuhangkaixcode/websocket/initprog"
    "net/http"
    "time"
)

func main() {
    defer func() {
        global.Global_RedisPoolInstance.Close() //redis释放
        global.Global_MysqlDbInstance.Close() //mysql释放
        global.Global_DealWsMsg_Pool.Release() //处理websocket msg 转发信息的pool释放
    }()

    //标准配置初始化
    initprog.InitBasicConfig()
    //初始化websocket配置
    core.InitWebSocketBasic()


    go func() {
        fmt.Println("====")
        for{
            time.Sleep(time.Second*5)
            fmt.Println("整个数据",core.HubHandle().GetAllportsMap())
        }
    }()

    //ginServer
    engine := gin.Default()
    ginserver.InitRouter(engine)
    s := &http.Server{
        Addr:           ":"+global.Global_Config_Manger.WebSocket.Port,
        Handler:        engine,
        ReadTimeout:    50 * time.Second,
        WriteTimeout:   50 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    fmt.Println("服务已经起来了.... 端口号是：",global.Global_Config_Manger.WebSocket.Port)



    core.LogDebug("debug信息")
    core.LogWarn("LogWarn信息")
    core.LogError("LogError信息")
    core.LogInfo("LogInfo信息")



    s.ListenAndServe()




}

