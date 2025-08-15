package main

/*
GOROOT is where the Go is installed in the system, it contains Go compiler, standard library and tools.
*/

import (
	"fmt"
	"os"
)

func main() {
	// goroot := runtime.GOROOT()
	goroot := os.Getenv("GOROOT")
	fmt.Println("GOROOT: ", goroot)

	entries, _ := os.ReadDir(goroot)
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("📁 %s\n", entry.Name())
		} else {
			fmt.Printf("🗒️ %s\n", entry.Name())
		}
	}
}
