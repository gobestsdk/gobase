package grpcserver

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "net/http/pprof" //开启pprof

	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/17bixin/gobase/log"
)

// Server 包含了grpc server, 集成assembly下不同的拦截器
type Server struct {
	name        string //服务名称
	hostIP      string //主机ip
	environment string //服务环境,用于区分环境配置

	pidTag  string //进程号
	pidFile string //进程文件

	RPCSvr        *grpc.Server   //rpc服务
	rpcPort       int            //rpc端口
	rpcRegister   func()         //rpc注册
	httpRegisters []HTTPRegister //http注册

	HTTPSvr  *http.Server //http服务
	httpPort int          //http端口

	quitChan    chan interface{}
	quitTimeout time.Duration

	cancel context.CancelFunc
}

// HTTPRegister 注册http回调类型
type HTTPRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

var (
	defaultQuitTimeout = 3 * time.Second
)

// New 生产Server实例
func New() *Server {
	var (
		hostIP      = os.Getenv("ENV_HOST_IP")
		serverName  = os.Getenv("ENV_SERVER_NAME")
		environment = os.Getenv("ENV_ENVIRONMENT")
	)

	return &Server{
		name:        serverName,
		hostIP:      hostIP,
		environment: environment,
		quitChan:    make(chan interface{}),
		quitTimeout: defaultQuitTimeout,
	}
}

// SetHostIP 设置服务ip
func (s *Server) SetHostIP(ip string) {
	s.hostIP = ip
}

// SetPort 设置服务端口
func (s *Server) SetPort(port int) {
	s.rpcPort = port
	s.httpPort = port + 1
}

// SetRPCRegister rpc注册
func (s *Server) SetRPCRegister(register func()) {
	s.rpcRegister = register
}

// SetHTTPRegisters http注册
func (s *Server) SetHTTPRegisters(registers []HTTPRegister) {
	s.httpRegisters = registers
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

// Run 启动server
func (s *Server) Run() {

	log.Info(log.Fields{"app": "Server Running"})
	{ //启动时候创建pid文件
		s.touchPidFile()
	}

	go s.rpcServer()  //rpc服务
	go s.httpServer() //代理http server

	<-s.quitChan
}

// Stop 停止server
func (s *Server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), s.quitTimeout)

	s.HTTPSvr.Shutdown(ctx)
	s.RPCSvr.GracefulStop()

	s.deletePidFile()
	log.Clear_buffer()
	close(s.quitChan)
}

func (s *Server) rpcServer() error {

	s.RPCSvr = grpc.NewServer()
	s.rpcRegister()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.rpcPort))
	if err != nil {

		log.Error(log.Fields{"app": "rpc  Listen failed!", "error": err})
	} else {

		log.Info(log.Fields{"app": "rpc  Listen", "port": s.rpcPort})
	}
	return s.RPCSvr.Serve(lis)
}

func (s *Server) httpServer() error {
	var (
		err    error
		ctx, _ = context.WithCancel(context.Background())
	)

	runtimeMux := runtime.NewServeMux()
	for _, register := range s.httpRegisters {
		err = register(ctx, runtimeMux, fmt.Sprintf("127.0.0.1:%d", s.rpcPort), []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Error(log.Fields{"app": "http  Regist failed", "error": err})
			return err
		}
	}

	handler := http.DefaultServeMux
	handler.Handle("/", runtimeMux)

	s.HTTPSvr = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.httpPort),
		Handler: handler,
	}

	log.Info(log.Fields{"app": "http  will Listen", "port": s.httpPort})
	err = s.HTTPSvr.ListenAndServe()
	if err != nil {
		log.Error(log.Fields{"app": "http    Listen failed", "error": err})
	}
	return err
}
