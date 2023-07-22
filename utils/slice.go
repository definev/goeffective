package utils

import "fmt"

func SliceExample() {
	sliceOne := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		sliceOne = append(sliceOne, i)
	}
	fmt.Printf("sliceOne: len(%d), cap(%d)\n", len(sliceOne), cap(sliceOne))
	subSliceOne := sliceOne[13:49]
	fmt.Printf("sliceOne: len(%d), cap(%d)\n", len(subSliceOne), cap(subSliceOne))
}
