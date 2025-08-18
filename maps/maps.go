package maps

func FilterAges(ages map[string]int, threshold int) []string {
	filter := []string{}
	for name, age := range ages {
		if age >= threshold {
			filter = append(filter, name)
		}
	}
	return filter
}

func Equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for kx, vx := range x {
		if vy, ok := y[kx]; !ok || vy != vx {
			return false
		}
	}

	return true
}
