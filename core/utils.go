package core

import (
	"encoding/json"
	"fmt"
	"github.com/liuhangkaixcode/websocket/global"
	"github.com/liuhangkaixcode/websocket/initprog"
	"go.uber.org/zap"
	"time"
)
//日志处理
func LogWarn(msg string,ops...LogOption)  {
	if global.Global_Config_Manger.Log.Mode == 1 {
		obj:=getinfoMessageStruct(msg,ops...)
		global.Global_LoggerInstance.Warn(msg,zap.String("status",obj.Status,),zap.String("serviceName",obj.ServiceName),zap.String("reqName",obj.ReqName),zap.String("logTime",obj.LogTime))
	}else{
         s:=getinfoMessageString(msg,"warn",ops...)
         sendRabbitMsg(s)
	}
}

func LogDebug(msg string,ops...LogOption)  {
	if global.Global_Config_Manger.Log.Mode == 1 {
		obj:=getinfoMessageStruct(msg,ops...)
		global.Global_LoggerInstance.Debug(msg,zap.String("status",obj.Status,),zap.String("serviceName",obj.ServiceName),zap.String("reqName",obj.ReqName),zap.String("logTime",obj.LogTime))
	}else{
		s:=getinfoMessageString(msg,"debug",ops...)
		sendRabbitMsg(s)
	}
}

func LogInfo(msg string,ops...LogOption)  {
	if global.Global_Config_Manger.Log.Mode == 1 {
		obj:=getinfoMessageStruct(msg,ops...)
		global.Global_LoggerInstance.Info(msg,zap.String("status",obj.Status,),zap.String("serviceName",obj.ServiceName),zap.String("reqName",obj.ReqName),zap.String("logTime",obj.LogTime))
	}else{
		s:=getinfoMessageString(msg,"info",ops...)
		sendRabbitMsg(s)
	}
}

func LogError(msg string,ops...LogOption)  {
	if global.Global_Config_Manger.Log.Mode == 1 {
		obj:=getinfoMessageStruct(msg,ops...)
		global.Global_LoggerInstance.Error(msg,zap.String("status",obj.Status,),zap.String("serviceName",obj.ServiceName),zap.String("reqName",obj.ReqName),zap.String("logTime",obj.LogTime))
	}else{
		s:=getinfoMessageString(msg,"error",ops...,)
		sendRabbitMsg(s)
	}
}


type LogOption func(l *log)
type log struct {
	Status string `json:"status"`
	ServiceName string `json:"serviceName"`
	ReqName string `json:"reqName"`
	Msg string  `json:"msg"`
	LogTime string  `json:"logTime"`
	Level string `json:level`

}

func WithStaus(s string) LogOption  {
	return func(l *log) {
		l.Status=s
	}
}



func WithServiceName(s string) LogOption  {
	return func(l *log) {
		l.ServiceName=s
	}
}
func WithReqName(s string)LogOption  {
	return func(l *log) {
		l.ReqName=s
	}
}

func getinfoMessageString(msg string,level string,ops...LogOption)  string{
	log1:=new(log)
	for _,op:=range ops{
		op(log1)
	}
	log1.Msg=msg
	log1.Level=level
	log1.LogTime=time.Now().Format("2006-01-02 15:04:05")
	bytes, _ := json.Marshal(log1)
	return string(bytes)
}

func getinfoMessageStruct(msg string,ops...LogOption)  *log{
	log1:=new(log)
	for _,op:=range ops{
		op(log1)
	}
	log1.Msg=msg
	log1.LogTime=time.Now().Format("2006-01-02 15:04:05")

	return log1
}

func sendRabbitMsg(msg string)  {
	var  MQURL string = fmt.Sprintf("amqp://%s:%s@%s/%s",global.Global_Config_Manger.RabbitMQ.UserName,global.Global_Config_Manger.RabbitMQ.Password,global.Global_Config_Manger.RabbitMQ.Host,global.Global_Config_Manger.RabbitMQ.Vhost)
	rabbit:=initprog.NewRabbitMq(global.Global_Config_Manger.RabbitMQ.QueueName,MQURL)
	rabbit.PublishSimple(string(msg))
}
