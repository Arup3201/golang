package main

import (
	"fmt"
	"io/ioutil"
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

		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			log.Fatalf("reading error: %s", err)
		}

		fmt.Printf("%s\n", b)
	}
}
