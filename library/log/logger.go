package log

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerCfg struct {
	Path   string `yaml:"path" json:"path"`
	Stdout bool   `yaml:"stdout" json:"stdout"`
	Access bool   `yaml:"access" json:"access"`
	Root   *Root  `yaml:"root" json:"root"`
}

type Root struct {
	FileName   string `yaml:"file_name" json:"file_name"`
	MaxSize    int    `yaml:"max_size" json:"max_size"`
	MaxAge     int    `yaml:"max_age" json:"max_age"`
	MaxBackups int    `yaml:"max_backups" json:"max_backups"`
	Compress   bool   `yaml:"compress" json:"compress"`
	Level      string `yaml:"level" json:"level"`
}

const (
	accessLogName       = "access.log"
	accessLogMaxSize    = 1
	accessLogMaxAge     = 1
	accessLogMaxBackups = 3
	accessLogCompress   = false
	accessLogLevel      = zapcore.InfoLevel
)

var (
	rootLogger   *zap.SugaredLogger
	accessLogger *zap.SugaredLogger
)

func Init(cfg *LoggerCfg) {
	// mkdir
	if cfg.Path != "" {
		_, err := os.Stat(cfg.Path)
		if err != nil && !os.IsExist(err) {
			if err = os.MkdirAll(cfg.Path, os.ModePerm); err != nil {
				panic(err)
			}
		}
	}

	// encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "ts"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// access
	aw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(cfg.Path, accessLogName),
		MaxSize:    accessLogMaxSize,
		MaxAge:     accessLogMaxAge,
		MaxBackups: accessLogMaxBackups,
		Compress:   accessLogCompress,
	})
	if cfg.Stdout {
		aw = zapcore.NewMultiWriteSyncer(aw, os.Stdout)
	}
	ac := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		aw,
		accessLogLevel,
	)
	accessLogger = zap.New(ac).Sugar()

	// root
	if cfg.Root != nil {
		rw := zapcore.AddSync(&lumberjack.Logger{
			Filename:   filepath.Join(cfg.Path, cfg.Root.FileName),
			MaxSize:    cfg.Root.MaxSize,
			MaxAge:     cfg.Root.MaxAge,
			MaxBackups: cfg.Root.MaxBackups,
			Compress:   cfg.Root.Compress,
		})
		if cfg.Stdout {
			rw = zapcore.NewMultiWriteSyncer(rw, os.Stdout)
		}
		rc := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			rw,
			getLevel(cfg.Root.Level),
		)
		rootLogger = zap.New(rc).Sugar()
	}
}

func Access(args ...interface{}) {
	if accessLogger != nil {
		accessLogger.Info(args)
	}
}

func Accessf(template string, args ...interface{}) {
	if accessLogger != nil {
		accessLogger.Infof(template, args...)
	}
}

func Error(args ...interface{}) {
	if rootLogger != nil {
		rootLogger.Error(args)
	}
}

func Errorf(template string, args ...interface{}) {
	if rootLogger != nil {
		rootLogger.Errorf(template, args...)
	}
}

func Info(args ...interface{}) {
	if rootLogger != nil {
		rootLogger.Info(args)
	}
}

func Infof(template string, args ...interface{}) {
	if rootLogger != nil {
		rootLogger.Infof(template, args...)
	}
}

func Debug(args ...interface{}) {
	if rootLogger != nil {
		rootLogger.Debug(args)
	}
}

func Debugf(template string, args ...interface{}) {
	if rootLogger != nil {
		rootLogger.Debugf(template, args...)
	}
}

func getLevel(level string) zapcore.Level {
	var logLevel zapcore.Level
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}
	return logLevel
}
