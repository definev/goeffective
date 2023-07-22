package utils

import (
	"fmt"
	"sync"
)

func MakeExample() {

	length := 5
	capacity := 10
	sliceMake := make([]int, length, capacity)
	mapMake := make(map[string]int, length)
	channelMake := make(chan int, 100)

	wait := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func(i int) {
			wait.Add(1)
			channelMake <- i
		}(i)
	}

	fmt.Printf("slice: len(%v) cap(%v)\n", len(sliceMake), cap(sliceMake))
	fmt.Printf("map: len(%v)\n", len(mapMake))
	fmt.Printf("channel: %v\n", channelMake)

	go func() {
		for value := range channelMake {
			wait.Done()
			fmt.Printf("Channel value: %v\n", value)
		}
	}()
	wait.Wait()

}
