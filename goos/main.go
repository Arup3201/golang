package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	if os == "windows" {
		fmt.Println("Running on windows")
	} else if os == "darwin" {
		fmt.Println("Running on Darwin")
	} else if os == "linux" {
		fmt.Println("Running on Linux")
	} else {
		fmt.Println("Running on", os, " platform")
	}
}
