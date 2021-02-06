package httpserver

import (
	"github.com/gobestsdk/gobase/trace"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Span struct {
	SpanID   string
	UsedTime time.Duration
	FileLine string
}

//HttpContext 作用
// 1 读写http双工
// 2 记录每个步骤耗费的时间
type HttpContext struct {
	ParentSpan   string
	Writer       http.ResponseWriter
	Realremoteip string
	Request      *http.Request
	//Query
	Query    map[string]interface{}
	Trace    string
	span     int
	UsedTime []Span
	//记录上次时间
	t time.Time
	//Store 注意，线程并不安全，请勿跨线程使用
	Store map[string]interface{}
}

func CreateHttpContext(w http.ResponseWriter, r *http.Request) (hctx *HttpContext) {
	hctx = new(HttpContext)
	hctx.Writer = w
	hctx.Request = r
	hctx.UsedTime = make([]Span, 0)
	hctx.t = time.Now()
	hctx.Store = make(map[string]interface{})
	return
}

func (hctx *HttpContext) Reset() *HttpContext {
	hctx.Writer = nil
	hctx.Request = nil
	hctx.UsedTime = make([]Span, 0)
	hctx.t = time.Unix(0, 0)
	hctx.Store = make(map[string]interface{})
	return hctx
}

func (hctx *HttpContext) HeaderIp(headeripkey string) *HttpContext {

	remote := hctx.Request.Header.Get(headeripkey)
	if len(remote) == 0 {
		remote = hctx.Request.Header.Get(headeripkey)
		if l := strings.Split(remote, ","); len(l) > 0 {
			remote = l[0]
		}
		if len(remote) == 0 {
			remote = hctx.Request.RemoteAddr
		}
	}
	hctx.Realremoteip = remote

	return hctx
}

func (hctx *HttpContext) HeaderTrace(headertracekey string) *HttpContext {
	traceid := hctx.Request.Header.Get(headertracekey)
	if len(traceid) == 0 {
		traceid = trace.NewtraceID(hctx.Realremoteip)
	}
	hctx.Trace = traceid
	return hctx
}
func (hctx *HttpContext) HeaderSpan(headerspankey string) *HttpContext {
	parentspanid := hctx.Request.Header.Get(headerspankey)

	hctx.ParentSpan = parentspanid
	return hctx
}

func (hctx *HttpContext) Step(name string) {
	hctx.span++
	n := time.Now()
	usedTime := n.Sub(hctx.t)
	hctx.t = n
	hctx.UsedTime = append(hctx.UsedTime,
		Span{
			SpanID:   hctx.ParentSpan + "." + strconv.Itoa(hctx.span) + name,
			UsedTime: usedTime,
			FileLine: trace.Getfileline(),
		})
}
func (hctx *HttpContext) Getspan() string {
	if hctx.ParentSpan != "" {
		return hctx.ParentSpan + "." + strconv.Itoa(hctx.span)
	} else {
		return strconv.Itoa(hctx.span)
	}
}

func (hctx *HttpContext) TimeUsedOver(span, total time.Duration) (re bool) {
	t := time.Duration(0)

	for _, u := range hctx.UsedTime {
		t += u.UsedTime
		if u.UsedTime > span {
			re = true
		}
	}
	if total < t {
		re = true
	}
	return
}
