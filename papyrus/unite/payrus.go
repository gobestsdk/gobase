package unite

type IPapyrus interface {
	NewPapyrus(pap Papyrus) IPapyrus
}

type Papyrus struct {
	AppID      string
	MchID      string
	PrivateKey string
	PublicKey  string
	AppKey     string
	AppSecret  string
}
