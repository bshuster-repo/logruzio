package logruzio

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	name := "test-new"
	h, err := New("faketoken", "app1", logrus.Fields{"user": "ruszio"})
	if err != nil {
		t.Errorf("%s expected not to fail: %s", name, err)
	}
	v, k := h.hookOpts.Context["token"]
	if !k {
		t.Errorf("%s token is missing", name)
	}
	token, k := v.(string)
	if !k {
		t.Errorf("%s: token is not a string", name)
	} else if token != "faketoken" {
		t.Errorf("%s expected token to be faketoken but got %s", name, token)
	}
	if h.hookOpts.Formatter == nil {
		t.Errorf("%s expected formatter to be not nil", name)
	}
	if h.hookOpts.Conn == nil {
		t.Errorf("%s expected writer to be not nil", name)
	}
}

func TestFire(t *testing.T) {
	name := "test-fire"
	var out bytes.Buffer
	h := Hook{HookOpts{
		Conn:      &out,
		Context:   logrus.Fields{"user": "ruszio"},
		Formatter: &logrus.JSONFormatter{},
	}}
	if err := h.Fire(&logrus.Entry{Data: logrus.Fields{"user": "bla"}}); err != nil {
		t.Errorf("%s expected not to fail: %s", name, err)
	}
	data, err := ioutil.ReadAll(&out)
	if err != nil {
		t.Errorf("%s expected not to fail: %s", name, err)
	}
	if expected := "{\"level\":\"panic\",\"msg\":\"\",\"time\":\"0001-01-01T00:00:00Z\",\"user\":\"bla\"}\n"; string(data) != expected {
		t.Errorf("%s expected data to be %s but got %s", name, expected, string(data))
	}
}
