package tree

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

// 工厂函数 来完成构造函数
func CreateTreeNode(val int) *Node {
	return &Node{Val: val} // 局部变量的地址也可以返回给外部函数使用
}

// 此处指定函数的接收者是treeNode，但是*treeNode也可调用该函数
func (node Node) Print() {
	fmt.Println(node.Val)
}

// 此处函数的接收者是*TreeNode,所以可以改变实际的val，若接收者改为treeNode则不行
// 但treeNode也可以调用该函数
func (node *Node) SetValue(val int) {
	node.Val = val
}
