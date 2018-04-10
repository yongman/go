//
// log_test.go
// Copyright (C) 2018 YanMing <yming0221@gmail.com>
//
// Distributed under terms of the MIT license.
//

package log

import (
	"os"
	"testing"
)

func TestLogInfo(t *testing.T) {
	SetLevel(DEBUG)

	SetOutput(os.Stdout)

	Debug("debug %d", 1)
	Info("info %d", 2)
	Error("error %d", 3)
	Debugf("debug %d", 1)
	Infof("info %d", 2)
	Errorf("error %d", 3)
	Infof("this is a string, %s %s", "blabla", "blabla")
}
