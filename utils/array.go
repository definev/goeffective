package utils

type Sortable interface {
	CompareWith(other Sortable) bool
}

type Int int

func (a Int) CompareWith(other Sortable) bool {
	b, ok := other.(Int)
	if !ok {
		return false
	}
	return a > b
}

type Float float64

func (a Float) CompareWith(other Sortable) bool {
	b, ok := other.(Float)
	if !ok {
		return false
	}
	return a > b
}

func Sort[T Sortable](arr *[]T, asc bool) {
	for i := 0; i < len(*arr); i++ {
		for j := i; j < len(*arr); j++ {
			first := (*arr)[i]
			second := (*arr)[j]
			compare := first.CompareWith(second)
			if asc {
				if compare {
					(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
				}
			} else {
				if !compare {
					(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
				}
			}
		}
	}
}
