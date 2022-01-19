package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var AccLogger *logrus.Logger
var TraceLogger *logrus.Logger

var AccLog *logrus.Entry
var TraceLog *logrus.Entry

func Initialize(isDebugMode bool) {
	AccLogger = logrus.New()
	AccLogger.SetLevel(logrus.InfoLevel)
	AccLogger.SetOutput(os.Stdout)
	AccLog = logrus.NewEntry(AccLogger)

	TraceLogger = logrus.New()
	if isDebugMode {
		TraceLogger.SetLevel(logrus.DebugLevel)
	} else {
		TraceLogger.SetLevel(logrus.InfoLevel)
	}
	TraceLogger.SetOutput(os.Stdout)
	TraceLog = logrus.NewEntry(TraceLogger)
}

func AddFileWriter() {
	mw1 := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename: "_log/access.log",
		MaxSize:  1000,
		MaxAge:   60,
		Compress: true,
	})
	AccLogger.SetOutput(mw1)

	mw2 := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename: "_log/trace.log",
		MaxSize:  1000,
		MaxAge:   60,
		Compress: true,
	})
	TraceLogger.SetOutput(mw2)
}
