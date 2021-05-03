package main

import (
	"bufio"
	"fmt"
	"io"
	"learngo/functional/fib"
	"strings"
)

//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//为函数实现接口：为fibonacci()实现Reader接口
//1.先为函数设成变量类型
//2.将该类型实现Read函数即实现了Reader接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	} // 只要next不大于10000就继续append下去并可读
	s := fmt.Sprintf("%d\n", next) //每次被读到的一行就是这个
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//f := fibonacci()
	//f := fib.Fibonacci() // 这样f的type是func() int并没有实现Reader接口
	var f intGen = fib.Fibonacci() // 这样定义f就是intGen实现过Reader接口
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	printFileContents(f)
}
