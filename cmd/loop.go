package main

import (
	"fmt"
	"sync"
)

var s = []string{"A", "B", "C"}

func main() {
	var wg sync.WaitGroup
	for i, v := range s {
		//fmt.Printf("index: %d, value: %s\n", i, v)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("index: %d, value: %s\n", i, v)
		}()
	}
	wg.Wait()
}
