package wechat

import (
	"reflect"
	"github.com/17bixin/gobase/chaos"
	"github.com/17bixin/gobase/papyrus"
)

//https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=8_1
//app pay method
type AppPay struct {
	TradeType      string            `json:"trade_type"`
	Body           string            `json:"body"`
	NonceStr       string            `json:"nonce_str"`
	NotifyUrl      string            `json:"notify_url"`
	SpbillCreateIp string            `json:"spbill_create_ip"`
	TotalFee       string            `json:"total_fee"`
	OutTradeNo     string            `json:"out_trade_no"`
	Attach         string            `json:"attach"`
	TimeExpire     string            `json:"time_expire"`
	Data           map[string]string `json:"-"`
}

func (*AppPay) ChooseHost() string {
	return ApiUnifyOrder
}

func (a *AppPay) SetExtraParam() {
}

func (a *AppPay) SetActionParam() {
	e := reflect.ValueOf(a)
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		if v := e.Field(i).Interface(); v != nil && t.Field(i).Tag.Get("json") != "-" {
			a.Data[t.Field(i).Tag.Get("json")] = v.(string)
		}
	}
}

func (a *AppPay) GetMapData() map[string]string {
	return a.Data
}

func (a *AppPay) Validator(chargeOp *papyrus.ChargeOption) bool {
	if chaos.IsAllNilString(
		chargeOp.OrderNo,
		chargeOp.NotifyUrl,
		chargeOp.Amount,
		chargeOp.ClientIP,
		chargeOp.NonceStr) {
		return false
	}
	return true
}

func (a *AppPay) Sign(chargeOp *papyrus.ChargeOption, pap *papyrus.Papyrus) (string, error) {
	a.Body = chargeOp.Body
	a.TimeExpire = chargeOp.TimeoutExpress
	a.Attach = chargeOp.Extra
	a.NotifyUrl = chargeOp.NotifyUrl
	a.TotalFee = chargeOp.Amount
	a.OutTradeNo = chargeOp.OrderNo
	a.NonceStr = chargeOp.NonceStr
	a.SpbillCreateIp = chargeOp.ClientIP
	a.TradeType = "APP"
	r, err := newWeChat(pap).doPost(a)
	return chaos.Byte2String(r), err
}
