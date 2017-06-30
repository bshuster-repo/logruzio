package main

import (
	"github.com/sirupsen/logrus"
	"github.com/bshuster-repo/logruzio"
)

func main() {
	ctx := logrus.Fields{
		"ExamplePath": "examples/simple/main.go",
		"ExampleName": "simple",
	}
	hook, err := logruzio.New("GNbHxMXxGLunaPMSHDBRdrhNgqHinusT", "example-1-app", ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.AddHook(hook)
	logrus.Info("Test example1")
}
