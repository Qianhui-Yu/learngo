package main

import "fmt"

type group struct {
	id string
	name string
}

func main() {
	var mm map[int]int
	fmt.Println(len(mm))
	m1 := map[string]string{
		"name":   "qianhui",
		"gender": "male",
		"age":    "25",
		"xxx":    "yyy",
	}
	m2 := make(map[string]int)       // m2 == empty map
	fmt.Println(len(m2))
	var m3 map[int]map[string]string // m3 == nil go中的nil可以安全参与运算

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	for k, v := range m1 {
		fmt.Println(k, v) // 此处为hashmap，key是无序的，所以遍历的结果中key的顺序不确定
	}

	delete(m1, "xxx")
	for k, v := range m1 {
		fmt.Println(k, v) // 此处为hashmap，key是无序的，所以遍历的结果中key的顺序不确定
	}

	m2["a"] = 1
	m2["b"] = 2
	for k, v := range m2 {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	v, ok := m2["c"]
	fmt.Println(v, ok) // "c"不存在，所以v是int的初始值0

	var groupMap = map[string]*group{}
	fmt.Println(groupMap) // map[]
	g := groupMap["xxx"]
	fmt.Println(g == nil) // true
	if g, ok := groupMap["yyy"]; ok {
		fmt.Println(g)
	} else {
		fmt.Println("val is nil")
	}
}
