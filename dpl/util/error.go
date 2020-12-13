package util

import (
	"fmt"
	"runtime"
	"strings"
)

func Error(err interface{}) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Errorf("%s (%s:%d)", err, file[strings.LastIndex(file, "/dueros-dpl/")+1:], line)
}
