package types

type StringSlice []string

func (s StringSlice) Map() (cnt map[string]int) {
	cnt = make(map[string]int)
	for _, v := range ([]string)(s) {
		cnt[v]++
	}
	return
}
