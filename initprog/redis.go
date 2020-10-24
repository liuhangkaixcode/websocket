package initprog

import (
	"fmt"
	"github.com/liuhangkaixcode/websocket/global"
	"time"
    "github.com/garyburd/redigo/redis"
)
func initRedis() {
	pool := &redis.Pool{
		MaxIdle:     global.Global_Config_Manger.Redis.MaxIdle,
		MaxActive:   global.Global_Config_Manger.Redis.MaxActive,
		IdleTimeout: time.Duration(global.Global_Config_Manger.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", global.Global_Config_Manger.Redis.HostAndPort,redis.DialPassword(global.Global_Config_Manger.Redis.PassWord),redis.DialDatabase(global.Global_Config_Manger.Redis.DB))
		},

	}
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("ping")
	if err != nil {
		panic("redis初始化失败"+fmt.Sprint("%v",err))
	}else{
		fmt.Println("redis已经起来了")
	}
	global.Global_RedisPoolInstance=pool


}
