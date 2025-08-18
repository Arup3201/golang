package maps

import "testing"

func TestMaps1(t *testing.T) {
	var ages = map[string]int{}
	ages["alice"] = 23
	ages["bob"] = 20

	want := 20
	if got := ages["bob"]; got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestMaps2(t *testing.T) {
	var ages = map[string]int{
		"alice": 23,
		"bob":   20,
	}

	want := 23
	if got := ages["alice"]; got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestMaps3(t *testing.T) {
	ages := make(map[string]int)
	ages["alice"] = 23
	ages["bob"] = 20

	want := 0
	if got := ages["frank"]; got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestDelete(t *testing.T) {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["bob"] = 32
	ages["frank"] = 20

	delete(ages, "frank")

	want := 0
	if got := ages["frank"]; got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestFilterAges(t *testing.T) {
	ages := map[string]int{
		"luffy":  17,
		"zoro":   17,
		"usopp":  17,
		"nami":   17,
		"sanji":  16,
		"franky": 27,
		"robin":  31,
		"zimbei": 60,
		"brook":  1200,
	}

	want := "brook"
	if got := FilterAges(ages, 100); got[0] != want {
		t.Errorf("FilterAges(m, x)=%s, want %s", got[0], want)
	}
}

func TestEqual(t *testing.T) {
	x := map[string]int{
		"alice": 21,
		"bob":   20,
	}
	y := map[string]int{
		"alice": 21,
		"bob":   20,
	}

	want := true
	if got := Equal(x, y); got != want {
		t.Errorf("Equal(x, y)=%t, want %t", got, want)
	}
}
