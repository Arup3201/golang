package slice

func GetCommon(first []string, second []string) string {

	for _, m1 := range first {
		for _, m2 := range second {
			if m1 == m2 {
				return m1
			}
		}
	}

	return ""
}

func Equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func Reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func Rotate(s []int, by int) {
	Reverse(s[:by])
	Reverse(s[by:])
	Reverse(s)
}

func NonZero(s []int) []int {
	i := 0
	for _, v := range s {
		if v != 0 {
			s[i] = v
			i++
		}
	}

	return s[:i]
}

func Remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func doSomething(s []int) []int {
	s = append(s, 1)
	return s
}

func SubtractOneFromLength(s []int) []int {
	s = s[0 : len(s)-1]
	return s
}
