# websocket
```
1.建立连接
ws://ip:8181/connect?userid=98

2.发送消息
http://ip:8181/sendmsg
content-type: appliction/json
{
  "fromid":"998", //发送者
  "toid":"98,998", //接受者
  "type":"1",     //消息类型
  "content":"12234354354354354524", //消息内容
  "otherinfo":"otherinfo"   //额外信息
  
}
**备注 发送的消息统一由线程池处理
3.日志收集 Rabbit+ELK 或者uber zap写本地
4.链路检测 jaeger

```


