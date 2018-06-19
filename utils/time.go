package utils

import (
	"fmt"
	"time"
)

const (
	DateTimeFormart = "2006-01-02 15:04:05"
	DateFormart     = "2006-01-02"
	rFC3339Local    = "2006-01-02T15:04:05"
)

type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(DateFormart))
	return []byte(stamp), nil
}
func (t *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DateFormart+`"`, string(data), time.Local)
	*t = now
	return
}
