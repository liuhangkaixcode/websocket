package core

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/liuhangkaixcode/websocket/global"
	"sync"
)

var (
	once sync.Once
	instance *hub

)

type hub struct {
	portsMap map[string]IPort
	locker  sync.RWMutex

}
type hubIF interface {
	AddPort(userId string ,conn IPort)error
	RemovePort(userId string)
	GetPort(userId string) (IPort, bool)
	GetAllportsMap() map[string]IPort
	GetAllportslens() int

	AddCache(userId string,m string)
	GetCache(userId string)[]string

}



func (h *hub)GetAllportslens()int  {
	h.locker.RLock()
	defer h.locker.RUnlock()
	return len(h.portsMap)
}


func (h *hub)GetAllportsMap()map[string]IPort  {
	h.locker.RLock()
	defer h.locker.RUnlock()
	return h.portsMap
}

//设置缓存
func (h *hub) AddCache(userId string,m string) {
     conn:=global.Global_RedisPoolInstance.Get()
     defer conn.Close()
     conn.Do("lpush",userId,m)
}

func (h *hub) GetCache(userId string)[]string {
	var data []string
	conn:=global.Global_RedisPoolInstance.Get()
	defer conn.Close()
	id2, err := redis.Values(conn.Do("LRANGE", userId,0,-1))
	if err != nil {
		fmt.Print("LRANGE=====",err)
		return nil
	}
	for _, v:=range id2{
		data=append(data,string(v.([]byte)))
	}
	conn.Do("del",userId)
	return data
}


func (h *hub)AddPort(userId string ,conn IPort) error {
	h.locker.Lock()
	defer h.locker.Unlock()

	h.portsMap[userId]=conn
	return nil

}

func (h *hub)RemovePort(userId string)  {
	h.locker.Lock()
	defer h.locker.Unlock()
	delete(h.portsMap,userId)

}

func (h *hub)GetPort(userId string) (IPort, bool){
	h.locker.RLock()
	defer h.locker.RUnlock()
	v,ok:=h.portsMap[userId]
	return v,ok
}
//获取单列
func HubHandle() hubIF {
	//单列
	once.Do(func() {
		fmt.Println("单列....")
		instance =&hub{portsMap: make(map[string]IPort,1000)}
	})
	return instance
}

