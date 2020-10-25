package core
//消息结构体
type WSMessage struct {
	FromId string    `form:"fromid" json:"fromid"`  //请求userid
	ToId  string   `form:"toid" json:"toid" `       //发送id 12,33,33 用，隔开
	Type string     `form:"type" json:"type" `      //消息的类型 0发送消息  1断开
	Content string   `form:"content" json:"content" `  //消息内容
	OtherInfo string `form:"otherinfo" json:"otherinfo" `  //消息额外的消息
	ToidArray []string `form:"-",json:"-"`

}
