package array

import "fmt"

func InitArray3() [3]int {
	var q = [3]int{}
	return q
}

func AssignArray() [3]int {
	var q = [3]int{}
	q[0] = 1
	return q
}

func InitArrayDots() int {
	var q = [...]int{1, 2, 3}
	return len(q)
}

func SearchArray(a [5]int, x int) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

func ArraySum(a [5]int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func GetCurrency(amount int, currency Currency) string {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	return fmt.Sprintf("%d %s", amount, symbol[currency])
}

func Zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
