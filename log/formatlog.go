package log

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	PRINT = 0
	INFO  = 1
	WARN  = 2
	ERROR = 3
	FATAL = 4
)

var (
	logpath string = "log"
	level   int
	buffer  []map[string]interface{}
)

type Fields map[string]interface{}

func Setlogfile(f string) {
	logpath = f
}
func Setlevel(l int) {
	level = l
}
func base_print(arg map[string]interface{}) {
	buffer = append(buffer, arg)
	if len(buffer) > 100 {
		write()
	}
}

func console_print(l int, arg map[string]interface{}) {
	bs, _ := json.Marshal(arg)
	var c Colortext
	switch l {
	case PRINT:
		c = White
	case INFO:
		c = Pink
	case WARN:
		c = Yellow
	case ERROR:
		c = LightRed
	case FATAL:
		c = Red
	}
	s := string(c) + string(bs) + string(EndColor)

	fmt.Println(s)

}
func print(l int, arg map[string]interface{}) {
	arg["_"] = time.Now().String()[:19]
	if l >= level {
		base_print(arg)
	}
	console_print(l, arg)
}

func Print(arg map[string]interface{}) {
	print(0, arg)
}

func Info(arg map[string]interface{}) {
	print(1, arg)
}

func Warn(arg map[string]interface{}) {
	print(2, arg)
}

func Error(arg map[string]interface{}) {
	print(3, arg)
}

func Fatal(arg map[string]interface{}) {
	print(4, arg)
}
