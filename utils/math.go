package utils

// Average returns the average of a series of numbers
//
// It takes a slice of float64s and returns a float64
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
