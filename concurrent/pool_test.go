package concurrent

import "testing"

func TestExecute(t *testing.T) {

	arrays := make([]string, 0, 10)
	ParallelExecute([]func() interface{}{func() interface{} {
		return "test"

	}}, func(data interface{}) {
		arrays = append(arrays, data.(string))
	})

	for _, value := range arrays {
		println(value)
	}
}
