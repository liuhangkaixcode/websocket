package initprog

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/liuhangkaixcode/websocket/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"

)

func InitBasicConfig()  {
	//初始化配置
    initConfig()
    //初始化日志配置
    initLogConfig()
    //初始化sql
    initSql()
    //初始化redis
    initRedis()
    //初始化jaeger
    initJaeger()


}

func initLogConfig()  {
	hook := lumberjack.Logger{
		Filename:   global.Global_Config_Manger.Log.LocalPath,  // 日志文件路径
		MaxSize:    500,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                       // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		//TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()

	switch global.Global_Config_Manger.Log.Level {
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	default:
		atomicLevel.SetLevel(zap.DebugLevel)
	}
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到控制台和文件zapcore.AddSync(os.Stdout),
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	//caller := zap.AddCaller()
	// 开启文件及行号
	//development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	//logger := zap.New(core, caller, development, filed)

	logger := zap.New(core)
	//logger.Info("log 初始化成功")
	//logger.Info("无法获取网址",
	//	zap.String("url", "http://www.baidu.com"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second))
	logger.Info("日志初始化配置成功",zap.String("logwirteTime",time.Now().Format("2006-01-02 15:04:05")))
	global.Global_LoggerInstance=logger


}

func initConfig()  {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.Global_Config_Manger); err != nil {
			fmt.Println("配置出错",err)
		}
	})
	if err := v.Unmarshal(&global.Global_Config_Manger); err != nil {
		fmt.Println("配置出错",err)
	}
	global.Global_viperInstance=v
	marshal, _ := json.Marshal(global.Global_Config_Manger)
	fmt.Println("所有配置的信息",string(marshal))

	//测试重新设置的问题
	//config.Global_V.Set("mysql",map[string]string{"host":"yyyy","port":"yyyyyyyyy-"})
	////config.Global_V.Set("redis",map[string]string{"host":"dfdf1","port":"port-"})
	//e := config.Global_V.WriteConfig()
	//fmt.Println(e)


}

