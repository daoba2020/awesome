package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Console = "console"
	File    = "file"
)

var (
	Leavel = zap.DebugLevel
	Target = Console
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}

func Init() {
	// 配置 sugaredLogger
	writer := zapcore.AddSync(os.Stdout)
	encoder := zapcore.NewConsoleEncoder(NewEncoderConfig()) // 设置 console 编码器
	core := zapcore.NewCore(encoder, writer, Leavel)
	Logger = zap.New(core, zap.AddCaller())

	Sugar = Logger.Sugar()
}
