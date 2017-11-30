package alipay

import (
	"encoding/json"

	"github.com/kafrax/chaos"
	"github.com/17bixin/gobase/papyrus/unite"
)

//https://docs.open.alipay.com/204/105465/
//app pay method
type AppPay struct {
	NotifyUrl      string `json:"-"`
	Body           string `json:"body"`
	Subject        string `json:"subject"`
	OutTradeNo     string `json:"out_trade_no"`
	TotalAmount    string `json:"total_amount"`
	ProductCode    string `json:"product_code"`
	PassbackParams string `json:"passback_params"`
	TimeoutExpress string `json:"timeout_express"`
}

func (a *AppPay) Validator(chargeOp *unite.ChargeOption) bool {
	if chaos.IsAllNilString(
		chargeOp.OrderNo,
		chargeOp.NotifyUrl,
		chargeOp.Subject,
		chargeOp.Amount,
		chargeOp.ClientIP) {
		return false
	}
	return true
}

func (a *AppPay) ChooseApi() string {
	return AliApiPay
}

func (a *AppPay) SetExtraParam() (isNil bool, m map[string]string) {
	var ex = make(map[string]string)
	ex["notify_url"] = a.NotifyUrl
	return false, ex
}

func (a *AppPay) BizContentKey() (isNil bool, key string) {
	return false, "biz_content"
}

func (a *AppPay) BizContentValue() string {
	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return chaos.Byte2String(b)
}

func (a *AppPay)Sign(chargeOp *unite.ChargeOption, pap *unite.Papyrus)string{
	a.Body = chargeOp.Body
	a.TimeoutExpress = chargeOp.TimeoutExpress
	a.PassbackParams = chargeOp.Extra
	a.NotifyUrl = chargeOp.NotifyUrl
	a.Subject = chargeOp.Subject
	a.TotalAmount = chargeOp.Amount
	a.ProductCode = "QUICK_MSECURITY_PAY"
	a.OutTradeNo = chargeOp.OrderNo
	return newAliPay(pap).URLEncode(a)
}