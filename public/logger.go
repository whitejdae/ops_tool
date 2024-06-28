package public

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"ops_tool/conf"
	"os"
	"time"
)

var Log *zap.SugaredLogger

func InitLogger() {
	// 日志写入配置
	writerSyncer := getLogWriter()
	// 写入字段配置
	encoder := getEncoder()
	// 第三个参数为日志级别默认debug
	logger := zap.New(zapcore.NewCore(encoder, writerSyncer, zap.DebugLevel), zap.AddCaller())
	Log = logger.Sugar()
	zap.ReplaceGlobals(Log.Desugar()) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	Log.Info("初始化zap日志完成!")
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			TimeKey:     "time",
			CallerKey:   "file",
			FunctionKey: "func",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.Conf.System.Log.LogPath,    //日志文件存放目录
		MaxSize:    conf.Conf.System.Log.MaxSize,    //文件大小限制,单位MB
		MaxBackups: conf.Conf.System.Log.MaxBackups, //日志文件保留天数
		MaxAge:     conf.Conf.System.Log.MaxAge,     //最大保留日志文件数量
		LocalTime:  false,
		Compress:   conf.Conf.System.Log.Compress, //是否压缩处理
	})
	// 打印日志到文件和终端
	ws := zapcore.NewMultiWriteSyncer(lumberJackLogger, zapcore.AddSync(os.Stdout))
	return ws
}
