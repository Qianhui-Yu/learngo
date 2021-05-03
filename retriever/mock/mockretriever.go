package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents:%s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	//如果要实现以下内容的更改，就需要用指针接收者
	r.Contents = form["contents"]
	return "ok"
}

//有函数是指针接收者之后，最好都改成指针接收者
func (r *Retriever) Get(url string) string {
	return r.Contents
}
