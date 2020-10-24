package core

//业务集中处理data
func dealWithMessage(m WSMessage)  {
	if m.Type == "1" {
		for _,v:=range m.ToidArray{
			kk:=HubHandle().GetPort(v)
			if kk == nil {
				HubHandle().AddCache(v,m.Content)
			}else{
				kk.SendMsg(m.Content)
			}

		}
	}else if m.Type == "0" {
		HubHandle().GetPort(m.FromId).Close()
		HubHandle().RemovePort(m.FromId)
	}else {
		HubHandle().GetPort(m.FromId).SendMsg(m.Content)

	}

}
