package mygologs

import "fmt"

type MyGoLogger struct{}

func NewMyGoLogger() *MyGoLogger {
	return &MyGoLogger{}
}

func (logger *MyGoLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args)
}
