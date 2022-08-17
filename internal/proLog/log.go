package proLog

import (
	"github.com/gin-gonic/gin"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

type Log struct {
	DebugLog *logrus.Logger
	InfoLog  *logrus.Logger
	WarnLog  *logrus.Logger
	ErrorLog *logrus.Logger
	FatalLog *logrus.Logger
	PanicLog *logrus.Logger
}

func (s *Log) DebugGin(c *gin.Context, format string, args ...any) {
	s.DebugLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).Debugf(format, args)
}

func (s *Log) Debug(format string, args ...any) {
	s.DebugLog.Debugf(format, args)
}

func (s *Log) InfoGin(c *gin.Context, format string, args ...any) {
	s.InfoLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).Infof(format, args)
}

func (s *Log) Info(format string, args ...any) {
	s.InfoLog.Infof(format, args)
}

func (s *Log) WarnGin(c *gin.Context, format string, args ...any) {
	s.WarnLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).Warnf(format, args)
}

func (s *Log) Warn(format string, args ...any) {
	s.WarnLog.Warnf(format, args)
}

func (s *Log) ErrorGin(c *gin.Context, format string, args ...any) {
	s.ErrorLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).Errorf(format, args)
}

func (s *Log) Error(format string, args ...any) {
	s.ErrorLog.Errorf(format, args)
}

func (s *Log) WithErrorGin(c *gin.Context, err error) {
	s.ErrorLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).WithError(err)
}

func (s *Log) WithError(err error) {
	s.ErrorLog.WithError(err)
}

func (s *Log) FatalGin(c *gin.Context, format string, args ...any) {
	s.FatalLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).Fatalf(format, args)
}

func (s *Log) Fatal(format string, args ...any) {
	s.FatalLog.Fatalf(format, args)
}

func (s *Log) Panic(format string, args ...any) {
	s.PanicLog.Panicf(format, args)
}

func (s *Log) PanicGin(c *gin.Context, format string, args ...any) {
	s.PanicLog.WithFields(logrus.Fields{
		"RequestId": getRequestId(c),
	}).Panicf(format, args)
}

func InitDebugLog(filePath, fileName string) *logrus.Logger {
	file := path.Join(filePath+"/debug", fileName)

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotate.New(
		file+"-%Y-%m-%d.log",
		rotate.WithMaxAge(180*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return logger
}

func InitInfoLog(filePath, fileName string) *logrus.Logger {
	file := path.Join(filePath+"/info", fileName)

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotate.New(
		file+"-%Y-%m-%d.log",
		rotate.WithMaxAge(180*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return logger
}

func InitWarnLog(filePath, fileName string) *logrus.Logger {
	file := path.Join(filePath+"/warn", fileName)

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotate.New(
		file+"-%Y-%m-%d.log",
		rotate.WithMaxAge(180*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.WarnLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return logger
}

func InitErrorLog(filePath, fileName string) *logrus.Logger {
	file := path.Join(filePath+"/error", fileName)

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotate.New(
		file+"-%Y-%m-%d.log",
		rotate.WithMaxAge(180*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.ErrorLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return logger
}

func InitFatalLog(filePath, fileName string) *logrus.Logger {
	file := path.Join(filePath+"/fatal", fileName)

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotate.New(
		file+"-%Y-%m-%d.log",
		rotate.WithMaxAge(180*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.FatalLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return logger
}

func InitPanicLog(filePath, fileName string) *logrus.Logger {
	file := path.Join(filePath+"/panic", fileName)
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotate.New(
		file+"-%Y-%m-%d.log",
		rotate.WithMaxAge(180*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return logger
}

func InitLogger(filePath, fileName string) *Log {
	return &Log{
		InfoLog:  InitInfoLog(filePath, fileName),
		DebugLog: InitDebugLog(filePath, fileName),
		WarnLog:  InitWarnLog(filePath, fileName),
		ErrorLog: InitErrorLog(filePath, fileName),
		FatalLog: InitFatalLog(filePath, fileName),
		PanicLog: InitPanicLog(filePath, fileName),
	}
}

func getRequestId(c *gin.Context) string {
	reqId, exist := c.Get("reqId")
	if !exist {
		return "Unknown"
	} else {
		return reqId.(string)
	}

}
