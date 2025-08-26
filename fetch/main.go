package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
test: go run main.go https://meowfacts.herokuapp.com/
*/

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("fetch error: %s", err)
			return
		}

		scanner := bufio.NewScanner(resp.Body)
		b := ""
		for scanner.Scan() {
			b += scanner.Text()
		}

		resp.Body.Close()
		fmt.Printf("%s\n", b)
	}
}
