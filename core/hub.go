package core

import (
	"fmt"
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
	GetPort(userId string) IPort
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

func (h *hub) AddCache(userId string,m string) {

}

func (h *hub) GetCache(userId string)[]string {
	return nil
}


func (h *hub)AddPort(userId string ,conn IPort) error {
	h.locker.Lock()
	defer h.locker.Unlock()
	if v,ok:= h.portsMap[userId];ok {
		v.Close()
		h.portsMap[userId]=conn
		return nil
	}

	if len(h.portsMap)+1 >= global.Global_Config_Manger.WebSocket.MaxClient{
		fmt.Println("===已经达到数量了")
		return fmt.Errorf("totalports two big")
	}
	h.portsMap[userId]=conn
	return nil

}

func (h *hub)RemovePort(userId string)  {
	h.locker.Lock()
	defer h.locker.Unlock()
	delete(h.portsMap,userId)

}

func (h *hub)GetPort(userId string) IPort {
	h.locker.RLock()
	defer h.locker.RUnlock()
	return h.portsMap[userId]
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

