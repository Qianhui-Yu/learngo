package main

import (
	"fmt"
	"learngo/testing"
)

//func retrieve(url string) string {
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	return string(bytes)
//}

//func getRetriever() testing.Retriever {
//	return testing.Retriever{}
//}

func getRetriever() retriever {
	return testing.Retriever{} // 这样只需改动这一处的实际Retriever即可
}

// ？: Something can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	url := "https://www.imooc.com"
	//retriever := infra.Retriever{}
	//var retriever ? = getRetriever() //需要一个？来实现类型的自由绑定
	retriever := getRetriever()
	fmt.Printf("%s\n", retriever.Get(url))


}
