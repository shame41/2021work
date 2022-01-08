package datastructure

import "go_test/Tool"

type BST struct{
	Parent *BST
	Left *BST
	Right *BST
	Data int
}

func BuildTreeNode(data int) *BST{
	BST := BST{nil,nil,nil, data}
	return &BST
}

func (x *BST)TreeSearch(k int) *BST{//通过迭代查找值为k的节点
	for x != nil && k != x.Data{
		if k < x.Data{
			x = x.Left
		}else{
			x = x.Right
		}
	}
	return x
}
func (x *BST)TreeMin() *BST{//找到最小节点
	for x.Left != nil{
		x = x.Left
	}
	return x
}
func (x *BST)TreeMax() *BST{//找到最大节点
	for x.Right != nil{
		x = x.Right
	}
	return x
}
func (x *BST)TreeInsert(k int){//向树中插入元素
	var temp *BST
	t := &BST{nil,nil,nil, k}
	for x != nil{
		temp = x
		if t.Data < x.Data {
			x = x.Left
		}else{
			x = x.Right
		}
	}
	t.Parent = temp
	if temp == nil {//树是空的
		*x = *t
	}else if t.Data < temp.Data {
		temp.Left = t
	}else{
		temp.Right = t
	}
}
func (x *BST)TreeDelect(k int){
	t := x.TreeSearch(k)
	if t.Left == nil {
		t.transplant(t.Right)
	}else if t.Right == nil {
		t.transplant(t.Left)
	}else{
		y := t.Right.TreeMin()
		if t.Parent != t{
			y.transplant(y.Right)
			y.Right = t.Right
			y.Right.Parent = y
		}
		t.transplant(y)
		y.Left = t.Left
		y.Left.Parent = y
	}
}
func (x *BST)transplant(newRoot *BST){//将树x接到newRoot上，并用newRoot替换掉x
	if x.Parent == nil {//x是树根
		x = newRoot
	}else if x == x.Parent.Left{//x是左孩子
		x.Parent.Left = newRoot
	}else{//x是右节点
		x.Parent.Right = newRoot
	}
	if newRoot != nil {
		newRoot.Parent = x.Parent
	}
}


func (x *BST)Print(){//中序遍历输出节点
	if x != nil {
		x.Left.Print()
		println(x.Data)
		x.Right.Print()
	}
}
func (t *BST)GetHeight() int{
	if(t == nil){
		return 0
	}
	return Tool.Max(t.Left.GetHeight()+1, t.Right.GetHeight()+1)
}