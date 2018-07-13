package log

import (
	"encoding/json"
	"os"
)

func WriteLog(w bool) {
	writelog = w
}
func makefile() {
	_, err := os.Stat(logpath)
	if err != nil {
		console_printjson(INFO, Fields{"log": "create log file", "filename": logpath})
		fs, err := os.OpenFile(logpath, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			console_printjson(FATAL, Fields{"err": err})
			return
		}
		fs.Close()
	}
}
func write() {
	makefile()

	afs, err := os.OpenFile(logpath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		console_printjson(FATAL, Fields{"log": "open log file", "err": err})
		return
	}
	defer afs.Close()
	for _, arg := range buffer {
		bs, _ := json.Marshal(arg)
		afs.Write(bs)
		afs.Write([]byte("\n"))
	}
	buffer = make([]map[string]interface{}, 0)
}
func Clear_buffer() {
	makefile()

	afs, err := os.OpenFile(logpath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		console_printjson(FATAL, Fields{"log": "open log file", "err": err})
		return
	}
	defer afs.Close()
	for _, arg := range buffer {
		bs, _ := json.Marshal(arg)
		afs.Write(bs)
		afs.Write([]byte("\n"))
	}
	afs.Write([]byte("\n"))
}
