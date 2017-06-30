package logruzio

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSimpleFormatterFormat(t *testing.T) {
	name := "test-simple-formatter-format"
	var out bytes.Buffer
	h := Hook{HookOpts{
		Conn:      &out,
		Context:   logrus.Fields{"user": "ruszio"},
		Formatter: &SimpleFormatter{},
	}}
	if err := h.Fire(&logrus.Entry{Data: logrus.Fields{"user": "bla"}}); err != nil {
		t.Errorf("%s expected not to fail: %s", name, err)
	}
	data, err := ioutil.ReadAll(&out)
	if err != nil {
		t.Errorf("%s expected not to fail: %s", name, err)
	}
	if expected := "{\"level\":0,\"message\":\"\",\"time\":\"0001-01-01T00:00:00Z\",\"user\":\"bla\"}\n"; string(data) != expected {
		t.Errorf("%s expected data to be %s but got %s", name, expected, string(data))
	}
}
