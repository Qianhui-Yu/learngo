package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

var (
	a = 11
	b = 22
	s = "abcdef"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitializeValue() {
	var a, b int = 1, 2
	var s string = "abc"
	var c bool = true
	fmt.Println(a, b, c, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 1, 2, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 1, 2, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, math.Pi*1i) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(math.Pi*1i)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enums() {
	const (
		c = iota
		cpp
		java
		python
		javascript
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(c, cpp, java, python, javascript, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello")
	fmt.Println(runtime.GOARCH)
	variableZeroValue()
	variableInitializeValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, s)
	euler()
	triangle()
	enums()
}
