package global

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type GLog struct {
	Path        string
	Prefix      string
	MaxLogFiles int
	Level       int // 日志级别，如：0=DEBUG, 1=INFO, 2=WARN, 3=ERROR
}

var GlobalRuntime_log = &GLog{
	Path:        "Logs/",
	MaxLogFiles: 20,
	Prefix:      "runtime",
	Level:       1,
}

var GlobalDebug_log = &GLog{
	Path:        "Logs/Debug/",
	MaxLogFiles: 20,
	Prefix:      "debug",
	Level:       0,
}

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006/01/02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s] [%s][%s %s] %s\n", "[MGL]", timestamp, entry.Level, funcVal, fileVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] [%s] %s\n", "[MGL]", timestamp, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

var (
	Log     *logrus.Logger
	LogFile *os.File
)
