package functions

import (
	"reflect"
	"testing"
)

var mylist = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

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

func TestSliceArray(t *testing.T) {
	expectede := []int{2, 4, 6, 8, 10}
	rete := SliceArray(mylist, 1)
	expectedo := []int{1, 3, 5, 7, 9}
	reto := SliceArray(mylist, 0)

	if !reflect.DeepEqual(rete, expectede) {
		t.Errorf("SliceArray() = %q, want %q", rete, expectede)
	}

	if !reflect.DeepEqual(reto, expectedo) {
		t.Errorf("SliceArray() = %q, want %q", reto, expectedo)
	}
}

func TestIndexArray(t *testing.T) {
	expected := 5

	if ret := IndexArray(mylist, 6); ret != expected {
		t.Errorf("SumArray() = %q, want %q", ret, expected)
	}
}

func TestHello(t *testing.T) {
	expected := "Hello World!"
	if ret := Hello(); ret != expected {
		t.Errorf("Hello() = %q, want %q", ret, expected)
	}
}

func TestReverseRunes(t *testing.T) {
	expected := "!dlroW olleH"
	if ret := ReverseRunes("Hello World!"); ret != expected {
		t.Errorf("ReverseRunes() = %q, want %q", ret, expected)
	}
}

func TestForwardDifference(t *testing.T) {
	expected := 4.00001000002703221980
	if ret := ForwardDifference(2.0, 1e-5); ret != expected {
		t.Errorf("ForwardDifference() = %.20f, want %.20f", ret, expected)
	}
}

func TestCentralDifference(t *testing.T) {
	expected := 4.00000000002620481609
	if ret := CentralDifference(2.0, 1e-5); ret != expected {
		t.Errorf("CentralDifference() = %.20f, want %.20f", ret, expected)
	}
}
