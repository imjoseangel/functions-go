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
func (arrayrange Range) RangeArray() []int {

	NewList := make([]int, arrayrange.MaxList-arrayrange.MinList+1)
	for Item := range NewList {
		NewList[Item] = arrayrange.MinList + Item
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

// Hello ...
// fmt.Println(functions.Hello())
func Hello() string {
	return "Hello World!"
}
