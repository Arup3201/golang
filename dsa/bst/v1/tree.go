package v1

type Tree struct {
	Root *Node
	Nodes int
}

func NewTree() *Tree {
	return &Tree{
		Root: nil, 
		Nodes: 0,
	}
}

/* Tree traversal - by default inorder */
func (t *Tree) Print() {
	t.Root.Print()
}

func (t *Tree) Insert(z *Node) {
	x := t.Root
	y := (*Node)(nil)
	for x!=nil {
		y = x
		
		if z.Value > x.Value {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	z.Parent = y
	if y!=nil{
		if z.Value > y.Value {
			y.Right = z
		} else {
			y.Left = z
		}
	} else {
		t.Root = z
	}
}

func (t *Tree) ShiftNodes(u, v *Node) {
	if u.Parent==nil{
		t.Root = v
	} else if u.Parent.Right==u {
		u.Parent.Right = v
	} else {
		u.Parent.Left = v
	}

	if v!=nil {
		v.Parent = u.Parent
	}
}

func (t *Tree) Delete(z *Node) {
	if z.Left==nil {
		t.ShiftNodes(z, z.Right)
	} else if z.Right==nil {
		t.ShiftNodes(z, z.Left)
	} else {
		y, _ := z.Successor()
		if y.Parent!=z{
			t.ShiftNodes(y, y.Right)
			y.Right=z.Right
			y.Right.Parent=y
		}
		t.ShiftNodes(z, y)
		y.Left = z.Left
		y.Left.Parent=y
	}
}

func (t *Tree) Search(v int) (*Node, bool) {
	x := t.Root
	for x!=nil{
		if x.Value==v{
			return x, true
		} else if v > x.Value {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	return (*Node)(nil), false
}
