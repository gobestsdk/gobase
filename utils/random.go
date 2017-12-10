package utils

import (
	"crypto/rand"
	"io"
	"strconv"
	"github.com/google/uuid"
	"math/big"
	"time"
	mrand "math/rand"
)

//uuid+unix time
func RandId() string {
	return strconv.FormatInt(int64(uuid.New().Time()/10000000000)*10000000000+Int64Range(100000000, 10000000000), 10)
}

func RandIdInt64() int64 {
	return int64(uuid.New().Time()/10000000000)*10000000000 + Int64Range(100000000, 10000000000)
}

//0123456789 select 6 password number
const C_RAND_TMP = "0123456789"

func RandPassword(length int, chars []byte) string {
	newPwd := make([]byte, length)
	random := make([]byte, length+(length/4)) // storage for random bytes.
	charsLength := byte(len(chars))
	max := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, random); err != nil {
			panic(err)
		}
		for _, c := range random {
			if c >= max {
				continue
			}
			newPwd[i] = chars[c%charsLength]
			i++
			if i == length {
				return string(newPwd)
			}
		}
	}
	panic("unreachable")
}

func Int64Range(min, max int64) int64 {
	var result int64
	maxRand := max - min
	b, err := rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
	if err != nil {
		return max
	}
	result = min + b.Int64()
	return result
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandomLetters(n int, alphabets ...byte) string {
	var bytes = make([]byte, n)
	var randby bool
	if num, err := r.Read(bytes); num != n || err != nil {
		mrand.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[mrand.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[mrand.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return Byte2String(bytes)
}
