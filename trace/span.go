package trace

import (
	"strconv"
)

// NewID 生成id
func NewspanID(parentspan string, index int) string {
	if parentspan == "" {
		return strconv.Itoa(index)
	} else {
		return parentspan + "." + strconv.Itoa(index)
	}
}
