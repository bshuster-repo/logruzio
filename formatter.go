package logruzio

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// SimpleFormatter represents a simple JSON formatter
type SimpleFormatter struct {
	jsonFmter logrus.JSONFormatter
}

// Format formats the log entry to a JSON format Logstash understands
func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	dataByte, err := f.jsonFmter.Format(entry)
	if err != nil {
		return nil, err
	}

	// This is a hack. I don't like it but it does the job.
	dataStr := string(dataByte)
	dataStr = strings.Replace(dataStr, "\"msg\":", "\"message\":", 1)

	// logz.io wants numeric levels not strings, otherwise you get a parse error
	// ie want "level":4 not "level":"info"
	//
	// could adapt the Format() function from logrus/json_formatter.go,
	// but this certainly "Formats" the output
	//
	// Logrus has six logging levels: Debug, Info, Warning, Error, Fatal and Panic
	// numeric levels are from logrus/logrus.go, var AllLevels
	//
	dataStr = strings.Replace(dataStr, "\"level\":\"panic\"", "\"level\":0", 1)
	dataStr = strings.Replace(dataStr, "\"level\":\"fatal\"", "\"level\":1", 1)
	dataStr = strings.Replace(dataStr, "\"level\":\"error\"", "\"level\":2", 1)
	dataStr = strings.Replace(dataStr, "\"level\":\"warning\"", "\"level\":3", 1)
	dataStr = strings.Replace(dataStr, "\"level\":\"info\"", "\"level\":4", 1)
	dataStr = strings.Replace(dataStr, "\"level\":\"debug\"", "\"level\":5", 1)

	return []byte(dataStr), nil
}
