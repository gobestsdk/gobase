package unite

import (
	"time"
	"github.com/17bixin/gobase/papyrus"
	"github.com/17bixin/gobase/papyrus/alipay"
	"github.com/17bixin/gobase/papyrus/wechat"
)

type ChargeAdapter interface {
	Validator(chargeOp *ChargeOption) bool //valid msg from C
	Sign(chargeOp *ChargeOption, pap *Papyrus) (string,error)
}

type ChargeOption struct {
	OrderNo        string `json:"order_no"`
	NotifyUrl      string `json:"notify_url"`
	TimeoutExpress string `json:"timeout_express"`
	Channel        string `json:"channel"`
	Amount         string `json:"amount"`
	NonceStr       string `json:"nonce_str"`
	ClientIP       string `json:"client_ip"`
	Subject        string `json:"subject"`
	Body           string `json:"body"`
	Extra          string `json:"extra"`
}

type TranStatus = string

const (
	TranStatusUnpaid TranStatus = "unpaid"
)

var TranMsg = make(map[TranStatus]string)

type ChargeBody struct {
	CreatedAt  int64              `json:"created"`
	Created    string             `json:"created_at"`
	TranStatus string             `json:"tran_status"`
	TranMsg    string             `json:"tran_msg"`
	ChargeStr  string             `json:"charge_str"`
	TimePaid   uint64             `json:"time_paid"`
	TimeExpire uint64             `json:"time_expire"`
	PapCode    papyrus.PapyrusCode`json:"pap_code"`
	PapMsg     string             `json:"pap_msg"`
	ChargeOption
}

func NewCharge(chargeOp *ChargeOption, pap *Papyrus) (*ChargeBody) {
	cb := &ChargeBody{}
	cb.CreatedAt = time.Now().Unix()
	cb.Created = time.Now().Format("2006-01-02 15:-5:05")
	cb.TranStatus = TranStatusUnpaid
	cb.TranMsg = "unpaid"
	cb.PapCode = papyrus.PapyrusSuccess
	cb.PapMsg = papyrus.GetMsgByType(papyrus.PapyrusSuccess)
	cb.ChargeOption = *chargeOp
	ad, err := chooseAdapter(chargeOp.Channel, pap)
	if err != nil {
		cb.PapCode = papyrus.PapyrusErrorChannel
		cb.PapMsg = papyrus.GetMsgByType(papyrus.PapyrusErrorChannel)+err.Error()
		return cb
	}
	if !ad.Validator(chargeOp) {
		cb.PapCode = papyrus.PapyrusErrorAlipayParams
		cb.PapMsg = papyrus.GetMsgByType(papyrus.PapyrusErrorAlipayParams)+err.Error()
		return cb
	}
	cb.ChargeStr,err = ad.Sign(chargeOp, pap)
	if err != nil {
		cb.PapCode = papyrus.PapyrusErrorAlipaySign
		cb.PapMsg = papyrus.GetMsgByType(papyrus.PapyrusErrorAlipaySign)+err.Error()
		return cb
	}
	return cb
}

func chooseAdapter(c string, pap *Papyrus) (ChargeAdapter, error) {
	switch c {
	case "alipay.app.pay":
		return &alipay.AppPay{}, nil
	case "wechat.app.pay":
		return &wechat.AppPay{},nil
	}
	
	return nil, papyrus.ErrorsNew(papyrus.PapyrusErrorChannel)
}
