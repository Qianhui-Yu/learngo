package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "fisher",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v", r, r) // 接口内含有接口类型和值
	fmt.Println()
	fmt.Print(" > Type switch: ")
	// 可以用r.(type)获得接口类型
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"This is fake contents of imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever)) //这里不能用r，因为r的类型是Retriever，而&retriever是*mock.Retriever同时实现了Poster和Retriever接口

	// Type assertion
	fmt.Println("Type assertion: ")
	if retriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(retriever.Contents)
	} else {
		fmt.Println("Not a mock Retriever.")
	}

}
