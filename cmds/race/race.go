package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// to fix the race condition we can just reassign the i within the for loop
		i := i
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
