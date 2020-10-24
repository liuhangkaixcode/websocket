package initprog

import (
	"fmt"
	"github.com/jinzhu/gorm"
	 _"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liuhangkaixcode/websocket/global"
)

func initSql(){
	sqlstring:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",global.Global_Config_Manger.Mysql.UserName,global.Global_Config_Manger.Mysql.PassWord,global.Global_Config_Manger.Mysql.Host,global.Global_Config_Manger.Mysql.Port,global.Global_Config_Manger.Mysql.DB)
	db,err:=gorm.Open("mysql",sqlstring)
	if err!=nil {
		panic(fmt.Sprintf("mysql初始化错误%v",err))
	}else{
		fmt.Println("init mysql success")
	}

	db.SingularTable(true) //表不会有复数
	// 启用Logger，显示详细日志
	db.LogMode(false)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(

	)

	global.Global_MysqlDbInstance=db

}
