package ginserver

import (
	"github.com/liuhangkaixcode/websocket/global"
	"github.com/muesli/cache2go"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

type Response struct {
	Code int
	Msg string
	Data map[string]interface{}
}
type Request struct {
	LastRequestTime  int64
	UserId string
	IsDeadTime bool
	ReqChan chan *Response


}
var (
	RequestLocker   sync.RWMutex
	RequestMap=make(map[string]Request,10000)
	CacheTool= cache2go.Cache("myCache")
	P *ants.PoolWithFunc
	ReqChan =make(chan *Request,1000)
)



//请求是否过快
func SetCache(userId string,r *Request)  {
	CacheTool.Add(userId,time.Second*20,r)
}
func GetCache(userId string) *Request  {
	item, e := CacheTool.Value(userId)

	if e==nil{
		return item.Data().(*Request)
	}

	return nil
}
func DeleteCache(userId string)  {
	CacheTool.Delete(userId)
}

//检查请求频率是否过快
func checkRequestIsFast(userId string) (bool ,*Request) {
	requst:=GetCache(userId)
	if requst == nil {
		r:=&Request{
			LastRequestTime:time.Now().Unix(),
		}
		SetCache(userId,r)
		return true,r
	}else{

		if time.Now().Unix()-requst.LastRequestTime<int64(global.Global_Config_Manger.WebSocket.InterValReq) {
			return false,requst
		}
		requst.LastRequestTime=time.Now().Unix()
		SetCache(userId,requst)
		return true,requst

	}
	return true,requst
}


