package string

import (
	"fmt"
	"testing"
)

func TestStringLen(t *testing.T) {
	str := "Hello, world"
	want := 12
	if got := len(str); got != want {
		t.Errorf("len(str)=%d, want %d", got, want)
	}
}

func TestStringInd(t *testing.T) {
	str := "Hello, world"
	want := byte(',')

	if got := str[5]; got != want {
		t.Errorf("str[5]=%d, want %d", got, want)
	}
}

func TestRawString(t *testing.T) {
	const GoUsage = `Go is a tool for managing Go source code.
	
	Usage: 
		go command [arguments]
		...`

	t.Log(GoUsage)
}

func TestUTF8Rune(t *testing.T) {
	utf8Rune := '世'
	want := 0x4e16

	if int32(utf8Rune) != int32(want) {
		t.Errorf("expected %x, got %x", want, utf8Rune)
	}
}

func TestStringIndex(t *testing.T) {
	str := "Hello, world"
	strByte := str[0]
	fmt.Printf("%T", strByte)
}

func TestStringLoop(t *testing.T) {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	for i := range len(sample) {
		fmt.Printf("%#U", sample[i])
	}
}
