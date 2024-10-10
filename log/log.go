package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(args ...any)
	Infoln(args ...any)
	Infof(template string, args ...any)
	Infow(msg string, keysAndValues ...any)
	Error(args ...any)
	Errorln(args ...any)
	Errorf(template string, args ...any)
	Errorw(msg string, keysAndValues ...any)
	Fatal(args ...any)
	Fatalln(args ...any)
	Fatalf(template string, args ...any)
	Fatalw(msg string, keysAndValues ...any)
	Debug(args ...any)
	Debugln(args ...any)
	Debugf(template string, args ...any)
	Debugw(msg string, keysAndValues ...any)
	Warn(args ...any)
	Warnln(args ...any)
	Warnf(template string, args ...any)
	Warnw(msg string, keysAndValues ...any)
}

var _ logger.Logger = (*WailsLogger)(nil)
var _ Logger = (*GoLogger)(nil)
var levelMap = map[logger.LogLevel]zapcore.Level{
	logger.DEBUG:   zap.DebugLevel,
	logger.INFO:    zap.InfoLevel,
	logger.WARNING: zap.WarnLevel,
	logger.ERROR:   zap.ErrorLevel,
	logger.TRACE:   zap.DebugLevel,
}

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

var levelToColor = map[zapcore.Level]Color{
	zapcore.DebugLevel:  Magenta,
	zapcore.InfoLevel:   Blue,
	zapcore.WarnLevel:   Yellow,
	zapcore.ErrorLevel:  Red,
	zapcore.DPanicLevel: Red,
	zapcore.PanicLevel:  Red,
	zapcore.FatalLevel:  Red,
}

func NewWailsLogger(level logger.LogLevel) *WailsLogger {
	return &WailsLogger{NewLogger(level, "Wails", Red, 1)}
}
func NewGoLogger(level logger.LogLevel) *GoLogger {
	return &GoLogger{NewLogger(level, "Go   ", Blue, 0)}
}
func NewJSLogger(level logger.LogLevel) *JSLogger {
	return &JSLogger{NewLogger(level, "JS   ", Yellow, 1)}
}
func NewLogger(level logger.LogLevel, tag string, tagColor Color, callerSkip int) *zap.SugaredLogger {
	colorTag := tagColor.Add(tag)
	consoleEncoderCfg := zap.NewProductionEncoderConfig()
	consoleEncoderCfg.ConsoleSeparator = "	"
	consoleEncoderCfg.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		// FIXME: 用更好的方式显示 tag
		pae.AppendString(t.Format("2006-01-02 15:04:05.000") + consoleEncoderCfg.ConsoleSeparator + colorTag)
	}
	consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderCfg.EncodeDuration = zapcore.StringDurationEncoder
	consoleEncoderCfg.EncodeCaller = func(ec zapcore.EntryCaller, pae zapcore.PrimitiveArrayEncoder) {}
	consoleEncoderCfg.EncodeName = zapcore.FullNameEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderCfg)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), levelMap[level])

	fileEncoderCfg := zap.NewProductionEncoderConfig()
	fileEncoderCfg.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000") + consoleEncoderCfg.ConsoleSeparator + tag)
	}
	fileEncoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoderCfg.EncodeDuration = zapcore.StringDurationEncoder
	fileEncoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	fileEncoderCfg.EncodeName = zapcore.FullNameEncoder

	fileEncoder := zapcore.NewJSONEncoder(fileEncoderCfg)
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting user config directory:", err)
	} else {
		// 保留最近的10个日志文件
		files, _ := filepath.Glob(filepath.Join(cfgDir, "flplugman", "flplugman*.log"))
		if len(files) > 10 {
			for i := 0; i < len(files)-10; i++ {
				os.Remove(files[i])
			}
		}
	}
	logFileName := filepath.Join(cfgDir, "flplugman", "flplugman"+time.Now().Format("2006-01-02")+".log")
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}
	fileCore := zapcore.NewCore(fileEncoder, zapcore.Lock(logFile), levelMap[level])

	core := consoleCore
	if err == nil {
		core = zapcore.NewTee(consoleCore, fileCore)
	}
	logger := zap.New(
		core,
		zap.AddCallerSkip(callerSkip), zap.AddCaller(),
	).Sugar()
	return logger
}

type WailsLogger struct {
	SugaredLogger *zap.SugaredLogger
}

func (l *WailsLogger) Debug(message string) {
	l.SugaredLogger.Debug(message)
}
func (l *WailsLogger) Error(message string) {
	l.SugaredLogger.Error(message)
}
func (l *WailsLogger) Fatal(message string) {
	l.SugaredLogger.Fatal(message)
}
func (l *WailsLogger) Info(message string) {
	l.SugaredLogger.Info(message)
}
func (l *WailsLogger) Print(message string) {
	l.SugaredLogger.Info(message)
}
func (l *WailsLogger) Trace(message string) {
	l.Debug(message)
}
func (l *WailsLogger) Warning(message string) {
	l.SugaredLogger.Warn(message)
}

type GoLogger struct {
	*zap.SugaredLogger
}
type JSLogger struct {
	SugaredLogger *zap.SugaredLogger
}

func (j *JSLogger) Debug(message string) {
	j.SugaredLogger.Debug(message)
}
func (j *JSLogger) Error(message string) {
	j.SugaredLogger.Error(message)
}
func (j *JSLogger) Fatal(message string) {
	j.SugaredLogger.Fatal(message)
}
func (j *JSLogger) Info(message string) {
	j.SugaredLogger.Info(message)
}
func (j *JSLogger) Warn(message string) {
	j.SugaredLogger.Warn(message)
}
