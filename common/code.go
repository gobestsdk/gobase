package common

type BiXinCode = uint64

const (
	BiXinCodeSuccess BiXinCode = 100000
	BiXinCodeFailure BiXinCode = BiXinCodeSuccess + 1
)

var biXinMsg = map[uint64]string{
	BiXinCodeSuccess: "SUCCESS",
	BiXinCodeFailure: "FAILURE",
}

func GetBiXinMsg(c BiXinCode) string {
	return biXinMsg[c]
}
