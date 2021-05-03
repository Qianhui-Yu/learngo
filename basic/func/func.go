package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (res int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func add(a, b int) int {
	return a + b
}

// 此处op是给定的函数的形参，实际的函数func(int, int)不确定，可以在调用apply时随意给
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 可变式参数列表
func sum(nums ...int) int {
	res := 0
	for i := range nums {
		//res += i // 此处i是参数列表的index而不是实际参数值
		res += nums[i]
	}
	return res
}

func swap(a, b *int) {
	// 不需要再使用临时变量
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	//fmt.Println(
	//	eval(1, 2, "+"),
	//	eval(1, 2, "-"),
	//	eval(1, 2, "*"),
	//	eval(1, 2, "/"),
	//	eval(1, 2, "<<"),
	//)
	if res, err := eval(1, 2, "<<"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(add, 12, 3))
	fmt.Println(apply(
		func(a int, b int) int {
			return a * b
		}, 13, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
