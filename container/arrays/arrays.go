package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	arr3 := [6]int{1, 3, 5, 7, 9, 0}         //用:=定义数组时必须给初始值
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8} //让编译器来数数组的大小
	var grid [5][2]bool

	fmt.Println(arr1, arr2, arr3, arr4)
	fmt.Println(grid)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
