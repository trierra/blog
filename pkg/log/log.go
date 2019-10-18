package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/blog/pkg/metrics"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuration
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmed everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

// MyHookImpL implements function Fire() for the logrus.Hook interface
type MyHookImpL struct {
}

// Levels sets log levels on which function Fire will be run
func (l *MyHookImpL) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

// Fire adds error metrics counter for levels: ErrorLevel, FatalLevel and PanicLevel
func (l *MyHookImpL) Fire(entry *logrus.Entry) error {
	metrics.ErrorInc()
	return nil
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
