package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
	Sync() error
	WithField(key string, value any) Logger
}

type logger struct {
	logger *zap.Logger
}

func NewLogger() (Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	loggerConfig, _ := config.Build()
	return &logger{logger: loggerConfig}, nil
}

func (l *logger) Debug(args ...any) {
	l.logger.Sugar().Debug(args...)
}

func (l *logger) Debugf(format string, args ...any) {
	l.logger.Sugar().Debugf(format, args...)
}

func (l *logger) Info(args ...any) {
	l.logger.Sugar().Info(args...)
}

func (l *logger) Infof(format string, args ...any) {
	l.logger.Sugar().Infof(format, args...)
}

func (l *logger) Warn(args ...any) {
	l.logger.Sugar().Warn(args...)
}

func (l *logger) Warnf(format string, args ...any) {
	l.logger.Sugar().Warnf(format, args...)
}

func (l *logger) Error(args ...any) {
	l.logger.Sugar().Error(args...)
}

func (l *logger) Errorf(format string, args ...any) {
	l.logger.Sugar().Errorf(format, args...)
}

func (l *logger) Fatal(args ...any) {
	l.logger.Sugar().Fatal(args...)
}

func (l *logger) Fatalf(format string, args ...any) {
	l.logger.Sugar().Fatalf(format, args...)
}

func (l *logger) Panic(args ...any) {
	l.logger.Sugar().Panic(args...)
}

func (l *logger) Panicf(format string, args ...any) {
	l.logger.Sugar().Panicf(format, args...)
}

func (l *logger) Sync() error {
	return l.logger.Sugar().Sync()
}

func (l *logger) WithField(key string, value any) Logger {
	return &logger{
		logger: l.logger.With(zap.Any(key, value)),
	}
}
