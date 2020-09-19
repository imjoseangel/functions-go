package functions

// RangeList ...
type RangeList struct {
	MinList int
	MaxList int
}

// MyArray ...
type MyArray []int

// MakeRange ...
func (listrange RangeList) MakeRange() []int {

	NewList := make([]int, listrange.MaxList-listrange.MinList+1)
	for Item := range NewList {
		NewList[Item] = listrange.MinList + Item
	}

	return NewList
}

// SumArray ...
func (array MyArray) SumArray() int {
	result := 0
	for _, numb := range array {
		result += numb
	}
	return result
}
