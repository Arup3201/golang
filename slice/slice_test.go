package slice

import "testing"

func TestGetCommon(t *testing.T) {
	months := []string{
		"", "Janurary", "February", "March", "April",
		"May", "June", "July", "Auguest", "September",
		"October", "November", "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	want := "June"
	if got := GetCommon(Q2, summer); got != want {
		t.Errorf("GetCommon()=%q, want %q", got, want)
	}
}

func TestEqual(t *testing.T) {
	s1 := []int{2, 4, 5, 7, 10}
	s2 := []int{2, 4, 5, 7, 10}
	want := true
	if got := Equal(s1, s2); got != want {
		t.Errorf("Equal(s1, s2)=%t, want %t", got, want)
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{2, 4, 5, 7, 10}
	want := []int{10, 7, 5, 4, 2}
	Reverse(numbers)
	if !Equal(numbers, want) {
		t.Errorf("Reverse(numbers)=%x, want %x", numbers, want)
	}
}

func TestRotate(t *testing.T) {
	numbers := []int{2, 4, 5, 7, 10}
	rotateBy := 2
	want := []int{5, 7, 10, 2, 4}

	Rotate(numbers, rotateBy)
	if !Equal(numbers, want) {
		t.Errorf("Rotate(s, by)=%x, want %x", numbers, want)
	}
}

func TestMake(t *testing.T) {
	s := make([]int, 3, 5) // []int{0, 0, 0} len=3, cap=5
	want := 3
	if got := len(s); got != want {
		t.Errorf("len(s)=%d, want %d", got, want)
	}

	want = 5
	if got := cap(s); got != want {
		t.Errorf("cap(slice)=%d, want %d", got, want)
	}
}

func TestAppend(t *testing.T) {
	s := []int{}
	for i := range 6 {
		s = append(s, i)
	}

	want := []int{0, 1, 2, 3, 4, 5}
	if !Equal(s, want) {
		t.Errorf("append %d, want %d", s, want)
	}
}

func TestNonEmpty(t *testing.T) {
	s := []int{1, 0, 3}
	want := []int{1, 3}

	if got := NonZero(s); !Equal(got, want) {
		t.Errorf("NonEmpty(s)=%x, want %x", got, want)
	}
}

func TestRemove(t *testing.T) {
	s := []int{1, 2, 10, 3}
	want := []int{1, 2, 3}

	if got := Remove(s, 2); !Equal(got, want) {
		t.Errorf("Remove(s)=%x, want %x", got, want)
	}
}
