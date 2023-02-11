package logs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type MyFormatter struct{}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	switch level {
	case "ERROR":
		level = color.RedString("%s ", "ERROR")
	case "WARN":
		level = color.YellowString("%s ", "WARN")
	case "INFO":
		level = color.GreenString("%s ", "INFO")
	case "DEBUG":
		level = color.HiBlueString("%s ", "DEBUG")
	default:
		level = color.HiMagentaString("%s ", level)
	}
	strList := strings.Split(entry.Caller.File, "/")
	fileName := strList[len(strList)-1]
	b.WriteString(fmt.Sprintf("%s - %s - [line:%d] - %s msg ▶ %s\n",
		entry.Time.Format("2006-01-02 15:04:05,678"), fileName,
		entry.Caller.Line, level, entry.Message))
	return b.Bytes(), nil
}

func NewLogger(filename string, display bool) *logrus.Logger {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}
	logger := logrus.New()
	if display {
		logger.SetOutput(io.MultiWriter(os.Stdout, f))
	} else {
		logger.SetOutput(io.MultiWriter(f))
	}
	logger.SetReportCaller(true)
	logger.SetFormatter(&MyFormatter{})
	return logger
}
