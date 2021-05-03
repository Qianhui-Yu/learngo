package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i) //参数i在defer语句时就已经运算好
		if i == 30 {
			panic("too many printed")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	//if err != nil {
	//	panic(err)
	//}
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// err = errors.New("this is a custom error") //自定义error
	// 出错处理：是否为已知类型的错误：是，输出信息return；不是，报错panic
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close() // 就近写这个defer，这里是最近的位置，因为如果err了file都没打开

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 将buffer中的内容flush进文件，而不是存在buffer中

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
