package papyrus

type Papyrus struct {
	AppID      string
	MchID      string
	PrivateKey string
	PublicKey  string
	AppKey     string
	AppSecret  string
}

type ChargeBody struct {
	CreatedAt  int64       `json:"created"`
	Created    string      `json:"created_at"`
	TranStatus string      `json:"tran_status"`
	TranMsg    string      `json:"tran_msg"`
	ChargeStr  string      `json:"charge_str"`
	TimePaid   uint64      `json:"time_paid"`
	TimeExpire uint64      `json:"time_expire"`
	PapCode    PapyrusCode `json:"pap_code"`
	PapMsg     string      `json:"pap_msg"`
	ChargeOption
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