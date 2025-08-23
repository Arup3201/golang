package v1

import "fmt"

type Node struct {
	Value int
	Parent, Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

/* Traverse the tree with `n` as the root node */
func (n *Node) Print() {
	if n==nil {
		return
	}

	n.Left.Print()
	fmt.Println(n.Value)
	n.Right.Print()
}

func (n *Node) Minimum() *Node {
	x := n
	for x.Left!=nil {
		x = x.Left
	}

	return x
}

func (n *Node) Successor() (*Node, bool) {
	if n.Right!=nil {
		return n.Right.Minimum(), true
	}

	x := n
	y := x.Parent
	for y!=nil && y.Right==x {
		x = y
		y = x.Parent
	}

	if y==nil {
		return (*Node)(nil), false
	} else {
		return y, true
	}
}
