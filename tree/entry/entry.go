package main

import (
	"fmt"
	"learngo/tree"
	"math"
)

// 利用组合来扩展别人的结构: 把别人的结构包装起来，再给这个新的结构挂载新函数，通过新结构访问别人的结构并使用新函数
type myTreeNode struct {
	node *tree.Node
}

// 内嵌 Embedding (可以使用语法糖来省略后续使用新结构时的很多.前缀)
type myTreeNodeEmbedded struct {
	//后续使用时可以省略最后一个变量前面的所有前缀
	//因为相当于把tree这个package里的Node这个结果内嵌在这个新结构中，所以不用再写前面的包名等
	*tree.Node
	//同时tree.Node的所有变量和函数都可以直接被新结构直接调用可以省去.Node, 新结构(.Node).func/var
}

func (myNode *myTreeNode) postOrderTraversal() {
	if myNode == nil || myNode.node == nil {
		return
	}
	//因为通过包装的新结构不能直接取地址，所以要先将其赋给一个新变量
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrderTraversal()
	right.postOrderTraversal()
	myNode.node.Print()
}

func (myNode *myTreeNodeEmbedded) postOrderTraversal() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	//因为通过包装/内嵌的新结构不能直接取地址，所以要先将其赋给一个新变量
	//可以直接调用myNode的Left，可以把.Node省略
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Node.Right}
	left.postOrderTraversal()
	right.postOrderTraversal()
	myNode.Print()
}

func main() {
	//root := tree.Node{Val: 2}
	//内嵌
	root := myTreeNodeEmbedded{&tree.Node{Val: 2}}
	//内嵌后tree.Node能调用的，新结构变量root也能直接调用，这才是内嵌扩展最方便的地方
	//原结构的变量和函数都能直接调用，还能挂载自己新的变量和函数
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Val: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(3)

	//nodes := []tree.Node{
	//	{Val: 3},
	//	{},
	//	{6, nil, &root}, //内嵌后这里&root要改成root.Node
	//}
	//
	//fmt.Println(nodes)

	root.Print()
	root.SetValue(20)
	root.Print()
	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(2)
	pRoot.Print()

	root.Traverse()

	//fmt.Println("Test combination's function: postOrderTraversal.")
	//myRoot := myTreeNode{&root}
	//myRoot.postOrderTraversal()
	fmt.Println("Test Embedding's function: postOrderTraversal.")
	//myRootEmbd := myTreeNodeEmbedded{&root}
	//myRootEmbd.postOrderTraversal()
	root.postOrderTraversal()

	fmt.Println("Test functional programming")
	root.Traverse()
	count := 0
	root.TraverseFunc(func(node *tree.Node) {
		count++
	})
	fmt.Println("Node count:", count)

	fmt.Println("Test Traverse with Channel")
	c := root.TraverseWithChannel()
	max := math.MinInt32
	for n := range c {
		fmt.Printf("Received %d from channel\n", n.Val)
		if n.Val > max {
			max = n.Val
		}
	}
	fmt.Println("Max Node Value =", max)
}
