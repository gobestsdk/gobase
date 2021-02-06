package simplerouter

import (
	"github.com/gobestsdk/gobase/httpserver"
	"net/http"
)

var (
	Apiserver  *httpserver.HttpServer
	apihandler func(*httpserver.HttpContext)
)

func CreateRouter(name, port string, api func(*httpserver.HttpContext)) {
	Apiserver = httpserver.New(name)
	Apiserver.SetPort(port)
	Apiserver.Mux.HandleFunc("/", handler)
	apihandler = api

}
func handler(writer http.ResponseWriter, request *http.Request) {
	ctx := httpserver.CreateHttpContext(writer, request)
	ctx.Query = httpserver.Getfilter(ctx.Request)
	defer httpserver.Recovery(ctx)
	apihandler(ctx)

}
