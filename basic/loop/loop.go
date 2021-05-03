package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	for ; num > 0; num /= 2 {
		res = strconv.Itoa(num%2) + res
	}
	return res
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

// 用io.Reader类型代替File类型，使得该函数的使用范围大了很多
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(1024),
		convertToBin(0),
	)
	printFile("basic/branch/abc.txt")

	s := `1111
www
xxx
ccc
x

ppppp`
	printFileContents(strings.NewReader(s))
}
