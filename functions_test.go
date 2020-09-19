package functions

import (
	"reflect"
	"testing"
)

var mylist = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestHello(t *testing.T) {
	expected := "Hello World!"
	if ret := Hello(); ret != expected {
		t.Errorf("Hello() = %q, want %q", ret, expected)
	}
}

func TestRangeArray(t *testing.T) {
	expected := mylist
	newarray := Range{
		MinList: 1,
		MaxList: 10,
	}
	ret := newarray.RangeArray()

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("RangeArray() = %q, want %q", ret, expected)
	}
}

func TestSumArray(t *testing.T) {
	expected := 55
	newarray := Array(mylist)
	if ret := newarray.SumArray(); ret != expected {
		t.Errorf("SumArray() = %q, want %q", ret, expected)
	}
}
