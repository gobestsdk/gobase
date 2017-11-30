package alipay

import (
	"encoding/json"
	"net/url"
	"sort"
	"time"

	"github.com/17bixin/gobase/papyrus"
	"github.com/17bixin/gobase/chaos"
	"github.com/17bixin/gobase/netask"
)

//todo Detecting the error of return system to avoid repeated payment
type AliApi = string

const (
	AliApiPay           AliApi = "AliPay.trade.app.pay"
	AliApiQuery         AliApi = "AliPay.trade.query"
	AliApiTransfer      AliApi = "AliPay.fund.trans.toaccount.transfer"
	AliApiTransferQuery AliApi = "AliPay.fund.trans.order.query"
)

type aliPay struct {
	AppId      string
	privateKey string
	publicKey  string
	domain     string
	format     string
	charset    string
	signType   string
	version    string
	values     *url.Values
}

func newAliPay(pap *papyrus.Papyrus) *aliPay {
	return &aliPay{
		domain:     "https://openapi.AliPay.com/gateway.do",
		format:     "JSON",
		charset:    "uft-8",
		signType:   "RSA2",
		version:    "1.0",
		privateKey: pap.PrivateKey,
		publicKey:  pap.PublicKey,
		AppId:      pap.AppID,
	}
}

// user demo:
// AliPay/AliPay.AppId="12345678"
// AliPay.URLEncode(AliPayer)
func (c *aliPay) URLEncode(a AliPayer) (string,error) {
	u := &url.Values{}
	u.Add("app_id", c.AppId)
	u.Add("method", a.ChooseApi())
	u.Add("format", c.format)
	u.Add("charset", c.charset)
	u.Add("sign_type", c.signType)
	u.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	u.Add("version", c.version)
	if ok, v := a.BizContentKey(); !ok {
		u.Add(v, a.BizContentValue())
	}
	if ok, v := a.SetExtraParam(); !ok {
		for key, value := range v {
			u.Add(key, value)
		}
	}
	sig, err := SignRsa2(allKeys(u), u, chaos.String2Byte(c.privateKey))
	if err != nil {
		return "",err
	}
	u.Add("sign", sig)
	return u.Encode(),nil
}

func allKeys(u *url.Values) (ret []string) {
	for k, _ := range *u {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return
}

//for http post
func (c *aliPay) doPost(method string, obj AliPayer, result interface{}) (err error) {
	r,err:=c.URLEncode(obj)
	if err != nil {
		return err
	}
	resp, err := netask.Post(
		c.domain,
		"application/x-www-form-urlencoded;charset=utf-8",
		false, chaos.String2Byte(r),
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, result)
	if err != nil {
		return err
	}
	return
}
