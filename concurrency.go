package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func processData(wg *sync.WaitGroup, result *int, data int) {

	defer wg.Done()

	processData := data * 2

	// lock.Lock()

	*result = processData

	// lock.Unlock()
}

func main() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5, 6, 7}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}
	wg.Wait()
	fmt.Print(result)
}
