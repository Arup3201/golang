package array

import "testing"

func TestInitArray3(t *testing.T) {
	want := [3]int{}
	if got := InitArray3(); got != want {
		t.Errorf("InitArray3()=%q, want %q", got, want)
	}
}

func TestAssignArray(t *testing.T) {
	want := [3]int{1, 0, 0}
	if got := AssignArray(); got != want {
		t.Errorf("AssignArray()=%q, want %q", got, want)
	}
}

func TestInitArrayDots(t *testing.T) {
	want := 3
	if got := InitArrayDots(); got != want {
		t.Errorf("InitArrayDots()=%q, want %q", got, want)
	}
}

func TestSearchArray(t *testing.T) {
	arr := [...]int{3, 1, 10, 4, 9}
	x := 4
	want := 3
	if got := SearchArray(arr, x); got != want {
		t.Errorf("SearchArray()=%q, want %q", got, want)
	}
}

func TestArraySum(t *testing.T) {
	arr := [...]int{3, 1, 10, 4, 9}
	want := 27
	if got := ArraySum(arr); got != want {
		t.Errorf("ArraySum()=%d, want %d", got, want)
	}
}

func TestGetCurrency(t *testing.T) {
	amount := 3
	want := "3 ¥"
	if got := GetCurrency(amount, RMB); got != want {
		t.Errorf("GetCurrency()=%s, want %s", got, want)
	}
}

func TestAssignFullArray(t *testing.T) {
	arr := [...]int{1, 2, 3}
	arr = [3]int{}
	if arr[0] != 0 {
		t.Errorf("got %d, want %d", 0, arr[0])
	}
}

func TestZero(t *testing.T) {
	arr := [32]byte{10: 1, 12: 1, 20: 1, 31: 1}
	want := [32]byte{}
	Zero(&arr)
	if arr != want {
		t.Errorf("Zero()=%q, want %q", arr, want)
	}
}
