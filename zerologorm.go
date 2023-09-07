package zerologorm

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func SwitchLvl(lvl zerolog.Level) gormlogger.LogLevel {
	switch lvl {
	case zerolog.DebugLevel, zerolog.ErrorLevel:
		return gormlogger.Error
	case zerolog.WarnLevel:
		return gormlogger.Warn
	case zerolog.InfoLevel:
		return gormlogger.Info
	default:
		return gormlogger.Silent
	}
}

type Logger struct {
	l                         *zerolog.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewLogger(l *zerolog.Logger, logLevel gormlogger.LogLevel) *Logger {
	return &Logger{
		l:                         l,
		LogLevel:                  logLevel,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		l:                         l.l,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.l.Info().Ctx(ctx).Msgf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.l.Warn().Ctx(ctx).Msgf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.l.Error().Ctx(ctx).Msgf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.l.Error().Ctx(ctx).Dur("elapsed", elapsed).Int64("rows", rows).Str("sql", sql).Err(err).Msg("trace")
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		l.l.Warn().Ctx(ctx).Dur("elapsed", elapsed).Int64("rows", rows).Str("sql", sql).Msg("trace")
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.l.Debug().Ctx(ctx).Dur("elapsed", elapsed).Int64("rows", rows).Str("sql", sql).Msg("trace")
	}
}
