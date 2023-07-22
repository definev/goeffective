package main

import (
	"fmt"

	"github.com/definev/goeffective/loop"
	"github.com/definev/goeffective/utils"
)

func main() {
	avgSlice := []float64{212, 12, 314, 21, 213, 12, 321}
	avg := utils.Average(avgSlice)
	floatAvgSlice := make([]utils.Float, len(avgSlice))
	{
		for i, v := range avgSlice {
			floatAvgSlice[i] = utils.Float(v)
		}
	}

	utils.Sort(&floatAvgSlice, false)
	fmt.Printf("Average: %f\n", avg)
	fmt.Printf("Sorted: %v\n", floatAvgSlice)
	utils.SliceExample()
	utils.MapExample()
	utils.MakeExample()
	loop.FourTypeOfFor()
}
