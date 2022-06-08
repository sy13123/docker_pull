package logtool

import (
	//"aopm-agnet/pkg/constant"
	//"github.com/go-ini/ini"
	//"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var SugLog *zap.SugaredLogger
var Logc *zap.Logger
func InitEvent() {
	//创建核心对象
	var coreArr []zapcore.Core
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig() //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	//encoderConfig.CallerKey=""
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	//配置日志级别
	level := getConfigLog()
	//info和debug级别,debug级别是最低的
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= level
	})
	//error级别
	//highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
	//	return lev >= zap.ErrorLevel
	//})
	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	//infoFileWriteSyncer := getInfoFileWriter()
	//errorFileWriteSyncer := getErrorFileWriter()
	//info文件writeSyncer
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer( zapcore.AddSync(os.Stdout)), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//error文件writeSyncer
	//errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//处理
	coreArr = append(coreArr, infoFileCore)
	//coreArr = append(coreArr, errorFileCore)
	//zap.AddCaller()为显示文件名和行号，可省略
	//log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(),zap.AddCallerSkip(1))
	log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
    //infoLog :=log.WithOptions(zap.AddCallerSkip(1))
	//获取
	SugLog = log.Sugar()
	Logc = log.WithOptions(zap.AddCallerSkip(1))
	//日志
	SugLog.Infof("**********日志初始化完成 输出级别=[%v]**********", level)
}

//格式获取当前日志级别
func getConfigLog() (level zapcore.Level) {
	//默认日志级别设置
	levelStr := "INFO"
	//读取配置获取日志输出的级别(直接读取配置文件)
	//cfg, err := ini.Load(constant.ConfigUrl)
	////如果配置文件存在有效
	//if err == nil {
	//	//获取日志级别
	//	configLevelStr := cfg.Section("").Key("log_level").String()
	//	//如果配置有效
	//	if configLevelStr != "" {
	//		//获取配置
	//		levelStr = configLevelStr
	//	}
	//}
	//默认日志级别
	level, _ = zapcore.ParseLevel(levelStr)
	//返回
	return level
}

