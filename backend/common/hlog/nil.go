package hlog

// NoneLogger is a wrapper for log nothing
type NoneLogger struct{}

// Info implements Info Logger interface function
func (l *NoneLogger) Info(args ...interface{}) {}

// Infof implements Infof Logger interface function
func (l *NoneLogger) Infof(format string, args ...interface{}) {}

// Infoln implements Infoln Logger interface function
func (l *NoneLogger) Infoln(args ...interface{}) {}

// Error implements Error Logger interface function
func (l *NoneLogger) Error(args ...interface{}) {}

// Errorf implements Errorf Logger interface function
func (l *NoneLogger) Errorf(format string, args ...interface{}) {}

// Errorln implements Errorln Logger interface function
func (l *NoneLogger) Errorln(args ...interface{}) {}

// Warn implements Warn Logger interface function
func (l *NoneLogger) Warn(args ...interface{}) {}

// Warnf implements Warnf Logger interface function
func (l *NoneLogger) Warnf(format string, args ...interface{}) {}

// Warnln implements Warnln Logger interface function
func (l *NoneLogger) Warnln(args ...interface{}) {}

// Debug implements Debug Logger interface function
func (l *NoneLogger) Debug(args ...interface{}) {}

// Debugf implements Debugf Logger interface function
func (l *NoneLogger) Debugf(format string, args ...interface{}) {}

// Debugln implements Debugln Logger interface function
func (l *NoneLogger) Debugln(args ...interface{}) {}

// Fatal implements Fatal Logger interface function
func (l *NoneLogger) Fatal(args ...interface{}) {}

// Fatalf implements Fatalf Logger interface function
func (l *NoneLogger) Fatalf(format string, args ...interface{}) {}

// Fatalln implements Fatalln Logger interface function
func (l *NoneLogger) Fatalln(args ...interface{}) {}

// WithFields implements WithFields Logger interface function
func (l *NoneLogger) WithFields(fields map[string]interface{}) Logger {
	return l
}
