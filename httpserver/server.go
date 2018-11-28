package httpserver

import (
	"fmt"
	"github.com/gobestsdk/gobase/log"

	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Server http的server
// router部分基于gorilla/mux
type Server struct {
	name string //服务名称

	pidTag  string //进程号
	pidFile string //进程文件

	Server http.Server

	httpPort int //http端口

	quitChan    chan interface{}
	quitTimeout time.Duration

	cancel context.CancelFunc
}

// New 生产Server实例
func New(serverName string) *Server {

	return &Server{
		name:     serverName,
		quitChan: make(chan interface{}),
	}
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
func (s *Server) Run() {
	{ //启动时候创建pid文件
		s.touchPidFile()
	}
	go s.httpServer()
	<-s.quitChan

}

func (s *Server) httpServer() {
	log.Info(log.Fields{"app": "http  will Listen", "port": s.httpPort})
	err := s.Server.ListenAndServe(fmt.Sprintf(":%d", s.httpPort), nil)
	if err != nil {
		log.Error(log.Fields{"app": "http    Listen failed", "error": err})
	}
}

// Stop 停止server
func (s *Server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), s.quitTimeout)
	s.Server.Shutdown(ctx)
	s.deletePidFile()
	log.Clear_buffer()
	close(s.quitChan)
}
