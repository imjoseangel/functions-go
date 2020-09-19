package functions

// RangeList ...
type RangeList struct {
	MinList int
	MaxList int
}

// Array ...
type Array []int

// MakeRange ...
func (listrange RangeList) MakeRange() []int {

	NewList := make([]int, listrange.MaxList-listrange.MinList+1)
	for Item := range NewList {
		NewList[Item] = listrange.MinList + Item
	}

	return NewList
}

// SumArray ...
func (array Array) SumArray() int {
	result := 0
	for _, numb := range array {
		result += numb
	}
	return result
}
