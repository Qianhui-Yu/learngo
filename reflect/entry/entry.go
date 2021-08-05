package main

import (
	"fmt"
	myref "learngo/reflect"
	"reflect"
)

func main() {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
	// 第三法则：要修改反射对象，其值必须可设置
	v := reflect.ValueOf(&author) // 获取变量指针
	v.Elem().SetString("yuqianhui")
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
	fmt.Println("author = ", author)

	err := fmt.Errorf("test")
	myref.SetErrNil(&err)
	fmt.Println(err)

	fmt.Println(myref.A)

	fmt.Println(myref.A)

	_ = testDefer()

	var obj interface{}
	results := myref.Indirect(reflect.ValueOf(&obj))
	fmt.Println(results.Kind())
}

func testDefer() (err error) {
	defer func() {
		fmt.Println(err)
	}()
	err = fmt.Errorf("err111")
	return func2()
}

func func2() (err error) {
	return fmt.Errorf("err222")
}