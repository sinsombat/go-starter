package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	count := make([]int, 1000)
	for i, _ := range count {
		count[i] = i
	}
	wg.Add(len(count))
	for i, v := range count {
		go func(i int, v int) {
			time.Sleep(time.Second)
			_, err := fmt.Fprintf(os.Stdout, "%d -> %d\n", i, v)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}
