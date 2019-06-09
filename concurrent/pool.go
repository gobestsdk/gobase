package concurrent

import "sync"

//ParallelExecute 并行处理
func ParallelExecute(execs []func() interface{}, merge func(data interface{})) {
	dataChan := make(chan interface{})
	wait := sync.WaitGroup{}

	for _, exec := range execs {
		wait.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {

				}
			}()
			dataChan <- exec()
		}()
	}

	closeChan := make(chan struct{})
	go func(dataChan chan interface{}) {
		for {
			select {
			case data := <-dataChan:
				merge(data)
				wait.Done()
			case <-closeChan:
				if len(dataChan) > 0 {
					merge(<-dataChan)
					wait.Done()
				}
				return
			}
		}
	}(dataChan)
	wait.Wait()
	close(closeChan)
}
