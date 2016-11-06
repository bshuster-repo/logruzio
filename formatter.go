package logruzio

import (
	"strings"

	"github.com/Sirupsen/logrus"
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
	return []byte(dataStr), nil
}
