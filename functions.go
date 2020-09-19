package functions

// run: `go get github.com/imjoseangel/functions-go`

// Range ...
type Range struct {
	MinList int
	MaxList int
}

// Array ...
type Array []int

// RangeArray ...
// Created this way to use struct with a function inside a module
func (arrayrange Range) RangeArray() []int {

	result := make([]int, arrayrange.MaxList-arrayrange.MinList+1)
	for Item := range result {
		result[Item] = arrayrange.MinList + Item
	}

	return result
}

// SumArray ...
// Created this way to use struct with a function inside a module
func (array Array) SumArray() int {
	result := 0
	for _, numb := range array {
		result += numb
	}
	return result
}

// SliceArray ...
func SliceArray(array []int, start int) []int {

	result := []int{}
	for i := start; i < len(array); i += 2 {
		result = append(result, array[i])
	}
	return result
}

// IndexArray ...
func IndexArray(array []int, item int) int {
	for result := range array {
		if array[result] == item {
			return result
		}
	}
	return -1
}

// Hello ...
// fmt.Println(functions.Hello())
func Hello() string {
	return "Hello World!"
}
