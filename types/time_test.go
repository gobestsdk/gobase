package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type Test struct {
	Date  JsonTime `json:"date"`
	Name  string   `json:"name"`
	State bool     `json:"state"`
}

func main() {

	var t = Test{}
	t.Date = JsonTime(time.Now())
	t.Name = "Hello World"
	t.State = true
	body, _ := json.Marshal(t)
	fmt.Println(string(body))
}
