package types

import (
	"fmt"
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	fmt.Println(BeijingToday().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Parse("2006-01-02", "2021-01-27"))
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2021-01-27 00:00:00"))

	fmt.Println(time.Parse("2006-1-2", "2021-12-3"))
	fmt.Println(struct {
		time.Time
	}{}.Unix())

}
