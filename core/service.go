package core

//业务集中处理
func dealWithMessage(m WSMessage)  {
	if m.Type == "1" {
		for _,v:=range m.ToidArray{
			iPort, _ := HubHandle().GetPort(v)

			if iPort == nil {
				HubHandle().AddCache(v,m.Content)
			}else{
				iPort.SendMsg(m.Content)
			}

		}
	}else if m.Type == "0" {
		iPort, _ := HubHandle().GetPort(m.FromId)
		iPort.Close()

		


	}else {

	}

}
