package string

import (
	"fmt"
	"slices"
	"strings"
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

	want := "string"
	if got := fmt.Sprintf("%T", GoUsage); got != want {
		t.Errorf("type expected %s, got %s", want, got)
	}
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

	want := "uint8"
	if got := fmt.Sprintf("%T", strByte); got != want {
		t.Errorf("type expected %s, got %s", want, got)
	}
}

func TestStringContain(t *testing.T) {
	title := "The lord of rings"

	want := true
	if got := strings.Contains(title, "lord"); got != want {
		t.Errorf("strings.Contains(title, substr)=%t, expected %t", got, want)
	}
}

func TestStringCount(t *testing.T) {
	title := "The lord of rings"

	want := 2
	if got := strings.Count(title, "o"); got != want {
		t.Errorf("strings.Count(title, str)=%d, expected %d", got, want)
	}
}

func TestStringHasPrefic(t *testing.T) {
	word := "testing"
	want := true

	if got := strings.HasPrefix(word, "tes"); got != want {
		t.Errorf("strings.HasPrefix(word, sub)=%t, expected %t", got, want)
	}
}

func TestStringHasSuffix(t *testing.T) {
	word := "testing"
	want := true
	if got := strings.HasSuffix(word, "ng"); got != want {
		t.Errorf("strings.HasPrefix(word, str)=%t, expected %t", got, want)
	}
}

func TestStringIndexFunc(t *testing.T) {
	word := "testing"
	want := 0
	if got := strings.Index(word, "t"); got != want {
		t.Errorf("strings.Index(str)=%d, expected %d", got, want)
	}
}

func TestStringJoin(t *testing.T) {
	helloWorld := []string{"Hello,", "world", "世界"}
	want := "Hello, world 世界"

	if got := strings.Join(helloWorld, " "); got != want {
		t.Errorf("strings.Join(str, sep)=%s, expected %s", got, want)
	}
}

func TestStringRepeat(t *testing.T) {
	a5 := strings.Repeat("a", 5)
	want := "aaaaa"

	if a5 != want {
		t.Errorf("strings.Repeat(str, times)=%s, expected %s", a5, want)
	}
}

func TestStringReplace(t *testing.T) {
	helloWorld := "Hello, world 世界"
	want := "Hell0, w0rld 世界"

	if got := strings.Replace(helloWorld, "o", "0", -1); got != want {
		t.Errorf("strings.Replace()=%s, expected %s", got, want)
	}
}

func TestStringSplit(t *testing.T) {
	helloWorld := "Hello, world 世界"
	want := []string{"Hello,", "world", "世界"}

	if got := strings.Split(helloWorld, " "); !slices.Equal(got, want) {
		t.Errorf("strings.Split()=%x, expected %x", got, want)
	}
}

func TestStringToLower(t *testing.T) {
	helloWorld := "Hello, WorlD 世界"
	want := "hello, world 世界"

	if got := strings.ToLower(helloWorld); got != want {
		t.Errorf("strings.ToLower()=%s, expected %s", got, want)
	}
}

func TestStringToUpper(t *testing.T) {
	helloWorld := "Hello, WorlD 世界"
	want := "HELLO, WORLD 世界"

	if got := strings.ToUpper(helloWorld); got != want {
		t.Errorf("strings.ToUpper()=%s, expected %s", got, want)
	}
}
