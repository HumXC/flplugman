package log

import (
	"os"

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

func NewLogger(level logger.LogLevel) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), levelMap[level])
	logger := zap.New(core, zap.AddCallerSkip(1)).Sugar()
	return logger
}

type WailsLogger struct {
	SugaredLogger *zap.SugaredLogger
}

func (l *WailsLogger) Debug(message string) {
	l.SugaredLogger.Debug("Wails | ", message)
}
func (l *WailsLogger) Error(message string) {
	l.SugaredLogger.Error("Wails | ", message)
}
func (l *WailsLogger) Fatal(message string) {
	l.SugaredLogger.Fatal("Wails | ", message)
}
func (l *WailsLogger) Info(message string) {
	l.SugaredLogger.Info("Wails | ", message)
}
func (l *WailsLogger) Print(message string) {
	l.SugaredLogger.Info("Wails | ", message)
}
func (l *WailsLogger) Trace(message string) {
	l.Debug(message)
}
func (l *WailsLogger) Warning(message string) {
	l.SugaredLogger.Warn("Wails | ", message)
}

type GoLogger struct {
	SugaredLogger *zap.SugaredLogger
}

func (g *GoLogger) Debug(args ...any) {
	g.SugaredLogger.Debugw("   Go | ", args...)
}
func (g *GoLogger) Debugf(template string, args ...any) {
	g.SugaredLogger.Debugf("   Go | "+template, args...)
}
func (g *GoLogger) Debugln(args ...any) {
	g.SugaredLogger.Debugln(append([]any{"   Go |"}, args...)...)
}
func (g *GoLogger) Debugw(msg string, keysAndValues ...any) {
	g.SugaredLogger.Debugw("   Go | "+msg, keysAndValues...)
}
func (g *GoLogger) Error(args ...any) {
	g.SugaredLogger.Errorw("   Go | ", args...)
}
func (g *GoLogger) Errorf(template string, args ...any) {
	g.SugaredLogger.Errorf("   Go | "+template, args...)
}
func (g *GoLogger) Errorln(args ...any) {
	g.SugaredLogger.Errorln(append([]any{"   Go |"}, args...)...)
}
func (g *GoLogger) Errorw(msg string, keysAndValues ...any) {
	g.SugaredLogger.Errorw("   Go | "+msg, keysAndValues...)
}
func (g *GoLogger) Fatal(args ...any) {
	g.SugaredLogger.Fatalw("   Go | ", args...)
}
func (g *GoLogger) Fatalf(template string, args ...any) {
	g.SugaredLogger.Fatalf("   Go | "+template, args...)
}
func (g *GoLogger) Fatalln(args ...any) {
	g.SugaredLogger.Fatalln(append([]any{"   Go |"}, args...)...)
}
func (g *GoLogger) Fatalw(msg string, keysAndValues ...any) {
	g.SugaredLogger.Fatalw("   Go | "+msg, keysAndValues...)
}
func (g *GoLogger) Info(args ...any) {
	g.SugaredLogger.Infow("   Go | ", args...)
}
func (g *GoLogger) Infof(template string, args ...any) {
	g.SugaredLogger.Infof("   Go | "+template, args...)
}
func (g *GoLogger) Infoln(args ...any) {
	g.SugaredLogger.Infoln(append([]any{"   Go |"}, args...)...)
}
func (g *GoLogger) Infow(msg string, keysAndValues ...any) {
	g.SugaredLogger.Infow("   Go | "+msg, keysAndValues...)
}
func (g *GoLogger) Warn(args ...any) {
	g.SugaredLogger.Warnw("   Go | ", args...)
}
func (g *GoLogger) Warnf(template string, args ...any) {
	g.SugaredLogger.Warnf("   Go | "+template, args...)
}
func (g *GoLogger) Warnln(args ...any) {
	g.SugaredLogger.Warnln(append([]any{"   Go |"}, args...)...)
}
func (g *GoLogger) Warnw(msg string, keysAndValues ...any) {
	g.SugaredLogger.Warnw("   Go | "+msg, keysAndValues...)
}

type JSLogger struct {
	SugaredLogger *zap.SugaredLogger
}

func (j *JSLogger) Debug(message string) {
	j.SugaredLogger.Debug("   JS | " + message)
}
func (j *JSLogger) Error(message string) {
	j.SugaredLogger.Error("   JS | ", message)
}
func (j *JSLogger) Fatal(message string) {
	j.SugaredLogger.Fatal("   JS | ", message)
}
func (j *JSLogger) Info(message string) {
	j.SugaredLogger.Info("   JS | ", message)
}
func (j *JSLogger) Warn(message string) {
	j.SugaredLogger.Warn("   JS | ", message)
}
