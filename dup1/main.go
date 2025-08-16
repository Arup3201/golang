package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}

		counts[input.Text()]++
	}

	for text, c := range counts {
		if c > 0 {
			fmt.Printf("%d\t%s\n", c, text)
		}
	}
}
