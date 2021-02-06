package trace

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

var autoincrease int

func init() {
	autoincrease = 1000
}

// NewID 生成id
//服务器 IP + 产生 ID 时候的时间 + 自增序列 + 当前进程号
func NewtraceID(ip string) (trace string) {

	ipstr := hex.EncodeToString(net.ParseIP(ip).To4())
	unix := time.Now().UnixNano()
	goid := GetGoroutineID()
	trace = fmt.Sprintf("%s%s%4d%05d", ipstr, fmt.Sprint(unix)[:13], autoincrease, goid)
	autoincrease++
	if autoincrease == 10000 {
		autoincrease = 1000
	}
	return trace
}
