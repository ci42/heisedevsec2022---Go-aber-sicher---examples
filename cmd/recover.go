package main

import (
	"examples/bar"
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("recovered from: %s", r)
		}
	}()
	bar.MaybePanic()
}
