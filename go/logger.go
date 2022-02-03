package 4seils_utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(dirPath, logName string) *zap.SugaredLogger {
	createLogDirectory(dirPath)
	writeSyncer := getLogWriter(dirPath, logName)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func createLogDirectory(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		_ = os.Mkdir(dirPath, os.ModePerm)
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(dirPath, logName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   dirPath + logName,
		MaxSize:    5,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

/*
func main() {
	logger := InitLogger("/ramdisk", "test")
	defer logger.Sync()
	logger.Info("logger initialized")
	logger.Debug("logging debug")
}
*/

/*
func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
func main() {
	hook := lumberjack.Logger{
		Filename:   "./logs/spikeProxy1.log", // log file path
		MaxSize:    5,                        // The maximum size of each log file saved Unit: M
		MaxBackups: 20,                       // How many backups the log file can save at most
		MaxAge:     21,                       // How many days the file is kept at most
		Compress:   true,                     // whether to compress
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // lowercase encoder
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC time format
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // full path encoder
		EncodeName:     zapcore.FullNameEncoder,
	}

	// set log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // encoder configuration
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // print to console and file
		atomicLevel, // log level
	)
	// open development mode, stack trace
	caller := zap.AddCaller()
	// open file and line number
	development := zap.Development()
	// set initialization fields
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// construct log
	logger := zap.New(core, caller, development, filed)

	logger.Info("log initialization succeeded")
	logger.Info("Unable to get URL",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}
*/
