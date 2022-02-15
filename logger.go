package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var atomicLevel zap.AtomicLevel

func InitLogger(dirPath, logName string, maxSize, maxBackups, maxAge int, compress bool) *zap.SugaredLogger {
	createLogDirectory(dirPath)
	writeSyncer := getLogWriter(dirPath, logName, maxSize, maxBackups, maxAge, compress)
	syncer := zap.CombineWriteSyncers(os.Stdout, writeSyncer)
	encoder := getEncoder()
	atomicLevel = zap.NewAtomicLevel()
	core := zapcore.NewCore(encoder, syncer, atomicLevel)
	logger := zap.New(core, zap.AddCaller())
	/*
		mux := http.NewServeMux()
		mux.Handle("/log_level", atom)
		go http.ListenAndServe(":1065", mux)
	*/
	return logger.Sugar()
}

func SetLoggerLevel(lvl string) {
	var level zapcore.Level
	switch lvl {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		// Requested log level not valid
		return
	}
	atomicLevel.SetLevel(level)
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

func getLogWriter(dirPath, logName string, maxSize, maxBackups, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   dirPath + "/" + logName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

/*
func main() {
	logger := InitLogger("/ramdisk", "test", 1, 1, 1, true)
	defer logger.Sync()
	var lvl int
	for {
		logger.Info("logger Info")
		logger.Debug("logger Debug")
		logger.Warn("logger Warn")
		if lvl == 0 {
			lvl += 1
			SetLoggerLevel("debug")
		} else {
			lvl = 0
			SetLoggerLevel("warn")
		}
		time.Sleep(1000 * time.Millisecond)
	}
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
