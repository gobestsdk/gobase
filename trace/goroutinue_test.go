package trace

import "testing"

func TestGetGoroutineID(t *testing.T) {
	println(GetGoroutineID())
}
