package unite

import (
	"github.com/17bixin/gobase/papyrus"
	"github.com/17bixin/gobase/papyrus/alipay"
)

type ChargeAdapter interface {
	Validator(charge map[string]string) bool           //valid msg from C
	Convert2SP(charge map[string]string) ChargeAdapter //convert msg C2Server provider
	Sign() string                                      //sign action
}

type ChargeOption struct {
	OrderNo  string `json:"order_no"`
	Channel  string `json:"channel"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	ClientIP string `json:"client_ip"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

//for pay
type ChargeBody struct {
	Id         string `json:"id"`
	Action     string `json:"action"`
	Created    uint64 `json:"created"`
	TranStatus string `json:"tran_status"`
	TranMsg    string `json:"tran_msg"`
	App        string `json:"app"`
	Channel    string `json:"channel"`
	OrderNo    string `json:"order_no"`
	ClientIp   string `json:"client_ip"`
	Amount     uint64 `json:"amount"`
	Currency   string `json:"currency"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Extra      string `json:"extra"`
	ChargeStr  string `json:"charge_str"`
	TimePaid   uint64 `json:"time_paid"`
	TimeExpire uint64 `json:"time_expire"`
	PapCode    uint32 `json:"pap_code"`
	PapMsg     string `json:"pap_msg"`
}

func NewCharge(chargeOp *ChargeOption, pap *Papyrus) *ChargeBody {
	return nil
}

func chooseAdapter(c string, pap *Papyrus) (ChargeAdapter, error) {
	switch c {
	case "alipay.app":
		return alipay.New(pap).URLEncode()
	}

	return nil, payrus.ErrorsNew("no channel", payrus.ErrorChannel)
}

func (c *ChargeOption) Validator() bool {
	panic("implement me")
}

func (c *ChargeOption) Convert2SP() ChargeAdapter {
	panic("implement me")
}

func (c *ChargeOption) Convert2C(ChargeAdapter) ChargeBody {
	panic("implement me")
}

//convert ChargeOption to ChargeResponse
func New(request *ChargeOption, pap *Papyrus) (c *ChargeBody, err error) {
	//chooseAdapter
	ad, err := chooseAdapter(request.Channel, pap)
	if err != nil {

	}
	//validator ChargeOption is correct
	//convert ChargeOption to ChargeAdapter
	//pay action
	//convert ChargeAdapter to ChargeResponse

}
