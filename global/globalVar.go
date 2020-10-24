package global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Global_Config_Manger  ConfigManger //配置管理
	Global_viperInstance   *viper.Viper  //配置管理实例
	Global_LoggerInstance  *zap.Logger   //日志实例
	Global_RedisPoolInstance *redis.Pool //redispool
	Global_MysqlDbInstance  *gorm.DB     //mysql

	Global_DealWsMsg_Pool  *ants.PoolWithFunc //处理websocket msg 转发信息的pool
)
