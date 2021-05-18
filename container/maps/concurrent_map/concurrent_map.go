package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var myMap map[int]int
var mySyncMap sync.Map

func main() {
	//initMap()
	//fmt.Println(myMap)
	//c := concurrentWrite()
	value, ok := mySyncMap.Load(1)
	if !ok {
		fmt.Println(value)
	}
	actual, loaded := mySyncMap.LoadOrStore(1, 10)
	fmt.Printf("actual = %v, loaded = %v\n", actual, loaded)

	resp := make(map[string]interface{})
	resp["name"] = "yuqianhui"
	resp["age"] = 25
	resp["job"] = map[string]interface{}{
		"cname": "bytedance",
		"team": "ecom_arch",
		"since": 2012,
	}

	fmt.Printf("%v", resp)

	var ts interface{}
	var ts64 int64
	ts = int64(100000000)
	ts64, ok = ts.(int64)
	fmt.Println(ts64)
	fmt.Println(ts.(int64))

	initSyncMap()
	c := concurrentWriteSyncMap()
	count := 0
	for {
		count += <-c
		if count == 10 {
			printSyncMap(&mySyncMap)
			break
		}
	}
}

func initMap() {
	myMap = make(map[int]int)
	for i := 0; i < 10; i++ {
		myMap[i] = i
	}
}

func initSyncMap() {
	for i := 0; i < 10; i++ {
		mySyncMap.Store(i, i)
	}
}

// 直接对map并发写会造成data race
func concurrentWrite() chan int {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			myMap[i] = 2 * i + 1
			c <- 1
		}(i)
	}
	return c
}

func concurrentWriteSyncMap() chan int {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			mySyncMap.Store(i, 2 * i + 1)
			c <- 1
		}(i)
	}
	return c
}

func printSyncMap(syncMap *sync.Map) {
	count := 0
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("no.%v key = %v value = %v\n", count, key, value)
		count++

		return true
	})
}