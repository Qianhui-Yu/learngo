package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(arr) // 7,8被替换成10,11但cap不变，append 12时，超过cap，系统会新建一个cap更大的数组并把原数组复制过来

	fmt.Println(len(s3), cap(s3)) // 3 4
	fmt.Println(len(s4), cap(s4)) // 4 4
	fmt.Println(len(s5), cap(s5)) // 5 8 This slice s5 view different array, not the original arr

}
