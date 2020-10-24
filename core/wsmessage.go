package core
type WSMessage struct {
	FromId string    `form:"fromid" json:"formid"`
	ToId  string   `form:"toid" json:"toid"`
	Type string     `form:"type" json:"type"`
	Content string   `form:"content" json:"content"`
	OtherInfo string `form:"otherinfo" json:"otherinfo"`
	ToidArray []string `form:"-",json:"-"`

}
