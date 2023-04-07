package hlogrus

import (
	"github.com/Ralphbaer/hubla/backend/common/hlog"
	"github.com/sirupsen/logrus"
)

// LogrusLogger is a wrapper of logging.Logger with logrus.Logger
type LogrusLogger struct {
	Logger *logrus.Logger
}

// LogrusEntryLogger is a wrapper of logging.Logger with logrus.Entry
type LogrusEntryLogger struct {
	Logger *logrus.Entry
}

// Info implements Info Logger interface function
func (l *LogrusLogger) Info(args ...interface{}) { l.Logger.Info(args...) }

// Infof implements Infof Logger interface function
func (l *LogrusLogger) Infof(format string, args ...interface{}) { l.Logger.Infof(format, args...) }

// Infoln implements Infoln Logger interface function
func (l *LogrusLogger) Infoln(args ...interface{}) { l.Logger.Infoln(args...) }

// Error implements Error Logger interface function
func (l *LogrusLogger) Error(args ...interface{}) { l.Logger.Error(args...) }

// Errorf implements Errorf Logger interface function
func (l *LogrusLogger) Errorf(format string, args ...interface{}) { l.Logger.Errorf(format, args...) }

// Errorln implements Errorln Logger interface function
func (l *LogrusLogger) Errorln(args ...interface{}) { l.Logger.Errorln(args...) }

// Warn implements Warn Logger interface function
func (l *LogrusLogger) Warn(args ...interface{}) { l.Logger.Warn(args...) }

// Warnf implements Warnf Logger interface function
func (l *LogrusLogger) Warnf(format string, args ...interface{}) { l.Logger.Warnf(format, args...) }

// Warnln implements Warnln Logger interface function
func (l *LogrusLogger) Warnln(args ...interface{}) { l.Logger.Warnln(args...) }

// Debug implements Debug Logger interface function
func (l *LogrusLogger) Debug(args ...interface{}) { l.Logger.Debug(args...) }

// Debugf implements Debugf Logger interface function
func (l *LogrusLogger) Debugf(format string, args ...interface{}) { l.Logger.Debugf(format, args...) }

// Debugln implements Debugln Logger interface function
func (l *LogrusLogger) Debugln(args ...interface{}) { l.Logger.Debugln(args...) }

// Fatal implements Fatal Logger interface function
func (l *LogrusLogger) Fatal(args ...interface{}) { l.Logger.Fatal(args...) }

// Fatalf implements Fatalf Logger interface function
func (l *LogrusLogger) Fatalf(format string, args ...interface{}) { l.Logger.Fatalf(format, args...) }

// Fatalln implements Fatalln Logger interface function
func (l *LogrusLogger) Fatalln(args ...interface{}) { l.Logger.Fatalln(args...) }

// WithFields implements WithFields Logger interface function
func (l *LogrusLogger) WithFields(fields map[string]interface{}) hlog.Logger {
	entry := l.Logger.WithFields(fields)
	return &LogrusEntryLogger{
		Logger: entry,
	}
}

// Info implements Info Logger interface function
func (l *LogrusEntryLogger) Info(args ...interface{}) { l.Logger.Info(args...) }

// Infof implements Info Infof interface function
func (l *LogrusEntryLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Infoln implements Infoln Logger interface function
func (l *LogrusEntryLogger) Infoln(args ...interface{}) { l.Logger.Infoln(args...) }

// Error implements Error Logger interface function
func (l *LogrusEntryLogger) Error(args ...interface{}) { l.Logger.Error(args...) }

// Errorf implements Errorf Logger interface function
func (l *LogrusEntryLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Errorln implements Errorln Logger interface function
func (l *LogrusEntryLogger) Errorln(args ...interface{}) { l.Logger.Errorln(args...) }

// Warn implements Warn Logger interface function
func (l *LogrusEntryLogger) Warn(args ...interface{}) { l.Logger.Warn(args...) }

// Warnf implements Warnf Logger interface function
func (l *LogrusEntryLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Warnln implements Warnln Logger interface function
func (l *LogrusEntryLogger) Warnln(args ...interface{}) { l.Logger.Warnln(args...) }

// Debug implements Debug Logger interface function
func (l *LogrusEntryLogger) Debug(args ...interface{}) { l.Logger.Debug(args...) }

// Debugf implements Debugf Logger interface function
func (l *LogrusEntryLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Debugln implements Debugln Logger interface function
func (l *LogrusEntryLogger) Debugln(args ...interface{}) { l.Logger.Debugln(args...) }

// Fatal implements Fatal Logger interface function
func (l *LogrusEntryLogger) Fatal(args ...interface{}) { l.Logger.Fatal(args...) }

// Fatalf implements Fatalf Logger interface function
func (l *LogrusEntryLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// Fatalln implements Fatalln Logger interface function
func (l *LogrusEntryLogger) Fatalln(args ...interface{}) { l.Logger.Fatalln(args...) }

// WithFields implements WithFields Logger interface function
func (l *LogrusEntryLogger) WithFields(fields map[string]interface{}) hlog.Logger {
	l.Logger = l.Logger.WithFields(fields)
	return l
}
