package middleware

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	logFilePath := "."
	logFileName := "swuops.log"

	fileName := path.Join(logFilePath, logFileName)
	open, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.Error("open logfile failed")
	}

	// init
	logger := logrus.New()
	// set output
	logger.Out = open
	// log level
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		// time format must set to 2006-01-02 15:04:05
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(lfHook)
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		requestMethod := ctx.Request.Method
		requestUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		logger.WithFields(logrus.Fields{
			"status_code":    statusCode,
			"latency_time":   latencyTime,
			"client_ip":      clientIP,
			"request_method": requestMethod,
			"request_uri":    requestUri,
		}).Info()
	}

}
