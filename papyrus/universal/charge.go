package universal

import (
	"time"
	"github.com/17bixin/gobase/papyrus"
	"github.com/17bixin/gobase/papyrus/alipay"
	"github.com/17bixin/gobase/papyrus/wechat"
)

type ChargeAdapter interface {
	Validator(chargeOp *papyrus.ChargeOption) bool //valid msg from C
	Sign(chargeOp *papyrus.ChargeOption, pap *papyrus.Papyrus) (string,error)
}

func NewCharge(chargeOp *papyrus.ChargeOption, pap *papyrus.Papyrus) (*papyrus.ChargeBody) {
	cb := &papyrus.ChargeBody{}
	cb.TranStatus = papyrus.TranStatusUnpaid
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

func chooseAdapter(c string, pap *papyrus.Papyrus) (ChargeAdapter, error) {
	switch c {
	case "alipay.app.pay":
		return &alipay.AppPay{}, nil
	case "wechat.app.pay":
		return &wechat.AppPay{},nil
	}
	
	return nil, papyrus.ErrorsNew(papyrus.PapyrusErrorChannel)
}
