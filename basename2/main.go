package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func basename(path string) string {
	var filebase string
	filebase = path

	slash := strings.LastIndex(filebase, "/")
	if slash > -1 {
		filebase = filebase[slash+1:]
	}

	dot := strings.LastIndex(filebase, ".")
	if dot > -1 {
		filebase = filebase[:dot]
	}

	return filebase
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	files := []string{}
	for reader.Scan() {
		files = append(files, reader.Text())
	}

	for _, file := range files {
		fmt.Printf("%s\n", basename(file))
	}
}
