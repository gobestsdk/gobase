package server

import (
	"fmt"
	"github.com/17bixin/gobase/log"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Server
type Server struct {
	name        string //服务名称
	hostIP      string //主机ip
	environment string //服务环境

	pidTag  string //进程号
	pidFile string //进程文件

	http.Server

	httpPort int //http端口

	quitChan    chan interface{}
	quitTimeout time.Duration

	cancel context.CancelFunc
}

// SetPort 设置服务端口
func (s *Server) SetPort(port int) {

	s.httpPort = port
}

// touchPidFile 创建pid文件
func (s *Server) touchPidFile() {
	var (
		file = fmt.Sprintf("./%s.pid", s.name)
		pid  = strconv.Itoa(os.Getpid())
	)

	err := ioutil.WriteFile(file, []byte(pid), 0777)
	if err != nil {

		log.Error(log.Fields{"app": "server touch pid file failed!"})
	}

	s.pidFile = file
	s.pidTag = pid

	log.Info(log.Fields{"app": "Process:file touched success.", "pid": s.pidTag})

}

// deletePidFile 删除pid文件
func (s *Server) deletePidFile() {
	_, err := os.Stat(s.pidFile)
	if err != nil {
		// no such file or dir
		return
	}

	os.Remove(s.pidFile)

	log.Info(log.Fields{"app": "Process:file deleted success.", "pid": s.pidTag})
}

// Run server on addr
func (s *Server) Run(addr string) (err error) {

	mux := http.ServeMux{}
	server := http.Server{Handler: mux, Addr: addr}
	err = server.ListenAndServe()

	return err
}

// Close shutdown server forcibly
func (s *Server) Close() error {
	return s.Close()
}
