package tree

import "fmt"

// go语言中nil空指针也能调用函数，但是不能取值
// nil的类型必须是指针，接口，slice等,所以此处接收者不能是treeNode只能是*TreeNode
//func (node *Node) Traverse() {
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}

// 应用函数式编程的更强大的TraverseFunc()
// 不仅能Print()还能传入别的函数干任意想做的事情
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node) // 可以做任何事情
	node.Right.TraverseFunc(f)
}

// 用TraverseFunc()传入Print实现最简单的traverse打印
func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	// 这一步的go表示TraverseFunc往chan里面放node的这个操作是异步的，不一定要完成才返回这个chan
	// 只要外面的goroutine有人读这个chan就行
	go func() {
		node.TraverseFunc(func(node *Node) {
			fmt.Printf("add %d into the channel\n", node.Val)
			out <- node
		})
		close(out)
	}()
	return out
}