package main

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
