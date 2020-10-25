package ginserver

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhangkaixcode/websocket/global"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"sort"
)

//userid校验
func checkUserId(c *gin.Context)  {
	userid:=c.Query("userid")
	if len(userid) == 0 {
		c.Abort()
		c.JSON(200,gin.H{"code":200,"message":"userid没传或者为空"})
	}else{
		c.Next()
	}
}

//sign签名校验
func checkSign(c *gin.Context)  {
	//name:["df"] kai:[]
	c.Request.ParseForm()
	map1:=c.Request.Form
	if len(map1["sign"]) == 0  || map1["sign"][0] == ""{
		c.JSON(200,"sign没有这个字段")
		c.Abort()
		return
	}
	sign:=map1["sign"][0]
	delete(map1,"sign")
	var keys []string
	for k,v:=range map1{
		if len(v)==0 || v[0] == ""{
			continue
		}
		fmt.Println(k,v,len(v))
		keys =append(keys,k)
	}
	sort.Strings(keys)
	result:=""
	if len(keys) == 0 {

	}else {
		for i:=0;i<len(keys);i++{
			result=fmt.Sprintf("%s&%s=%s",result,keys[i],map1[keys[i]][0])
		}
	}

	if result !="" {
		result=result[1:]
	}
	h := md5.New()
	h.Write([]byte(result))
	md5result:= hex.EncodeToString(h.Sum(nil))

	fmt.Printf("最终结果的字符=> %s  md5==> %s sign=> %s \n",result,md5result,sign)
	if md5result == sign{
		c.Next()

	}else{
		c.JSON(200,"sign验签名不通过")
		c.Abort()

	}



}
//链路检测
func jaegerCheck(ctx *gin.Context)  {
	path := ctx.Request.URL.Path

	span := global.Global_Jaeger.StartSpan(path,

		ext.SpanKindRPCServer)
	ext.HTTPUrl.Set(span, path)
	ext.HTTPMethod.Set(span, ctx.Request.Method)
	c := opentracing.ContextWithSpan(context.Background(), span)

	ctx.Set("ctx", c)

	ctx.Next()

	ext.HTTPStatusCode.Set(span, uint16(ctx.Writer.Status()))
	span.Finish()
}
