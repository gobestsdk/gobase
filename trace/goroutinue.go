package trace

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func Getfileandline() (file string, line int) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		//f := runtime.FuncForPC(pc)
		return file, line
	}
	return
}

func Getfileline() (file string) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		//f := runtime.FuncForPC(pc)
		return file + ":" + strconv.Itoa(line)
	}
	return
}
