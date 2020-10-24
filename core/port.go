package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/liuhangkaixcode/websocket/global"
	"net/http"
	"sync"
	"time"
)
const (
	writeWait = 2 * time.Second
	pongWait = 11 * time.Second
	pingPeriod = 10 *time.Second
)

var(
	upgrader = websocket.Upgrader{
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

//func GetWebSocketConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
//  return upgrader.Upgrade(w,r,responseHeader)
//}

type opFunc  func(h *Port)
type Port struct {
	userId string
	conn *websocket.Conn
	isAdd bool
	l   sync.Mutex
}

type IPort interface {
	ConnectHubToWork()
	SendMsg(m string) error
	Close() chan int
	GetConn() *websocket.Conn
	GetUserID() string
}

//func WithUserIdOption(userId string) opFunc {
//	return func(h *Hub) {
//		h.userId=userId
//	}
//}

func NewPort(userId string,w http.ResponseWriter, r *http.Request, responseHeader http.Header,ops ...opFunc) (IPort ,error){
	conn, e := upgrader.Upgrade(w, r, responseHeader)
	if e!=nil {
		return nil,e
	}
	a:=new(Port)
	a.conn=conn
	a.userId=userId
	for _, op:=range ops{
		op(a)
	}

	if HubHandle().GetAllportslens()+1 >= global.Global_Config_Manger.WebSocket.MaxClient{
		conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf("连接数已经达到最大数-%d",global.Global_Config_Manger.WebSocket.MaxClient)))
		conn.Close()
		return nil,fmt.Errorf("已经超过了")
	}


	if v,ok:= HubHandle().GetPort(userId);ok{
		select {
		case <-v.Close():
			e = HubHandle().AddPort(userId, a)
		case <-time.After(time.Second*5):
			return nil,fmt.Errorf("超时了")

		}
	}else{
		e = HubHandle().AddPort(userId, a)
	}



  data:=HubHandle().GetCache(a.userId)
  for _,d:=range data{
  	a.GetConn().WriteMessage(websocket.TextMessage,[]byte(d))
  }



	return a,nil

}

func (p *Port)Close() chan int{
	ch:=make(chan int,1)
	 p.l.Lock()
	 defer p.l.Unlock()
	if p.conn==nil {

	}else{
		p.conn.Close()
		p.conn=nil
		HubHandle().RemovePort(p.userId)

	}
	ch<-1
	return ch

}

func (p *Port)SendMsg(m string)  error{
	p.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := p.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
		fmt.Println("writeMessage::error",err)
		return err
	}
	return nil
}
func (p *Port) ConnectHubToWork(){
	ch:=make(chan int ,1)
	go p.readerMessage(ch)
	go p.pingWrite11(ch)
}

func (p *Port)GetConn() *websocket.Conn{
	return p.conn
}
func (p *Port)GetUserID() string{
	return p.userId
}
func (p*Port)readerMessage(closech chan int) {
	defer func() {
		fmt.Println("读的时候-关闭连接了")
	}()

	p.conn.SetReadLimit(512)
	p.conn.SetReadDeadline(time.Now().Add(pongWait))
	p.conn.SetPongHandler(func(t string) (ee error ){
		fmt.Println("pong callBack-->",t,ee,time.Now().String())
		p.conn.SetReadDeadline(time.Now().Add(pongWait))
		return
	})
	for {
		messageType, _, err := p.conn.ReadMessage()
		//这里是总控 无论是ping 还是主动发送信息 都是在这里
		fmt.Println("读取所有的信息-->",messageType,err)
		if err != nil || messageType == websocket.CloseMessage {
			<-p.Close()
			close(closech)
			return
		}
		//这个地方发送内容
		//p=append(p,[]byte("|| server")...)
		//ws.SetWriteDeadline(time.Now().Add(writeWait))
		//if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
		//	fmt.Println("writeMessage::error",err)
		//	return
		//}

	}
}
func (p *Port)pingWrite11(closech chan int)  {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		fmt.Println("因为ping关闭了")

	}()

	for{
		select {
		case <-pingTicker.C:
			{
				p.conn.SetWriteDeadline(time.Now().Add(writeWait))
				control := p.conn.WriteControl(websocket.PingMessage, []byte("a"), time.Now().Add(writeWait))
				if control !=nil{
					return
				}
			}
		case <-closech:
			return


		}
	}
}
