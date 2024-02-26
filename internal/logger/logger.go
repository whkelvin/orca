package logger

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var (
	defaultLogger = getDefaultLogger()
)

func getDefaultLogger() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
		TimeFormat:      time.Kitchen,
		Prefix:          "Orca CLI",
		Level:           log.DebugLevel,
	})
	logger.SetStyles(getLogStyles())
	return logger
}

func getLogStyles() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().SetString("DEBU:").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("6"))

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().SetString("INFO:").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("2"))

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().SetString("WARN:").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("3"))

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().SetString("ERRO:").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("1"))

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().SetString("FATA:").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("5"))

	return styles
}

func Debug(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Debug(msg, keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Info(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Warn(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Error(msg, keyvals...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Fatal(msg, keyvals...)
}

func Print(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Print(msg, keyvals...)
}

func Debugf(msg string, args ...interface{}) {
	defaultLogger.Debugf(msg, args...)
}
func Infof(msg string, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}
func Warnf(msg string, args ...interface{}) {
	defaultLogger.Warnf(msg, args...)
}
func Errorf(msg string, args ...interface{}) {
	defaultLogger.Errorf(msg, args...)
}
func Fatalf(msg string, args ...interface{}) {
	defaultLogger.Fatalf(msg, args...)
}
func Printf(msg string, args ...interface{}) {
	defaultLogger.Printf(msg, args...)
}

type ILogger interface {
	Debug(msg interface{}, keyvals ...interface{})
	Info(msg interface{}, keyvals ...interface{})
	Warn(msg interface{}, keyvals ...interface{})
	Error(msg interface{}, keyvals ...interface{})
	Fatal(msg interface{}, keyvals ...interface{})
	Print(msg interface{}, keyvals ...interface{})

	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	Printf(msg string, args ...interface{})
}

type CapsuleLogger struct {
	logger *log.Logger
}

func NewCapsuleLogger() *CapsuleLogger {
	return &CapsuleLogger{
		logger: getDefaultLogger(),
	}
}

func (l *CapsuleLogger) Debug(msg interface{}, keyvals ...interface{}) {
	l.logger.Debug(msg, keyvals...)
}

func (l *CapsuleLogger) Info(msg interface{}, keyvals ...interface{}) {
	l.logger.Info(msg, keyvals...)
}

func (l *CapsuleLogger) Warn(msg interface{}, keyvals ...interface{}) {
	l.logger.Warn(msg, keyvals...)
}

func (l *CapsuleLogger) Error(msg interface{}, keyvals ...interface{}) {
	l.logger.Error(msg, keyvals...)
}

func (l *CapsuleLogger) Fatal(msg interface{}, keyvals ...interface{}) {
	l.logger.Fatal(msg, keyvals...)
}

func (l *CapsuleLogger) Print(msg interface{}, keyvals ...interface{}) {
	l.logger.Print(msg, keyvals...)
}

func (l *CapsuleLogger) Debugf(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}
func (l *CapsuleLogger) Infof(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}
func (l *CapsuleLogger) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}
func (l *CapsuleLogger) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}
func (l *CapsuleLogger) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}
func (l *CapsuleLogger) Printf(msg string, args ...interface{}) {
	l.logger.Printf(msg, args...)
}
