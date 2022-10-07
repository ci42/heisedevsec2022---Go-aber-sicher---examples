package main

import (
	"fmt"
	"sync"
)

var j = 0

func StartConcurrentAdd(wg sync.WaitGroup) {
	go func() {
		defer wg.Done()
		j++
	}()
}

func main() {
	//var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		//wg.Add(1)
		//StartConcurrentAdd(wg)
		go func() {
			//	defer wg.Done()
			j++
		}()
	}
	//wg.Wait()

	fmt.Println(j)
}
