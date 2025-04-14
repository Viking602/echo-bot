package logger

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

// Logger 是一个封装了 zerolog 的日志记录器
type Logger struct {
	log zerolog.Logger
}

var globalLogLevel = zerolog.InfoLevel
var defaultFields = map[string]interface{}{}

//func init() {
//	// 自定义函数：只显示文件名而不显示完整路径
//	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
//		return fmt.Sprintf("%s:%d", filepath.Base(file), line) // 仅提取文件名
//	}
//}

// SetGlobalLogLevel 设置全局日志级别
func SetGlobalLogLevel(level zerolog.Level) {
	globalLogLevel = level
}

// SetDefaultFields 设置全局默认字段
func SetDefaultFields(fields map[string]interface{}) {
	defaultFields = fields
}

// NewLogger 创建一个新的 Logger 实例
func NewLogger() *Logger {
	// 设置日志输出到标准输出
	output := zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: false,
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	output.TimeFormat = "2006-01-02T15:04:05.000"

	// 通过自定义 FormatLevel 设置日志级别颜色
	output.FormatLevel = func(i interface{}) string {
		if level, ok := i.(string); ok {
			// 根据日志级别自定义颜色
			switch strings.ToUpper(level) {
			case "DEBUG":
				return fmt.Sprintf("\033[36m[%s]\033[0m", strings.ToUpper(level)) // 青色
			case "INFO":
				return fmt.Sprintf("\033[32m[%s]\033[0m", strings.ToUpper(level)) // 绿色
			case "WARN":
				return fmt.Sprintf("\033[33m[%s]\033[0m", strings.ToUpper(level)) // 黄色
			case "ERROR":
				return fmt.Sprintf("\033[31m[%s]\033[0m", strings.ToUpper(level)) // 红色
			default:
				return fmt.Sprintf("[%-5s]", strings.ToUpper(level)) // 默认无颜色
			}
		}
		return "[UNKNOWN]"
	}
	output.FormatMessage = func(i interface{}) string { return fmt.Sprintf("%s", i) }
	output.FormatFieldName = func(i interface{}) string { return fmt.Sprintf("%s: ", i) }
	output.FormatFieldValue = func(i interface{}) string { return fmt.Sprintf("%v", i) }

	// 创建一个新的 zerolog 实例，并设置日志级别为全局日志级别
	logger := zerolog.New(output).
		With().
		Timestamp().
		Fields(defaultFields).
		Logger().
		Level(globalLogLevel)

	return &Logger{
		log: logger,
	}
}

// Info 记录一个 Info 级别的日志
func (l *Logger) Info() *zerolog.Event {
	return l.log.Info()
}

// Error 记录一个 Error 级别的日志
func (l *Logger) Error() *zerolog.Event {
	return l.log.Error()
}

// Debug 记录一个 Debug 级别的日志
func (l *Logger) Debug() *zerolog.Event {
	return l.log.Debug()
}

// Warn 记录一个 Warn 级别的日志
func (l *Logger) Warn() *zerolog.Event {
	return l.log.Warn()
}

// Infof 格式化记录 Info 级别日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.Info().Msgf(format, args)
}

// Errorf 格式化记录 Error 级别日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.Error().Msgf(format, args)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log.Warn().Msgf(format, args)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log.Debug().Msg(fmt.Sprintf(format, args...))

}

// WithField 添加一个字段到日志上下文中，并返回一个新的 Logger 实例
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return l.With(map[string]interface{}{key: value})
}

// WithFields 添加多个字段到日志上下文中，并返回一个新的 Logger 实例
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	return &Logger{
		log: l.log.With().Fields(fields).Logger(),
	}
}

// Str 添加一个字符串类型的字段到日志上下文中，并返回一个新的 Logger 实例
func (l *Logger) Str(key, value string) *Logger {
	return l.WithField(key, value)
}

// Float64 添加一个浮点数类型的字段到日志上下文中，并返回一个新的 Logger 实例
func (l *Logger) Float64(key string, value float64) *Logger {
	return l.WithField(key, value)
}

// With 添加动态字段到日志上下文中
func (l *Logger) With(fields map[string]interface{}) *Logger {
	return &Logger{
		log: l.log.With().Fields(fields).Logger(),
	}
}

// WithDefaultFields 添加全局默认字段
func (l *Logger) WithDefaultFields() *Logger {
	return l.With(defaultFields)
}

// Send 记录一个日志（带日志级别和动态字段）
func (l *Logger) Send(level zerolog.Level, msg string, fields map[string]interface{}) {
	event := l.log.With().Fields(fields).Logger()

	switch level {
	case zerolog.DebugLevel:
		event.Debug().Msg(msg)
	case zerolog.InfoLevel:
		event.Info().Msg(msg)
	case zerolog.WarnLevel:
		event.Warn().Msg(msg)
	case zerolog.ErrorLevel:
		event.Error().Msg(msg)
	default:
		event.Log().Msg(msg) // 通用日志处理
	}
}

// Ctx 从 context 中获取已绑定的 Logger，如果没有则返回默认 Logger
func Ctx(ctx context.Context) *Logger {
	logger, ok := ctx.Value("logger").(*Logger)
	if ok {
		return logger
	}
	return NewLogger()
}

// WithContext 将 Logger 注入 context 中
func (l *Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "logger", l)
}
