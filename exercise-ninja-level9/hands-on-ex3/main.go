package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	wg := sync.WaitGroup{}

	wg.Add(110)
	for i := 0; i < 110; i++ {
		go func() {
			val := count
			runtime.Gosched()
			val += 1
			count = val
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}