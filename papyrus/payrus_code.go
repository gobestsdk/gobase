package papyrus

type PapyrusCode = uint64

const (
	PapyrusSuccess           PapyrusCode = 100000
	PapyrusErrServer         PapyrusCode = PapyrusSuccess + 1
	PapyrusErrorChannel      PapyrusCode = PapyrusSuccess + 2
	PapyrusErrorAlipayParams PapyrusCode = PapyrusSuccess + 3
	PapyrusErrorAlipaySign   PapyrusCode = PapyrusSuccess + 4
)

var papyrusMsg = map[PapyrusCode]string{
	PapyrusSuccess:           "SUCCESS",
	PapyrusErrorAlipayParams: "params error ",
	PapyrusErrorChannel:      "channel not found ",
	PapyrusErrorAlipaySign:   "sign error ",
}

type PapyrusMsg struct {
	Code PapyrusCode
	Msg  string
}

func (e *PapyrusMsg) Error() string {
	return e.Msg
}

func ErrorsNew(t PapyrusCode) error {
	return &PapyrusMsg{t, GetMsgByType(t)}
}

func GetMsgByType(e PapyrusCode) string {
	return papyrusMsg[e]
}
