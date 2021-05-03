package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	for idx, c := range s {
		fmt.Printf("(%d %c)", idx, c)
	}
	fmt.Println()
	for idx, c := range []byte(s) {
		fmt.Printf("(%d %c)", idx, c)
	}
	fmt.Println()
	for idx, c := range []rune(s) {
		fmt.Printf("(%d %c)", idx, c)
	}
	fmt.Println()

	n := utf8.RuneCountInString(s)
	fmt.Println(n)

	ss := " 123 456   哇好神奇  ！ "
	fields := strings.Fields(ss)
	for idx, x := range fields {
		fmt.Println(idx, x)
	}

	res := strings.Join([]string{"aa", "bb", "cc"}, "-")
	fmt.Println(res)

}
